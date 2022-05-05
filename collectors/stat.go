package collectors

import (
	"github.com/xray-team/xray-agent-linux/conf"
	"github.com/xray-team/xray-agent-linux/dto"
	"github.com/xray-team/xray-agent-linux/logger"
)

type StatDataSource interface {
	GetData() (*dto.Stat, error)
}

type StatCollector struct {
	Config     *conf.StatConf
	DataSource StatDataSource
}

// NewStatCollector returns a new collector object.
func NewStatCollector(cfg *conf.CollectorsConf, dataSource StatDataSource) dto.Collector {
	if cfg == nil || dataSource == nil {
		logger.Log.Error.Printf(logger.MessageInitCollectorError, dto.CollectorNameStat)

		return nil
	}

	// exit if collector disabled
	if cfg.Stat == nil || !cfg.Stat.Enabled {
		return nil
	}

	return &StatCollector{
		Config:     cfg.Stat,
		DataSource: dataSource,
	}
}

// GetName returns the collector's name.
func (c *StatCollector) GetName() string {
	return dto.CollectorNameStat
}

// Collect collects and returns metrics.
func (c *StatCollector) Collect() ([]dto.Metric, error) {
	stat, err := c.DataSource.GetData()
	if err != nil {
		return nil, err
	}

	metrics := make([]dto.Metric, 0, (len(stat.PerCPU))*11+25)

	// CPU Total
	resourceAttr := dto.MetricAttribute{
		Name:  dto.ResourceAttr,
		Value: dto.ResourceStat,
	}

	totalAttr := dto.MetricAttribute{
		Name:  dto.SetNameCPUProcessor,
		Value: dto.SetValueCPUProcessorTotal,
	}

	usageAttr := dto.MetricAttribute{
		Name:  dto.SetNameCPUSet,
		Value: dto.SetValueCPUSetUsage,
	}

	softIRQAttr := dto.MetricAttribute{
		Name:  dto.SetNameCPUSet,
		Value: dto.SetValueCPUSetSoftIRQ,
	}

	countAttr := dto.MetricAttribute{
		Name:  dto.SetNameCPUSet,
		Value: dto.SetValueCPUSetCount,
	}

	metrics = append(metrics,
		// Ctxt
		dto.Metric{
			Name:       dto.MetricCPUCtxt,
			Value:      stat.Ctxt,
			Attributes: []dto.MetricAttribute{resourceAttr, totalAttr},
		},
		// Intr
		dto.Metric{
			Name:       dto.MetricCPUIntr,
			Value:      stat.Intr,
			Attributes: []dto.MetricAttribute{resourceAttr, totalAttr},
		},
		// CPU Usage
		dto.Metric{
			Name:       dto.MetricCPUUsageTotal,
			Value:      calculateTotalCPUUsage(stat.CPU),
			Attributes: []dto.MetricAttribute{resourceAttr, totalAttr, usageAttr},
		},
		dto.Metric{
			Name:       dto.MetricCPUUsageUser,
			Value:      stat.CPU.User,
			Attributes: []dto.MetricAttribute{resourceAttr, totalAttr, usageAttr},
		},
		dto.Metric{
			Name:       dto.MetricCPUUsageSystem,
			Value:      stat.CPU.System,
			Attributes: []dto.MetricAttribute{resourceAttr, totalAttr, usageAttr},
		},
		dto.Metric{
			Name:       dto.MetricCPUUsageNice,
			Value:      stat.CPU.Nice,
			Attributes: []dto.MetricAttribute{resourceAttr, totalAttr, usageAttr},
		},
		dto.Metric{
			Name:       dto.MetricCPUUsageIdle,
			Value:      stat.CPU.Idle,
			Attributes: []dto.MetricAttribute{resourceAttr, totalAttr, usageAttr},
		},
		dto.Metric{
			Name:       dto.MetricCPUUsageIOwait,
			Value:      stat.CPU.IOwait,
			Attributes: []dto.MetricAttribute{resourceAttr, totalAttr, usageAttr},
		},
		dto.Metric{
			Name:       dto.MetricCPUUsageGuest,
			Value:      stat.CPU.Guest,
			Attributes: []dto.MetricAttribute{resourceAttr, totalAttr, usageAttr},
		},
		dto.Metric{
			Name:       dto.MetricCPUUsageGuestNice,
			Value:      stat.CPU.GuestNice,
			Attributes: []dto.MetricAttribute{resourceAttr, totalAttr, usageAttr},
		},
		dto.Metric{
			Name:       dto.MetricCPUUsageIRQ,
			Value:      stat.CPU.IRQ,
			Attributes: []dto.MetricAttribute{resourceAttr, totalAttr, usageAttr},
		},
		dto.Metric{
			Name:       dto.MetricCPUUsageSoftIRQ,
			Value:      stat.CPU.SoftIRQ,
			Attributes: []dto.MetricAttribute{resourceAttr, totalAttr, usageAttr},
		},
		dto.Metric{
			Name:       dto.MetricCPUUsageSteal,
			Value:      stat.CPU.Steal,
			Attributes: []dto.MetricAttribute{resourceAttr, totalAttr, usageAttr},
		},
		// SoftIRQ
		dto.Metric{
			Name:       dto.MetricCPUSoftIRQTotal,
			Value:      stat.SoftIRQ.Total,
			Attributes: []dto.MetricAttribute{resourceAttr, totalAttr, softIRQAttr},
		},
		dto.Metric{
			Name:       dto.MetricCPUSoftIRQHi,
			Value:      stat.SoftIRQ.Hi,
			Attributes: []dto.MetricAttribute{resourceAttr, totalAttr, softIRQAttr},
		},
		dto.Metric{
			Name:       dto.MetricCPUSoftIRQTimer,
			Value:      stat.SoftIRQ.Timer,
			Attributes: []dto.MetricAttribute{resourceAttr, totalAttr, softIRQAttr},
		},
		dto.Metric{
			Name:       dto.MetricCPUSoftIRQNetRx,
			Value:      stat.SoftIRQ.NetRx,
			Attributes: []dto.MetricAttribute{resourceAttr, totalAttr, softIRQAttr},
		},
		dto.Metric{
			Name:       dto.MetricCPUSoftIRQNetTx,
			Value:      stat.SoftIRQ.NetTx,
			Attributes: []dto.MetricAttribute{resourceAttr, totalAttr, softIRQAttr},
		},
		dto.Metric{
			Name:       dto.MetricCPUSoftIRQBlock,
			Value:      stat.SoftIRQ.Block,
			Attributes: []dto.MetricAttribute{resourceAttr, totalAttr, softIRQAttr},
		},
		dto.Metric{
			Name:       dto.MetricCPUSoftIRQIRQPoll,
			Value:      stat.SoftIRQ.IRQPoll,
			Attributes: []dto.MetricAttribute{resourceAttr, totalAttr, softIRQAttr},
		},
		dto.Metric{
			Name:       dto.MetricCPUSoftIRQTasklet,
			Value:      stat.SoftIRQ.Tasklet,
			Attributes: []dto.MetricAttribute{resourceAttr, totalAttr, softIRQAttr},
		},
		dto.Metric{
			Name:       dto.MetricCPUSoftIRQSched,
			Value:      stat.SoftIRQ.Sched,
			Attributes: []dto.MetricAttribute{resourceAttr, totalAttr, softIRQAttr},
		},
		dto.Metric{
			Name:       dto.MetricCPUSoftIRQHRTimer,
			Value:      stat.SoftIRQ.HRTimer,
			Attributes: []dto.MetricAttribute{resourceAttr, totalAttr, softIRQAttr},
		},
		dto.Metric{
			Name:       dto.MetricCPUSoftIRQRCU,
			Value:      stat.SoftIRQ.RCU,
			Attributes: []dto.MetricAttribute{resourceAttr, totalAttr, softIRQAttr},
		})

	// Per CPU
	for cpuNumber, cpuStat := range stat.PerCPU {
		attrs := []dto.MetricAttribute{
			resourceAttr,
			{
				Name:  dto.SetNameCPUProcessor,
				Value: cpuNumber,
			},
			usageAttr,
		}

		metrics = append(metrics,
			dto.Metric{
				Name:       dto.MetricCPUUsageTotal,
				Value:      calculateTotalCPUUsage(cpuStat),
				Attributes: attrs,
			},
			dto.Metric{
				Name:       dto.MetricCPUUsageUser,
				Value:      cpuStat.User,
				Attributes: attrs,
			},
			dto.Metric{
				Name:       dto.MetricCPUUsageSystem,
				Value:      cpuStat.System,
				Attributes: attrs,
			},
			dto.Metric{
				Name:       dto.MetricCPUUsageNice,
				Value:      cpuStat.Nice,
				Attributes: attrs,
			},
			dto.Metric{
				Name:       dto.MetricCPUUsageIdle,
				Value:      cpuStat.Idle,
				Attributes: attrs,
			},
			dto.Metric{
				Name:       dto.MetricCPUUsageIOwait,
				Value:      cpuStat.IOwait,
				Attributes: attrs,
			},
			dto.Metric{
				Name:       dto.MetricCPUUsageGuest,
				Value:      cpuStat.Guest,
				Attributes: attrs,
			},
			dto.Metric{
				Name:       dto.MetricCPUUsageGuestNice,
				Value:      cpuStat.GuestNice,
				Attributes: attrs,
			},
			dto.Metric{
				Name:       dto.MetricCPUUsageIRQ,
				Value:      cpuStat.IRQ,
				Attributes: attrs,
			},
			dto.Metric{
				Name:       dto.MetricCPUUsageSoftIRQ,
				Value:      cpuStat.SoftIRQ,
				Attributes: attrs,
			},
			dto.Metric{
				Name:       dto.MetricCPUUsageSteal,
				Value:      cpuStat.Steal,
				Attributes: attrs,
			},
		)
	}

	// Count.CPUs
	metrics = append(metrics,
		// Ctxt
		dto.Metric{
			Name:       dto.MetricCPUCountCPUs,
			Value:      len(stat.PerCPU),
			Attributes: []dto.MetricAttribute{resourceAttr, countAttr},
		},
	)

	return metrics, nil
}

// calculateTotalCPUUsage - Calculates Total CPU Usage
func calculateTotalCPUUsage(stats dto.CPUStats) uint64 {
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
