package nginx

import (
	"github.com/xray-team/xray-agent-linux/conf"
	"github.com/xray-team/xray-agent-linux/dto"
)

type StubStatusCollector struct {
	Config     *conf.NginxStubStatus
	DataSource NginxStubStatusDataSource
}

// NewStubStatusCollector returns a new collector object.
func NewStubStatusCollector(cfg *conf.CollectorsConf, dataSource NginxStubStatusDataSource) dto.Collector {
	if cfg == nil || dataSource == nil {
		return nil
	}

	// exit if collector disabled
	if cfg.NginxStubStatus == nil || !cfg.NginxStubStatus.Enabled {
		return nil
	}

	return &StubStatusCollector{
		Config:     cfg.NginxStubStatus,
		DataSource: dataSource,
	}
}

type NginxStubStatusDataSource interface {
	GetData() (*StubStatus, error)
}

// GetName returns the collector's name.
func (c *StubStatusCollector) GetName() string {
	return CollectorName
}

// Collect collects and returns metrics.
func (c *StubStatusCollector) Collect() ([]dto.Metric, error) {
	data, err := c.DataSource.GetData()
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
			Name:       MetricStubStatusActive,
			Value:      data.Active,
			Attributes: attrs,
		},
		{
			Name:       MetricStubStatusAccepts,
			Value:      data.Accepts,
			Attributes: attrs,
		},
		{
			Name:       MetricStubStatusHandled,
			Value:      data.Handled,
			Attributes: attrs,
		},
		{
			Name:       MetricStubStatusRequests,
			Value:      data.Requests,
			Attributes: attrs,
		},
		{
			Name:       MetricStubStatusReading,
			Value:      data.Reading,
			Attributes: attrs,
		},
		{
			Name:       MetricStubStatusWriting,
			Value:      data.Writing,
			Attributes: attrs,
		},
		{
			Name:       MetricStubStatusWaiting,
			Value:      data.Waiting,
			Attributes: attrs,
		},
	}, nil
}
