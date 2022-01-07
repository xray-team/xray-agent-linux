package collectors

import (
	"errors"

	"github.com/xray-team/xray-agent-linux/conf"
	"github.com/xray-team/xray-agent-linux/dto"
	"github.com/xray-team/xray-agent-linux/logger"
)

type PSDataSource interface {
	GetData() (*dto.PS, error)
}

type PSCollector struct {
	Config     *conf.PSConf
	DataSource PSDataSource
}

// GetName returns the collector's name.
func (c *PSCollector) GetName() string {
	return dto.CollectorNamePS
}

// NewPSCollector returns a new collector object.
func NewPSCollector(cfg *conf.CollectorsConf, dataSource PSDataSource) dto.Collector {
	if cfg == nil || dataSource == nil {
		logger.LogWarning(logger.CollectorInitPrefix, errors.New("PS collector init params error"))
		return nil
	}

	// exit if collector disabled
	if cfg.PS == nil || !cfg.PS.Enabled {
		return nil
	}

	return &PSCollector{
		Config:     cfg.PS,
		DataSource: dataSource,
	}
}

// Collect collects and returns metrics.
func (c *PSCollector) Collect() ([]dto.Metric, error) {
	ps, err := c.DataSource.GetData()
	if err != nil {
		return nil, err
	}

	attrs := []dto.MetricAttribute{
		{
			Name:  dto.ResourceAttr,
			Value: dto.ResourceProcesses,
		},
	}

	return []dto.Metric{
		{
			Name:       dto.MetricProcessesCount,
			Value:      ps.Count,
			Attributes: attrs,
		},
		{
			Name:       dto.MetricProcessesLimit,
			Value:      ps.Limit,
			Attributes: attrs,
		},
		{
			Name:       dto.MetricProcessesInStateRunning,
			Value:      ps.InStateRunning,
			Attributes: attrs,
		},
		{
			Name:       dto.MetricProcessesInStateIdle,
			Value:      ps.InStateIdle,
			Attributes: attrs,
		},
		{
			Name:       dto.MetricProcessesInStateDead,
			Value:      ps.InStateDead,
			Attributes: attrs,
		},
		{
			Name:       dto.MetricProcessesInStateStopped,
			Value:      ps.InStateStopped,
			Attributes: attrs,
		},
		{
			Name:       dto.MetricProcessesInStateSleeping,
			Value:      ps.InStateSleeping,
			Attributes: attrs,
		},
		{
			Name:       dto.MetricProcessesInStateDiskSleep,
			Value:      ps.InStateDiskSleep,
			Attributes: attrs,
		},
		{
			Name:       dto.MetricProcessesInStateZombie,
			Value:      ps.InStateZombie,
			Attributes: attrs,
		},
		{
			Name:       dto.MetricProcessesThreadsCount,
			Value:      ps.Threads,
			Attributes: attrs,
		},
		{
			Name:       dto.MetricProcessesThreadsLimit,
			Value:      ps.ThreadsLimit,
			Attributes: attrs,
		},
	}, nil
}
