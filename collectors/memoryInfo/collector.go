package memoryInfo

import (
	"github.com/xray-team/xray-agent-linux/dto"
	"github.com/xray-team/xray-agent-linux/logger"
)

type DataSource interface {
	GetData() (*MemoryInfo, error)
}

type Collector struct {
	Config     *Config
	DataSource DataSource
}

// CreateCollector returns a new collector object.
func CreateCollector(rawConfig []byte) dto.Collector {
	config := NewConfig()

	err := config.Parse(rawConfig)
	if err != nil {
		logger.Log.Error.Printf(logger.MessageError, CollectorName, err.Error())

		return nil
	}

	return NewCollector(
		config,
		NewDataSource(MemInfoPath, CollectorName),
	)
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
			Name:       MetricMemoryTotal,
			Value:      data.MemTotal,
			Attributes: attrs,
		},
		{
			Name:       MetricMemoryFree,
			Value:      data.MemFree,
			Attributes: attrs,
		},
		{
			Name:       MetricMemoryAvailable,
			Value:      data.MemAvailable,
			Attributes: attrs,
		},
		{
			Name:       MetricMemoryUsed,
			Value:      data.MemTotal - data.MemAvailable,
			Attributes: attrs,
		},
		{
			Name:       MetricMemoryBuffers,
			Value:      data.Buffers,
			Attributes: attrs,
		},
		{
			Name:       MetricMemoryCached,
			Value:      data.Cached,
			Attributes: attrs,
		},
		{
			Name:       MetricMemorySwapTotal,
			Value:      data.SwapTotal,
			Attributes: attrs,
		},
		{
			Name:       MetricMemorySwapFree,
			Value:      data.SwapFree,
			Attributes: attrs,
		},
	}, nil
}
