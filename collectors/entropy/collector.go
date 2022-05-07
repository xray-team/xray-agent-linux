package entropy

import (
	"github.com/xray-team/xray-agent-linux/conf"
	"github.com/xray-team/xray-agent-linux/dto"
	"github.com/xray-team/xray-agent-linux/logger"
)

type EntropyDataSource interface {
	GetData() (*Entropy, error)
}

type EntropyCollector struct {
	Config     *conf.EntropyConf
	DataSource EntropyDataSource
}

// NewEntropyCollector returns a new collector object.
func NewEntropyCollector(cfg *conf.CollectorsConf, dataSource EntropyDataSource) dto.Collector {
	if cfg == nil || dataSource == nil {
		logger.Log.Error.Printf(logger.MessageInitCollectorError, CollectorName)
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

// GetName returns the collector's name.
func (c *EntropyCollector) GetName() string {
	return CollectorName
}

// Collect collects and returns metrics.
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
