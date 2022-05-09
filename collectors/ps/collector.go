package ps

import (
	"github.com/xray-team/xray-agent-linux/dto"
	"github.com/xray-team/xray-agent-linux/logger"
)

type DataSource interface {
	GetData() (*PS, error)
}

type Collector struct {
	Config     *Config
	DataSource DataSource
}

// GetName returns the collector's name.
func (c *Collector) GetName() string {
	return CollectorName
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
			Name:       MetricProcessesCount,
			Value:      data.Count,
			Attributes: attrs,
		},
		{
			Name:       MetricProcessesLimit,
			Value:      data.Limit,
			Attributes: attrs,
		},
		{
			Name:       MetricProcessesInStateRunning,
			Value:      data.InStateRunning,
			Attributes: attrs,
		},
		{
			Name:       MetricProcessesInStateIdle,
			Value:      data.InStateIdle,
			Attributes: attrs,
		},
		{
			Name:       MetricProcessesInStateDead,
			Value:      data.InStateDead,
			Attributes: attrs,
		},
		{
			Name:       MetricProcessesInStateStopped,
			Value:      data.InStateStopped,
			Attributes: attrs,
		},
		{
			Name:       MetricProcessesInStateSleeping,
			Value:      data.InStateSleeping,
			Attributes: attrs,
		},
		{
			Name:       MetricProcessesInStateDiskSleep,
			Value:      data.InStateDiskSleep,
			Attributes: attrs,
		},
		{
			Name:       MetricProcessesInStateZombie,
			Value:      data.InStateZombie,
			Attributes: attrs,
		},
		{
			Name:       MetricProcessesThreadsCount,
			Value:      data.Threads,
			Attributes: attrs,
		},
		{
			Name:       MetricProcessesThreadsLimit,
			Value:      data.ThreadsLimit,
			Attributes: attrs,
		},
	}, nil
}
