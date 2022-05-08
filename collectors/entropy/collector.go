package entropy

import (
	"github.com/xray-team/xray-agent-linux/conf"
	"github.com/xray-team/xray-agent-linux/dto"
	"github.com/xray-team/xray-agent-linux/logger"
)

type DataSource interface {
	GetData() (*Entropy, error)
}

type Collector struct {
	Config     *conf.EntropyConf
	DataSource DataSource
}

// NewCollector returns a new collector object.
func NewCollector(cfg *conf.CollectorsConf, dataSource DataSource) dto.Collector {
	if cfg == nil || dataSource == nil {
		logger.Log.Error.Printf(logger.MessageInitCollectorError, CollectorName)
		return nil
	}

	// exit if collector disabled
	if cfg.Entropy == nil || !cfg.Entropy.Enabled {
		return nil
	}

	return &Collector{
		Config:     cfg.Entropy,
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

	return []dto.Metric{
		{
			Name:  MetricEntropyAvailable,
			Value: data.Available,
			Attributes: []dto.MetricAttribute{
				{
					Name:  dto.ResourceAttr,
					Value: ResourceName,
				},
			},
		},
	}, nil
}
