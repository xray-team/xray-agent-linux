package loadAvg

import (
	"github.com/xray-team/xray-agent-linux/conf"
	"github.com/xray-team/xray-agent-linux/dto"
	"github.com/xray-team/xray-agent-linux/logger"
)

type LoadAvgDataSource interface {
	GetData() (*LoadAvg, error)
}

type LoadAvgCollector struct {
	Config     *conf.LoadAvgConf
	DataSource LoadAvgDataSource
}

// NewLoadAvgCollector returns a new collector object.
func NewLoadAvgCollector(cfg *conf.CollectorsConf, dataSource LoadAvgDataSource) dto.Collector {
	if cfg == nil || dataSource == nil {
		logger.Log.Error.Printf(logger.MessageInitCollectorError, CollectorName)

		return nil
	}

	// exit if collector disabled
	if cfg.LoadAvg == nil || !cfg.LoadAvg.Enabled {
		return nil
	}

	return &LoadAvgCollector{
		Config:     cfg.LoadAvg,
		DataSource: dataSource,
	}
}

// GetName returns the collector's name.
func (c *LoadAvgCollector) GetName() string {
	return CollectorName
}

// Collect collects and returns metrics.
func (c *LoadAvgCollector) Collect() ([]dto.Metric, error) {
	loadAvg, err := c.DataSource.GetData()
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
			Name:       MetricLast,
			Value:      loadAvg.Last,
			Attributes: attrs,
		},
		{
			Name:       MetricLast5m,
			Value:      loadAvg.Last5m,
			Attributes: attrs,
		},
		{
			Name:       MetricLast15m,
			Value:      loadAvg.Last15m,
			Attributes: attrs,
		},
		{
			Name:       MetricKernelSchedulingEntities,
			Value:      loadAvg.KernelSchedulingEntities,
			Attributes: attrs,
		},
	}, nil
}
