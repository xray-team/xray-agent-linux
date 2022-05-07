package netSNMP6

import (
	"github.com/xray-team/xray-agent-linux/conf"
	"github.com/xray-team/xray-agent-linux/dto"
	"github.com/xray-team/xray-agent-linux/logger"
)

type SNMP6DataSource interface {
	GetData() (*dto.NetSNMP6, error)
}

type NetSNMP6Collector struct {
	Config     *conf.NetSNMP6Conf
	DataSource SNMP6DataSource
}

// NewNetSNMP6Collector returns a new collector object.
func NewNetSNMP6Collector(cfg *conf.CollectorsConf, dataSource SNMP6DataSource) dto.Collector {
	if cfg == nil || dataSource == nil {
		logger.Log.Error.Printf(logger.MessageInitCollectorError, CollectorName)

		return nil
	}

	// exit if collector disabled
	if cfg.NetSNMP6 == nil || !cfg.NetSNMP6.Enabled {
		return nil
	}

	return &NetSNMP6Collector{
		Config:     cfg.NetSNMP6,
		DataSource: dataSource,
	}
}

// GetName returns the collector's name.
func (c *NetSNMP6Collector) GetName() string {
	return CollectorName
}

// Collect collects and returns metrics.
func (c *NetSNMP6Collector) Collect() ([]dto.Metric, error) {
	snmp6, err := c.DataSource.GetData()
	if err != nil {
		return nil, err
	}

	metrics := make([]dto.Metric, 0, len(snmp6.Counters))

	for counterName, value := range snmp6.Counters {
		metrics = append(metrics,
			dto.Metric{
				Name:  counterName,
				Value: value,
				Attributes: []dto.MetricAttribute{
					{
						Name:  dto.ResourceAttr,
						Value: dto.ResourceNetSNMP6,
					},
				},
			},
		)
	}

	return metrics, nil
}
