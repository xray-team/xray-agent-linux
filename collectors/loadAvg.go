package collectors

import (
	"errors"

	"github.com/xray-team/xray-agent-linux/conf"
	"github.com/xray-team/xray-agent-linux/dto"
	"github.com/xray-team/xray-agent-linux/logger"
)

type LoadAvgDataSource interface {
	GetData() (*dto.LoadAvg, error)
}

type LoadAvgCollector struct {
	Config     *conf.LoadAvgConf
	DataSource LoadAvgDataSource
}

func NewLoadAvgCollector(cfg *conf.CollectorsConf, dataSource LoadAvgDataSource) dto.Collector {
	if cfg == nil || dataSource == nil {
		logger.LogWarning(logger.CollectorInitPrefix, errors.New("load avg collector init params error"))

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

func (c *LoadAvgCollector) GetName() string {
	return dto.CollectorNameLoadAvg
}

func (c *LoadAvgCollector) Collect() ([]dto.Metric, error) {
	loadAvg, err := c.DataSource.GetData()
	if err != nil {
		return nil, err
	}

	attrs := []dto.MetricAttribute{
		{
			Name:  dto.ResourceAttr,
			Value: dto.ResourceLoadAvg,
		},
	}

	return []dto.Metric{
		{
			Name:       dto.MetricLoadAvgLast,
			Value:      loadAvg.Last,
			Attributes: attrs,
		},
		{
			Name:       dto.MetricLoadAvgLast5m,
			Value:      loadAvg.Last5m,
			Attributes: attrs,
		},
		{
			Name:       dto.MetricLoadAvgLast15m,
			Value:      loadAvg.Last15m,
			Attributes: attrs,
		},
		{
			Name:       dto.MetricLoadAvgKernelSchedulingEntities,
			Value:      loadAvg.KernelSchedulingEntities,
			Attributes: attrs,
		},
	}, nil
}
