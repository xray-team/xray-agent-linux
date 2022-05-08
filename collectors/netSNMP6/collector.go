package netSNMP6

import (
	"github.com/xray-team/xray-agent-linux/conf"
	"github.com/xray-team/xray-agent-linux/dto"
	"github.com/xray-team/xray-agent-linux/logger"
)

type SNMP6DataSource interface {
	GetData() (*NetSNMP6, error)
}

type Collector struct {
	Config     *conf.NetSNMP6Conf
	DataSource SNMP6DataSource
}

// NewCollector returns a new collector object.
func NewCollector(cfg *conf.CollectorsConf, dataSource SNMP6DataSource) dto.Collector {
	if cfg == nil || dataSource == nil {
		logger.Log.Error.Printf(logger.MessageInitCollectorError, CollectorName)

		return nil
	}

	// exit if collector disabled
	if cfg.NetSNMP6 == nil || !cfg.NetSNMP6.Enabled {
		return nil
	}

	return &Collector{
		Config:     cfg.NetSNMP6,
		DataSource: dataSource,
	}
}

// GetName returns the collector's name.
func (c *Collector) GetName() string {
	return CollectorName
}

// Collect collects and returns metrics.
func (c *Collector) Collect() ([]dto.Metric, error) {
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
						Value: ResourceName,
					},
				},
			},
		)
	}

	return metrics, nil
}
