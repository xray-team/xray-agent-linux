package vmStat

import (
	"github.com/xray-team/xray-agent-linux/dto"
	"github.com/xray-team/xray-agent-linux/logger"
)

type DataSource interface {
	GetData() (*VMStat, error)
}

type Collector struct {
	Config     *Config
	DataSource DataSource
}

// CreateCollector returns a new collector object.
func CreateCollector(rawConfig []byte) dto.Collector {
	config := NewConfig()

	err := config.Parse(rawConfig)
	if err != nil {
		logger.Log.Error.Printf(logger.MessageError, CollectorName, err.Error())

		return nil
	}

	return NewCollector(
		config,
		NewDataSource(VMStatPath, CollectorName),
	)
}

// NewCollector returns a new collector object.
func NewCollector(config *Config, dataSource DataSource) dto.Collector {
	if config == nil || dataSource == nil {
		logger.Log.Error.Printf(logger.MessageInitCollectorError, CollectorName)
		return nil
	}

	// exit if collector disabled
	if !config.Enabled {
		return nil
	}

	return &Collector{
		Config:     config,
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

	attrs := []dto.MetricAttribute{
		{
			Name:  dto.ResourceAttr,
			Value: ResourceName,
		},
	}

	return []dto.Metric{
		{
			Name:       MetricPgPgIn,
			Value:      data.PgPgIn,
			Attributes: attrs,
		},
		{
			Name:       MetricPgPgOut,
			Value:      data.PgPgOut,
			Attributes: attrs,
		},
		{
			Name:       MetricPSwpIn,
			Value:      data.PSwpIn,
			Attributes: attrs,
		},
		{
			Name:       MetricPSwpOut,
			Value:      data.PSwpOut,
			Attributes: attrs,
		},
		{
			Name:       MetricPgFault,
			Value:      data.PgFault,
			Attributes: attrs,
		},
		{
			Name:       MetricPgMajFault,
			Value:      data.PgMajFault,
			Attributes: attrs,
		},
		{
			Name:       MetricPgFree,
			Value:      data.PgFree,
			Attributes: attrs,
		},
		{
			Name:       MetricPgActivate,
			Value:      data.PgActivate,
			Attributes: attrs,
		},
		{
			Name:       MetricPgDeactivate,
			Value:      data.PgDeactivate,
			Attributes: attrs,
		},
		{
			Name:       MetricPgLazyFree,
			Value:      data.PgLazyFree,
			Attributes: attrs,
		},
		{
			Name:       MetricPgLazyFreed,
			Value:      data.PgLazyFreed,
			Attributes: attrs,
		},
		{
			Name:       MetricPgRefill,
			Value:      data.PgRefill,
			Attributes: attrs,
		},
		{
			Name:       MetricNumaHit,
			Value:      data.NumaHit,
			Attributes: attrs,
		},
		{
			Name:       MetricNumaMiss,
			Value:      data.NumaMiss,
			Attributes: attrs,
		},
		{
			Name:       MetricNumaForeign,
			Value:      data.NumaForeign,
			Attributes: attrs,
		},
		{
			Name:       MetricNumaInterleave,
			Value:      data.NumaInterleave,
			Attributes: attrs,
		},
		{
			Name:       MetricNumaLocal,
			Value:      data.NumaLocal,
			Attributes: attrs,
		},
		{
			Name:       MetricNumaOther,
			Value:      data.NumaOther,
			Attributes: attrs,
		},
		{
			Name:       MetricOOMKill,
			Value:      data.OOMKill,
			Attributes: attrs,
		},
	}, nil
}
