package collectors

import (
	"errors"

	"xray-agent-linux/conf"
	"xray-agent-linux/dto"
	"xray-agent-linux/logger"
)

type UptimeDataSource interface {
	GetData() (*dto.Uptime, error)
}

type UptimeCollector struct {
	Config     *conf.UptimeConf
	DataSource UptimeDataSource
}

func NewUptimeCollector(cfg *conf.CollectorsConf, dataSource UptimeDataSource) dto.Collector {
	if cfg == nil || dataSource == nil {
		logger.LogWarning(logger.CollectorInitPrefix, errors.New("uptime collector init params error"))
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

func (c *UptimeCollector) GetName() string {
	return dto.CollectorNameUptime
}

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
