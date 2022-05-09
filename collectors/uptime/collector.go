package uptime

import (
	"github.com/xray-team/xray-agent-linux/dto"
	"github.com/xray-team/xray-agent-linux/logger"
)

type DataSource interface {
	GetData() (*Uptime, error)
}

type Collector struct {
	Config     *Config
	DataSource DataSource
}

// NewCollector returns a new collector object.
func NewCollector(config *Config, dataSource DataSource) dto.Collector {
	if config == nil || dataSource == nil {
		logger.Log.Error.Printf(logger.MessageInitCollectorError, CollectorName)

		return nil
	}

	// exit if collector disabled
	if !config.Enabled {
		return nil
	}

	return &Collector{
		Config:     config,
		DataSource: dataSource,
	}
}

// GetName returns the collector's name.
func (c *Collector) GetName() string {
	return CollectorName
}

// Collect collects and returns metrics.
func (c *Collector) Collect() ([]dto.Metric, error) {
	data, err := c.DataSource.GetData()
	if err != nil {
		return nil, err
	}

	attrs := []dto.MetricAttribute{
		{
			Name:  dto.ResourceAttr,
			Value: ResourceName,
		},
	}

	return []dto.Metric{
		{
			Name:       MetricUptime,
			Value:      data.Uptime,
			Attributes: attrs,
		},
		{
			Name:       MetricIdle,
			Value:      data.Idle,
			Attributes: attrs,
		},
	}, nil
}
