package collectors

import (
	"fmt"

	"github.com/xray-team/xray-agent-linux/conf"
	"github.com/xray-team/xray-agent-linux/dto"
	"github.com/xray-team/xray-agent-linux/logger"
)

type DiskStatDataSource interface {
	GetData() ([]dto.DiskStat, error)
}

type ClassBlockDataSource interface {
	GetData() (map[string]dto.ClassBlock, error)
}

type DiskStatCollector struct {
	Config               *conf.DiskStatConf
	DataSource           DiskStatDataSource
	ClassBlockDataSource ClassBlockDataSource
}

// NewDiskStatCollector returns a new collector object.
func NewDiskStatCollector(cfg *conf.CollectorsConf, diskStatDataSource DiskStatDataSource, classBlockDataSource ClassBlockDataSource) dto.Collector {
	if cfg == nil || diskStatDataSource == nil || classBlockDataSource == nil {
		logger.LogWarning(logger.CollectorInitPrefix, fmt.Errorf("%s collector init params error", dto.CollectorNameDiskStat))
		return nil
	}

	// exit if collector disabled
	if cfg.DiskStat == nil || !cfg.DiskStat.Enabled {
		return nil
	}

	return &DiskStatCollector{
		Config:               cfg.DiskStat,
		DataSource:           diskStatDataSource,
		ClassBlockDataSource: classBlockDataSource,
	}
}

// GetName returns the collector's name.
func (c *DiskStatCollector) GetName() string {
	return dto.CollectorNameDiskStat
}

// Collect collects and returns metrics.
func (c *DiskStatCollector) Collect() ([]dto.Metric, error) {
	// Block Dev Inventory
	inventory, err := c.ClassBlockDataSource.GetData()
	if err != nil {
		return nil, err
	}

	// Applying Filters
	inventory = c.filterBlockDevByMajor(inventory)
	inventory = c.excludeBlockDevPartitions(inventory)
	inventory = c.excludeBlockDevByName(inventory)

	diskStats, err := c.DataSource.GetData()
	if err != nil {
		return nil, err
	}

	// Slice for results
	metrics := make([]dto.Metric, 0, len(inventory)*8)

	// fill out
	for _, ds := range diskStats {
		stat := ds

		// Skip if blockDev not found in inventory
		_, ok := inventory[stat.Dev]
		if !ok {
			continue
		}

		attrs := []dto.MetricAttribute{
			{
				Name:  dto.ResourceAttr,
				Value: dto.ResourceDiskStat,
			},
			{
				Name:  dto.SetNameDiskStatDev,
				Value: stat.Dev,
			},
		}
		metrics = append(metrics, genMetricsDiskStat(attrs, stat)...)
	}

	return metrics, nil
}

func genMetricsDiskStat(attrs []dto.MetricAttribute, diskStat dto.DiskStat) []dto.Metric {
	return []dto.Metric{
		{
			Name:       dto.MetricDiskStatReadsCompletedSuccessfully,
			Value:      diskStat.ReadsCompletedSuccessfully,
			Attributes: attrs,
		},
		{
			Name:       dto.MetricDiskStatReadsMerged,
			Value:      diskStat.ReadsMerged,
			Attributes: attrs,
		},
		{
			Name:       dto.MetricDiskStatSectorsRead,
			Value:      diskStat.SectorsRead,
			Attributes: attrs,
		},
		{
			Name:       dto.MetricDiskStatTimeSpentReading,
			Value:      diskStat.TimeSpentReading,
			Attributes: attrs,
		},
		{
			Name:       dto.MetricDiskStatWritesCompleted,
			Value:      diskStat.WritesCompleted,
			Attributes: attrs,
		},
		{
			Name:       dto.MetricDiskStatWritesMerged,
			Value:      diskStat.WritesMerged,
			Attributes: attrs,
		},
		{
			Name:       dto.MetricDiskStatSectorsWritten,
			Value:      diskStat.SectorsWritten,
			Attributes: attrs,
		},
		{
			Name:       dto.MetricDiskStatTimeSpentWriting,
			Value:      diskStat.TimeSpentWriting,
			Attributes: attrs,
		},
		{
			Name:       dto.MetricDiskStatIOsCurrentlyInProgress,
			Value:      diskStat.IOsCurrentlyInProgress,
			Attributes: attrs,
		},
		{
			Name:       dto.MetricDiskStatTimeSpentDoingIOs,
			Value:      diskStat.TimeSpentDoingIOs,
			Attributes: attrs,
		},
		{
			Name:       dto.MetricDiskStatWeightedTimeSpentDoingIOs,
			Value:      diskStat.WeightedTimeSpentDoingIOs,
			Attributes: attrs,
		},
		{
			Name:       dto.MetricDiskStatDiscardsCompletedSuccessfully,
			Value:      diskStat.DiscardsCompletedSuccessfully,
			Attributes: attrs,
		},
		{
			Name:       dto.MetricDiskStatDiscardsMerged,
			Value:      diskStat.DiscardsMerged,
			Attributes: attrs,
		},
		{
			Name:       dto.MetricDiskStatSectorsDiscarded,
			Value:      diskStat.SectorsDiscarded,
			Attributes: attrs,
		},
		{
			Name:       dto.MetricDiskStatTimeSpentDiscarding,
			Value:      diskStat.TimeSpentDiscarding,
			Attributes: attrs,
		},
		{
			Name:       dto.MetricDiskStatFlushRequestsCompletedSuccessfully,
			Value:      diskStat.FlushRequestsCompletedSuccessfully,
			Attributes: attrs,
		},
		{
			Name:       dto.MetricDiskStatTimeSpentFlushing,
			Value:      diskStat.TimeSpentFlushing,
			Attributes: attrs,
		},
	}
}

func (c *DiskStatCollector) filterBlockDevByMajor(m map[string]dto.ClassBlock) map[string]dto.ClassBlock {
	out := make(map[string]dto.ClassBlock)

	for devName, dev := range m {
		for _, major := range c.Config.MonitoredDiskTypes {
			if dev.Uevent.Major == major {
				out[devName] = dev
			}
		}
	}

	return out
}

func (c *DiskStatCollector) excludeBlockDevByName(m map[string]dto.ClassBlock) map[string]dto.ClassBlock {
	for _, devName := range c.Config.ExcludeByName {
		delete(m, devName)
	}

	return m
}

func (c *DiskStatCollector) excludeBlockDevPartitions(m map[string]dto.ClassBlock) map[string]dto.ClassBlock {
	if !c.Config.ExcludePartitions {
		return m
	}

	out := make(map[string]dto.ClassBlock)

	for devName, dev := range m {
		if dev.Uevent.DevType != dto.BlockDevTypePartition {
			out[devName] = dev
		}
	}

	return out
}
