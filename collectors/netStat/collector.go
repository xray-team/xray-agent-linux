package netStat

import (
	"github.com/xray-team/xray-agent-linux/conf"
	"github.com/xray-team/xray-agent-linux/dto"
	"github.com/xray-team/xray-agent-linux/logger"
)

type DataSource interface {
	GetData() (*Netstat, error)
}

type Collector struct {
	Config     *conf.NetStatConf
	DataSource DataSource
}

// NewCollector returns a new collector object.
func NewCollector(cfg *conf.CollectorsConf, dataSource DataSource) dto.Collector {
	if cfg == nil || dataSource == nil {
		logger.Log.Error.Printf(logger.MessageInitCollectorError, CollectorName)
		return nil
	}

	// exit if collector disabled
	if cfg.NetStat == nil || !cfg.NetStat.Enabled {
		return nil
	}

	return &Collector{
		Config:     cfg.NetStat,
		DataSource: dataSource,
	}
}

// GetName returns the collector's name.
func (c *Collector) GetName() string {
	return CollectorName
}

// Collect collects and returns metrics.
func (c *Collector) Collect() ([]dto.Metric, error) {
	netstat, err := c.DataSource.GetData()
	if err != nil {
		return nil, err
	}

	metrics := make([]dto.Metric, 0, 160)

	for extName, stat := range netstat.Ext {
		for metricName, value := range stat {
			metrics = append(metrics,
				dto.Metric{
					Name:  metricName,
					Value: value,
					Attributes: []dto.MetricAttribute{
						{
							Name:  dto.ResourceAttr,
							Value: ResourceName,
						},
						{
							Name:  SetNameExt,
							Value: extName,
						},
					},
				},
			)
		}
	}

	return metrics, nil
}