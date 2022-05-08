package netARP

import (
	"strings"

	"github.com/xray-team/xray-agent-linux/conf"
	"github.com/xray-team/xray-agent-linux/dto"
	"github.com/xray-team/xray-agent-linux/logger"
)

type DataSource interface {
	GetData() ([]ARPEntry, error)
}

type Collector struct {
	Config     *conf.NetARPConf
	DataSource DataSource
}

// NewCollector returns a new collector object.
func NewCollector(cfg *conf.CollectorsConf, dataSource DataSource) dto.Collector {
	if cfg == nil || dataSource == nil {
		logger.Log.Error.Printf(logger.MessageInitCollectorError, CollectorName)

		return nil
	}

	// exit if collector disabled
	if cfg.NetARP == nil || !cfg.NetARP.Enabled {
		return nil
	}

	return &Collector{
		Config:     cfg.NetARP,
		DataSource: dataSource,
	}
}

// GetName returns the collector's name.
func (c *Collector) GetName() string {
	return CollectorName
}

// Collect collects and returns metrics.
func (c *Collector) Collect() ([]dto.Metric, error) {
	netArp, err := c.GetNetArp()
	if err != nil {
		return nil, err
	}

	metrics := make([]dto.Metric, 0, len(netArp.Entries)+len(netArp.IncompleteEntries))

	for devName, value := range netArp.Entries {
		devName = strings.ReplaceAll(devName, ".", "_")

		metrics = append(metrics,
			dto.Metric{
				Name:  MetricEntries,
				Value: value,
				Attributes: []dto.MetricAttribute{
					{
						Name:  dto.ResourceAttr,
						Value: ResourceName,
					},
					{
						Name:  SetNameInterface,
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
				Name:  MetricIncompleteEntries,
				Value: value,
				Attributes: []dto.MetricAttribute{
					{
						Name:  dto.ResourceAttr,
						Value: ResourceName,
					},
					{
						Name:  SetNameInterface,
						Value: devName,
					},
				},
			},
		)
	}

	return metrics, nil
}

func (c *Collector) GetNetArp() (*NetArp, error) {
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
