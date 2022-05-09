package cpuInfo

import (
	"strconv"

	"github.com/xray-team/xray-agent-linux/dto"
	"github.com/xray-team/xray-agent-linux/logger"
)

type DataSource interface {
	GetData() (*CPUInfo, error)
}

type Collector struct {
	Config     *Config
	DataSource DataSource
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
	cpuInfo, err := c.DataSource.GetData()
	if err != nil {
		return nil, err
	}

	metrics := make([]dto.Metric, 0, len(cpuInfo.CPU))

	for _, cpu := range cpuInfo.CPU {
		metrics = append(metrics, dto.Metric{
			Name: MetricMHz,
			Attributes: []dto.MetricAttribute{
				{
					Name:  dto.ResourceAttr,
					Value: ResourceName,
				},
				{
					Name:  SetNameProcessor,
					Value: strconv.Itoa(cpu.ProcessorNumber),
				},
			},
			Value: cpu.MHz,
		})
	}

	return metrics, nil
}
