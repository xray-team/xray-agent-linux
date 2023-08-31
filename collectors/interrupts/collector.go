package interrupts

import (
	"fmt"

	"github.com/xray-team/xray-agent-linux/dto"
	"github.com/xray-team/xray-agent-linux/logger"
)

type DataSource interface {
	GetData() (*Interrupts, error)
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
		NewDataSource(InterruptsPath, CollectorName),
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

	var result []dto.Metric

	for cpuNum, v := range data.PerCPU {
		metric := dto.Metric{
			Name:       fmt.Sprintf("CPU%d", cpuNum),
			Attributes: attrs,
		}

		metric.Value = v
		result = append(result, metric)
	}

	// Appending total interrupts count
	result = append(result, dto.Metric{
		Name:       "Total",
		Attributes: attrs,
		Value:      data.Total,
	})

	return result, nil
}
