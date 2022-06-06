package stat

import (
	"github.com/xray-team/xray-agent-linux/dto"
	"github.com/xray-team/xray-agent-linux/logger"
)

type DataSource interface {
	GetData() (*Stat, error)
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
		NewDataSource(StatPath, CollectorName),
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
	stat, err := c.DataSource.GetData()
	if err != nil {
		return nil, err
	}

	metrics := make([]dto.Metric, 0, (len(stat.PerCPU))*11+25)

	// CPU Total
	resourceAttr := dto.MetricAttribute{
		Name:  dto.ResourceAttr,
		Value: ResourceName,
	}

	totalAttr := dto.MetricAttribute{
		Name:  SetNameProcessor,
		Value: SetValueProcessorTotal,
	}

	usageAttr := dto.MetricAttribute{
		Name:  SetNameCPUSet,
		Value: SetValueCPUSetUsage,
	}

	softIRQAttr := dto.MetricAttribute{
		Name:  SetNameCPUSet,
		Value: SetValueCPUSetSoftIRQ,
	}

	countAttr := dto.MetricAttribute{
		Name:  SetNameCPUSet,
		Value: SetValueCPUSetCount,
	}

	metrics = append(metrics,
		// Ctxt
		dto.Metric{
			Name:       MetricCtxt,
			Value:      stat.Ctxt,
			Attributes: []dto.MetricAttribute{resourceAttr, totalAttr},
		},
		// Intr
		dto.Metric{
			Name:       MetricIntr,
			Value:      stat.Intr,
			Attributes: []dto.MetricAttribute{resourceAttr, totalAttr},
		},
		// Processes
		dto.Metric{
			Name:       MetricProcesses,
			Value:      stat.Processes,
			Attributes: []dto.MetricAttribute{resourceAttr, totalAttr},
		},
		// ProcessesRunning (procs_running)
		dto.Metric{
			Name:       MetricProcessesRunning,
			Value:      stat.ProcessesRunning,
			Attributes: []dto.MetricAttribute{resourceAttr, totalAttr},
		},
		// ProcessesBlocked (procs_blocked)
		dto.Metric{
			Name:       MetricProcessesBlocked,
			Value:      stat.ProcessesBlocked,
			Attributes: []dto.MetricAttribute{resourceAttr, totalAttr},
		},
		// BootTime (btime)
		dto.Metric{
			Name:       MetricBootTime,
			Value:      stat.Btime,
			Attributes: []dto.MetricAttribute{resourceAttr, totalAttr},
		},
		// CPU Usage
		dto.Metric{
			Name:       MetricCPUUsageTotal,
			Value:      calculateTotalCPUUsage(stat.CPU),
			Attributes: []dto.MetricAttribute{resourceAttr, totalAttr, usageAttr},
		},
		dto.Metric{
			Name:       MetricCPUUsageUser,
			Value:      stat.CPU.User,
			Attributes: []dto.MetricAttribute{resourceAttr, totalAttr, usageAttr},
		},
		dto.Metric{
			Name:       MetricCPUUsageSystem,
			Value:      stat.CPU.System,
			Attributes: []dto.MetricAttribute{resourceAttr, totalAttr, usageAttr},
		},
		dto.Metric{
			Name:       MetricCPUUsageNice,
			Value:      stat.CPU.Nice,
			Attributes: []dto.MetricAttribute{resourceAttr, totalAttr, usageAttr},
		},
		dto.Metric{
			Name:       MetricCPUUsageIdle,
			Value:      stat.CPU.Idle,
			Attributes: []dto.MetricAttribute{resourceAttr, totalAttr, usageAttr},
		},
		dto.Metric{
			Name:       MetricCPUUsageIOwait,
			Value:      stat.CPU.IOwait,
			Attributes: []dto.MetricAttribute{resourceAttr, totalAttr, usageAttr},
		},
		dto.Metric{
			Name:       MetricCPUUsageGuest,
			Value:      stat.CPU.Guest,
			Attributes: []dto.MetricAttribute{resourceAttr, totalAttr, usageAttr},
		},
		dto.Metric{
			Name:       MetricCPUUsageGuestNice,
			Value:      stat.CPU.GuestNice,
			Attributes: []dto.MetricAttribute{resourceAttr, totalAttr, usageAttr},
		},
		dto.Metric{
			Name:       MetricCPUUsageIRQ,
			Value:      stat.CPU.IRQ,
			Attributes: []dto.MetricAttribute{resourceAttr, totalAttr, usageAttr},
		},
		dto.Metric{
			Name:       MetricCPUUsageSoftIRQ,
			Value:      stat.CPU.SoftIRQ,
			Attributes: []dto.MetricAttribute{resourceAttr, totalAttr, usageAttr},
		},
		dto.Metric{
			Name:       MetricCPUUsageSteal,
			Value:      stat.CPU.Steal,
			Attributes: []dto.MetricAttribute{resourceAttr, totalAttr, usageAttr},
		},
		// SoftIRQ
		dto.Metric{
			Name:       MetricSoftIRQTotal,
			Value:      stat.SoftIRQ.Total,
			Attributes: []dto.MetricAttribute{resourceAttr, totalAttr, softIRQAttr},
		},
		dto.Metric{
			Name:       MetricSoftIRQHi,
			Value:      stat.SoftIRQ.Hi,
			Attributes: []dto.MetricAttribute{resourceAttr, totalAttr, softIRQAttr},
		},
		dto.Metric{
			Name:       MetricSoftIRQTimer,
			Value:      stat.SoftIRQ.Timer,
			Attributes: []dto.MetricAttribute{resourceAttr, totalAttr, softIRQAttr},
		},
		dto.Metric{
			Name:       MetricSoftIRQNetRx,
			Value:      stat.SoftIRQ.NetRx,
			Attributes: []dto.MetricAttribute{resourceAttr, totalAttr, softIRQAttr},
		},
		dto.Metric{
			Name:       MetricSoftIRQNetTx,
			Value:      stat.SoftIRQ.NetTx,
			Attributes: []dto.MetricAttribute{resourceAttr, totalAttr, softIRQAttr},
		},
		dto.Metric{
			Name:       MetricSoftIRQBlock,
			Value:      stat.SoftIRQ.Block,
			Attributes: []dto.MetricAttribute{resourceAttr, totalAttr, softIRQAttr},
		},
		dto.Metric{
			Name:       MetricSoftIRQIRQPoll,
			Value:      stat.SoftIRQ.IRQPoll,
			Attributes: []dto.MetricAttribute{resourceAttr, totalAttr, softIRQAttr},
		},
		dto.Metric{
			Name:       MetricSoftIRQTasklet,
			Value:      stat.SoftIRQ.Tasklet,
			Attributes: []dto.MetricAttribute{resourceAttr, totalAttr, softIRQAttr},
		},
		dto.Metric{
			Name:       MetricSoftIRQSched,
			Value:      stat.SoftIRQ.Sched,
			Attributes: []dto.MetricAttribute{resourceAttr, totalAttr, softIRQAttr},
		},
		dto.Metric{
			Name:       MetricSoftIRQHRTimer,
			Value:      stat.SoftIRQ.HRTimer,
			Attributes: []dto.MetricAttribute{resourceAttr, totalAttr, softIRQAttr},
		},
		dto.Metric{
			Name:       MetricSoftIRQRCU,
			Value:      stat.SoftIRQ.RCU,
			Attributes: []dto.MetricAttribute{resourceAttr, totalAttr, softIRQAttr},
		})

	// Per CPU
	for cpuNumber, cpuStat := range stat.PerCPU {
		attrs := []dto.MetricAttribute{
			resourceAttr,
			{
				Name:  SetNameProcessor,
				Value: cpuNumber,
			},
			usageAttr,
		}

		metrics = append(metrics,
			dto.Metric{
				Name:       MetricCPUUsageTotal,
				Value:      calculateTotalCPUUsage(cpuStat),
				Attributes: attrs,
			},
			dto.Metric{
				Name:       MetricCPUUsageUser,
				Value:      cpuStat.User,
				Attributes: attrs,
			},
			dto.Metric{
				Name:       MetricCPUUsageSystem,
				Value:      cpuStat.System,
				Attributes: attrs,
			},
			dto.Metric{
				Name:       MetricCPUUsageNice,
				Value:      cpuStat.Nice,
				Attributes: attrs,
			},
			dto.Metric{
				Name:       MetricCPUUsageIdle,
				Value:      cpuStat.Idle,
				Attributes: attrs,
			},
			dto.Metric{
				Name:       MetricCPUUsageIOwait,
				Value:      cpuStat.IOwait,
				Attributes: attrs,
			},
			dto.Metric{
				Name:       MetricCPUUsageGuest,
				Value:      cpuStat.Guest,
				Attributes: attrs,
			},
			dto.Metric{
				Name:       MetricCPUUsageGuestNice,
				Value:      cpuStat.GuestNice,
				Attributes: attrs,
			},
			dto.Metric{
				Name:       MetricCPUUsageIRQ,
				Value:      cpuStat.IRQ,
				Attributes: attrs,
			},
			dto.Metric{
				Name:       MetricCPUUsageSoftIRQ,
				Value:      cpuStat.SoftIRQ,
				Attributes: attrs,
			},
			dto.Metric{
				Name:       MetricCPUUsageSteal,
				Value:      cpuStat.Steal,
				Attributes: attrs,
			},
		)
	}

	// Count.CPUs
	metrics = append(metrics,
		// Ctxt
		dto.Metric{
			Name:       MetricCountCPUs,
			Value:      len(stat.PerCPU),
			Attributes: []dto.MetricAttribute{resourceAttr, countAttr},
		},
	)

	return metrics, nil
}

// calculateTotalCPUUsage - Calculates Total CPU Usage
func calculateTotalCPUUsage(stats CPUStats) uint64 {
	return stats.User +
		stats.GuestNice +
		stats.Guest +
		stats.Steal +
		stats.SoftIRQ +
		stats.IOwait +
		stats.Nice +
		stats.System +
		stats.IRQ
}
