package nginx

import (
	"net/http"
	"time"

	"github.com/xray-team/xray-agent-linux/dto"
	"github.com/xray-team/xray-agent-linux/logger"
)

type StubStatusCollector struct {
	Config     *Config
	DataSource StubStatusDataSource
}

// CreateCollector returns a new collector object.
func CreateCollector(rawConfig []byte) dto.Collector {
	config := NewConfig()

	err := config.Parse(rawConfig)
	if err != nil {
		logger.Log.Error.Printf(logger.MessageError, CollectorName, err.Error())

		return nil
	}

	err = config.Validate()
	if err != nil {
		logger.Log.Error.Printf(logger.MessageError, CollectorName, err.Error())

		return nil
	}

	return NewStubStatusCollector(
		config,
		NewStubStatusClient(
			config,
			&http.Client{
				Timeout: time.Second * time.Duration(config.Timeout),
			},
			CollectorName,
		),
	)
}

// NewStubStatusCollector returns a new collector object.
func NewStubStatusCollector(config *Config, dataSource StubStatusDataSource) dto.Collector {
	if config == nil || dataSource == nil {
		return nil
	}

	// exit if collector disabled
	if !config.Enabled {
		return nil
	}

	return &StubStatusCollector{
		Config:     config,
		DataSource: dataSource,
	}
}

type StubStatusDataSource interface {
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
