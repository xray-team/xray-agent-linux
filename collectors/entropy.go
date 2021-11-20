package collectors

import (
	"errors"

	"xray-agent-linux/conf"
	"xray-agent-linux/dto"
	"xray-agent-linux/logger"
)

type EntropyDataSource interface {
	GetData() (*dto.Entropy, error)
}

type EntropyCollector struct {
	Config     *conf.EntropyConf
	DataSource EntropyDataSource
}

func NewEntropyCollector(cfg *conf.CollectorsConf, dataSource EntropyDataSource) dto.Collector {
	if cfg == nil || dataSource == nil {
		logger.LogWarning(logger.CollectorInitPrefix, errors.New("entropy collector init params error"))
		return nil
	}

	// exit if collector disabled
	if cfg.Entropy == nil || !cfg.Entropy.Enabled {
		return nil
	}

	return &EntropyCollector{
		Config:     cfg.Entropy,
		DataSource: dataSource,
	}
}

func (c *EntropyCollector) GetName() string {
	return dto.CollectorNameEntropy
}

func (c *EntropyCollector) Collect() ([]dto.Metric, error) {
	data, err := c.DataSource.GetData()
	if err != nil {
		return nil, err
	}

	return []dto.Metric{
		{
			Name:  dto.MetricEntropyAvailable,
			Value: data.Available,
			Attributes: []dto.MetricAttribute{
				{
					Name:  dto.ResourceAttr,
					Value: dto.ResourceEntropy,
				},
			},
		},
	}, nil
}
