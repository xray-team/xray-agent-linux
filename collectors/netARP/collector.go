package netARP

import (
	"strings"

	"github.com/xray-team/xray-agent-linux/conf"
	"github.com/xray-team/xray-agent-linux/dto"
	"github.com/xray-team/xray-agent-linux/logger"
)

type NetARPDataSource interface {
	GetData() ([]ARPEntry, error)
}

type NetARPCollector struct {
	Config     *conf.NetARPConf
	DataSource NetARPDataSource
}

// NewNetARPCollector returns a new collector object.
func NewNetARPCollector(cfg *conf.CollectorsConf, dataSource NetARPDataSource) dto.Collector {
	if cfg == nil || dataSource == nil {
		logger.Log.Error.Printf(logger.MessageInitCollectorError, dto.CollectorNameNetARP)

		return nil
	}

	// exit if collector disabled
	if cfg.NetARP == nil || !cfg.NetARP.Enabled {
		return nil
	}

	return &NetARPCollector{
		Config:     cfg.NetARP,
		DataSource: dataSource,
	}
}

// GetName returns the collector's name.
func (c *NetARPCollector) GetName() string {
	return dto.CollectorNameNetARP
}

// Collect collects and returns metrics.
func (c *NetARPCollector) Collect() ([]dto.Metric, error) {
	netArp, err := c.GetNetArp()
	if err != nil {
		return nil, err
	}

	metrics := make([]dto.Metric, 0, len(netArp.Entries)+len(netArp.IncompleteEntries))

	for devName, value := range netArp.Entries {
		devName = strings.ReplaceAll(devName, ".", "_")

		metrics = append(metrics,
			dto.Metric{
				Name:  dto.MetricNetARPEntries,
				Value: value,
				Attributes: []dto.MetricAttribute{
					{
						Name:  dto.ResourceAttr,
						Value: dto.ResourceNetARP,
					},
					{
						Name:  dto.SetNameNetARPInterface,
						Value: devName,
					},
				},
			},
		)
	}

	for devName, value := range netArp.IncompleteEntries {
		devName = strings.ReplaceAll(devName, ".", "_")

		metrics = append(metrics,
			dto.Metric{
				Name:  dto.MetricNetARPIncompleteEntries,
				Value: value,
				Attributes: []dto.MetricAttribute{
					{
						Name:  dto.ResourceAttr,
						Value: dto.ResourceNetARP,
					},
					{
						Name:  dto.SetNameNetARPInterface,
						Value: devName,
					},
				},
			},
		)
	}

	return metrics, nil
}

func (c *NetARPCollector) GetNetArp() (*NetArp, error) {
	arpTable, err := c.DataSource.GetData()
	if err != nil {
		return nil, err
	}

	var out NetArp
	out.Entries = make(map[string]uint)
	out.IncompleteEntries = make(map[string]uint)

	out.Entries["Total"] = 0
	out.IncompleteEntries["Total"] = 0

	for _, entry := range arpTable {
		if _, ok := out.Entries[entry.Device]; !ok {
			out.Entries[entry.Device] = 0
		}

		if _, ok := out.IncompleteEntries[entry.Device]; !ok {
			out.IncompleteEntries[entry.Device] = 0
		}

		out.Entries["Total"]++
		out.Entries[entry.Device]++

		if entry.HWAddress == "00:00:00:00:00:00" {
			out.IncompleteEntries["Total"]++
			out.IncompleteEntries[entry.Device]++
		}
	}

	return &out, err
}
