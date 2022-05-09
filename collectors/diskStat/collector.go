package diskStat

import (
	"github.com/xray-team/xray-agent-linux/dto"
	"github.com/xray-team/xray-agent-linux/logger"
)

type DiskStatDataSource interface {
	GetData() ([]DiskStat, error)
}

type ClassBlockDataSource interface {
	GetData() (map[string]dto.ClassBlock, error)
}

type Collector struct {
	Config               *Config
	DataSource           DiskStatDataSource
	ClassBlockDataSource ClassBlockDataSource
}

// NewCollector returns a new collector object.
func NewCollector(config *Config, diskStatDataSource DiskStatDataSource, classBlockDataSource ClassBlockDataSource) dto.Collector {
	if config == nil || diskStatDataSource == nil || classBlockDataSource == nil {
		logger.Log.Error.Printf(logger.MessageInitCollectorError, CollectorName)

		return nil
	}

	// exit if collector disabled
	if !config.Enabled {
		return nil
	}

	return &Collector{
		Config:               config,
		DataSource:           diskStatDataSource,
		ClassBlockDataSource: classBlockDataSource,
	}
}

// GetName returns the collector's name.
func (c *Collector) GetName() string {
	return CollectorName
}

// Collect collects and returns metrics.
func (c *Collector) Collect() ([]dto.Metric, error) {
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
				Value: ResourceName,
			},
			{
				Name:  SetNameDev,
				Value: stat.Dev,
			},
		}
		metrics = append(metrics, genMetricsDiskStat(attrs, stat)...)
	}

	return metrics, nil
}

func genMetricsDiskStat(attrs []dto.MetricAttribute, diskStat DiskStat) []dto.Metric {
	return []dto.Metric{
		{
			Name:       MetricReadsCompletedSuccessfully,
			Value:      diskStat.ReadsCompletedSuccessfully,
			Attributes: attrs,
		},
		{
			Name:       MetricReadsMerged,
			Value:      diskStat.ReadsMerged,
			Attributes: attrs,
		},
		{
			Name:       MetricSectorsRead,
			Value:      diskStat.SectorsRead,
			Attributes: attrs,
		},
		{
			Name:       MetricTimeSpentReading,
			Value:      diskStat.TimeSpentReading,
			Attributes: attrs,
		},
		{
			Name:       MetricWritesCompleted,
			Value:      diskStat.WritesCompleted,
			Attributes: attrs,
		},
		{
			Name:       MetricWritesMerged,
			Value:      diskStat.WritesMerged,
			Attributes: attrs,
		},
		{
			Name:       MetricSectorsWritten,
			Value:      diskStat.SectorsWritten,
			Attributes: attrs,
		},
		{
			Name:       MetricTimeSpentWriting,
			Value:      diskStat.TimeSpentWriting,
			Attributes: attrs,
		},
		{
			Name:       MetricIOsCurrentlyInProgress,
			Value:      diskStat.IOsCurrentlyInProgress,
			Attributes: attrs,
		},
		{
			Name:       MetricTimeSpentDoingIOs,
			Value:      diskStat.TimeSpentDoingIOs,
			Attributes: attrs,
		},
		{
			Name:       MetricWeightedTimeSpentDoingIOs,
			Value:      diskStat.WeightedTimeSpentDoingIOs,
			Attributes: attrs,
		},
		{
			Name:       MetricDiscardsCompletedSuccessfully,
			Value:      diskStat.DiscardsCompletedSuccessfully,
			Attributes: attrs,
		},
		{
			Name:       MetricDiscardsMerged,
			Value:      diskStat.DiscardsMerged,
			Attributes: attrs,
		},
		{
			Name:       MetricSectorsDiscarded,
			Value:      diskStat.SectorsDiscarded,
			Attributes: attrs,
		},
		{
			Name:       MetricTimeSpentDiscarding,
			Value:      diskStat.TimeSpentDiscarding,
			Attributes: attrs,
		},
		{
			Name:       MetricFlushRequestsCompletedSuccessfully,
			Value:      diskStat.FlushRequestsCompletedSuccessfully,
			Attributes: attrs,
		},
		{
			Name:       MetricTimeSpentFlushing,
			Value:      diskStat.TimeSpentFlushing,
			Attributes: attrs,
		},
	}
}

func (c *Collector) filterBlockDevByMajor(m map[string]dto.ClassBlock) map[string]dto.ClassBlock {
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

func (c *Collector) excludeBlockDevByName(m map[string]dto.ClassBlock) map[string]dto.ClassBlock {
	for _, devName := range c.Config.ExcludeByName {
		delete(m, devName)
	}

	return m
}

func (c *Collector) excludeBlockDevPartitions(m map[string]dto.ClassBlock) map[string]dto.ClassBlock {
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
