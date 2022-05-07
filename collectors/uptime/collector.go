package uptime

import (
	"github.com/xray-team/xray-agent-linux/conf"
	"github.com/xray-team/xray-agent-linux/dto"
	"github.com/xray-team/xray-agent-linux/logger"
)

type UptimeDataSource interface {
	GetData() (*Uptime, error)
}

type UptimeCollector struct {
	Config     *conf.UptimeConf
	DataSource UptimeDataSource
}

// NewUptimeCollector returns a new collector object.
func NewUptimeCollector(cfg *conf.CollectorsConf, dataSource UptimeDataSource) dto.Collector {
	if cfg == nil || dataSource == nil {
		logger.Log.Error.Printf(logger.MessageInitCollectorError, dto.CollectorNameUptime)

		return nil
	}

	// exit if collector disabled
	if cfg.Uptime == nil || !cfg.Uptime.Enabled {
		return nil
	}

	return &UptimeCollector{
		Config:     cfg.Uptime,
		DataSource: dataSource,
	}
}

// GetName returns the collector's name.
func (c *UptimeCollector) GetName() string {
	return dto.CollectorNameUptime
}

// Collect collects and returns metrics.
func (c *UptimeCollector) Collect() ([]dto.Metric, error) {
	uptime, err := c.DataSource.GetData()
	if err != nil {
		return nil, err
	}

	attrs := []dto.MetricAttribute{
		{
			Name:  dto.ResourceAttr,
			Value: dto.ResourceUptime,
		},
	}

	return []dto.Metric{
		{
			Name:       dto.MetricUptimeUptime,
			Value:      uptime.Uptime,
			Attributes: attrs,
		},
		{
			Name:       dto.MetricUptimeIdle,
			Value:      uptime.Idle,
			Attributes: attrs,
		},
	}, nil
}
