package collectors

import (
	"errors"

	"github.com/xray-team/xray-agent-linux/conf"
	"github.com/xray-team/xray-agent-linux/dto"
	"github.com/xray-team/xray-agent-linux/logger"
)

type NetStatDataSource interface {
	GetData() (*dto.Netstat, error)
}

type NetStatCollector struct {
	Config     *conf.NetStatConf
	DataSource NetStatDataSource
}

func NewNetStatCollector(cfg *conf.CollectorsConf, dataSource NetStatDataSource) dto.Collector {
	if cfg == nil || dataSource == nil {
		logger.LogWarning(logger.CollectorInitPrefix, errors.New("net stat collector init params error"))
		return nil
	}

	// exit if collector disabled
	if cfg.NetStat == nil || !cfg.NetStat.Enabled {
		return nil
	}

	return &NetStatCollector{
		Config:     cfg.NetStat,
		DataSource: dataSource,
	}
}

func (c *NetStatCollector) GetName() string {
	return dto.CollectorNameNetStat
}

func (c *NetStatCollector) Collect() ([]dto.Metric, error) {
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
							Value: dto.ResourceNetStat,
						},
						{
							Name:  dto.SetNameNetStatExt,
							Value: extName,
						},
					},
				},
			)
		}
	}

	return metrics, nil
}
