package netSNMP

import (
	"github.com/xray-team/xray-agent-linux/collectors/netStat"
	"github.com/xray-team/xray-agent-linux/conf"
	"github.com/xray-team/xray-agent-linux/dto"
	"github.com/xray-team/xray-agent-linux/logger"
)

type netSNMPCollector struct {
	Config     *conf.NetSNMPConf
	DataSource netStat.NetStatDataSource
}

// NewNetSNMPCollector returns a new collector object.
func NewNetSNMPCollector(cfg *conf.CollectorsConf, dataSource netStat.NetStatDataSource) dto.Collector {
	if cfg == nil || dataSource == nil {
		logger.Log.Error.Printf(logger.MessageInitCollectorError, CollectorName)

		return nil
	}

	// exit if collector disabled
	if cfg.NetSNMP == nil || !cfg.NetSNMP.Enabled {
		return nil
	}

	return &netSNMPCollector{
		Config:     cfg.NetSNMP,
		DataSource: dataSource,
	}
}

// GetName returns the collector's name.
func (c *netSNMPCollector) GetName() string {
	return CollectorName
}

// Collect collects and returns metrics.
func (c *netSNMPCollector) Collect() ([]dto.Metric, error) {
	netstat, err := c.DataSource.GetData()
	if err != nil {
		return nil, err
	}

	metrics := make([]dto.Metric, 0, 128)

	for extName, stat := range netstat.Ext {
		for metricName, value := range stat {
			metrics = append(metrics,
				dto.Metric{
					Name:  metricName,
					Value: value,
					Attributes: []dto.MetricAttribute{
						{
							Name:  dto.ResourceAttr,
							Value: dto.ResourceNetSNMP,
						},
						{
							Name:  dto.SetNameNetSNMPExt,
							Value: extName,
						},
					},
				},
			)
		}
	}

	return metrics, nil
}
