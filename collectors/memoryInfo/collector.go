package memoryInfo

import (
	"github.com/xray-team/xray-agent-linux/conf"
	"github.com/xray-team/xray-agent-linux/dto"
	"github.com/xray-team/xray-agent-linux/logger"
)

type MemoryDataSource interface {
	GetData() (*MemoryInfo, error)
}

type MemoryInfoCollector struct {
	Config     *conf.MemoryInfoConf
	DataSource MemoryDataSource
}

// NewMemoryInfoCollector returns a new collector object.
func NewMemoryInfoCollector(cfg *conf.CollectorsConf, dataSource MemoryDataSource) dto.Collector {
	if cfg == nil || dataSource == nil {
		logger.Log.Error.Printf(logger.MessageInitCollectorError, CollectorName)

		return nil
	}

	// exit if collector disabled
	if cfg.MemoryInfo == nil || !cfg.MemoryInfo.Enabled {
		return nil
	}

	return &MemoryInfoCollector{
		Config:     cfg.MemoryInfo,
		DataSource: dataSource,
	}
}

// GetName returns the collector's name.
func (c *MemoryInfoCollector) GetName() string {
	return CollectorName
}

// Collect collects and returns metrics.
func (c *MemoryInfoCollector) Collect() ([]dto.Metric, error) {
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
