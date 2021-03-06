package diskSpace

import (
	"fmt"
	"strings"

	"golang.org/x/sys/unix"

	"github.com/xray-team/xray-agent-linux/dto"
	"github.com/xray-team/xray-agent-linux/logger"
)

type MountsDataSource interface {
	GetData() ([]Mounts, error)
}

type Collector struct {
	Config           *Config
	MountsDataSource MountsDataSource
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

	return NewCollector(
		config,
		NewMountsDataSource(MountsPath, CollectorName),
	)
}

// NewCollector returns a new collector object.
func NewCollector(config *Config, mountsDataSource MountsDataSource) dto.Collector {
	if config == nil || mountsDataSource == nil {
		logger.Log.Error.Printf(logger.MessageInitCollectorError, CollectorName)

		return nil
	}

	// exit if collector disabled
	if !config.Enabled {
		return nil
	}

	return &Collector{
		Config:           config,
		MountsDataSource: mountsDataSource,
	}
}

// GetName returns the collector's name.
func (c *Collector) GetName() string {
	return CollectorName
}

// Collect collects and returns metrics.
func (c *Collector) Collect() ([]dto.Metric, error) {
	mounts, err := c.MountsDataSource.GetData()
	if err != nil {
		return nil, fmt.Errorf("cannot parse mounts file: %s", err)
	}

	// Filter mounts
	mounts = c.filterMounts(mounts)

	// Slice for results
	metrics := make([]dto.Metric, 0, len(mounts)*7)

	for _, mount := range mounts {
		attrs := []dto.MetricAttribute{
			{
				Name:  dto.ResourceAttr,
				Value: ResourceName,
			},
			{
				Name:  SetNameMountPoint,
				Value: rewriteMount(mount.MountPoint),
			},
		}

		diskSpace, err := c.getDiskSpaceUsage(mount.MountPoint)
		if err != nil {
			return nil, err
		}

		metrics = append(metrics, genMetricsDiskSpace(attrs, diskSpace)...)
	}

	return metrics, nil
}

func (c *Collector) filterMounts(mounts []Mounts) []Mounts {
	out := make([]Mounts, 0)

	// FS type
	for _, mount := range mounts {
		for _, fs := range c.Config.MonitoredFileSystemTypes {
			if mount.FileSystemType == fs {
				out = append(out, mount)
			}
		}
	}

	return out
}

// getDiskSpaceUsage returns disk usage info and error if any.
func (c *Collector) getDiskSpaceUsage(path string) (*DiskSpaceUsage, error) {
	var (
		out  DiskSpaceUsage
		stat unix.Statfs_t
	)

	if err := unix.Statfs(path, &stat); err != nil {
		return nil, err
	}

	// bytes
	out.Bytes.Available = uint64(stat.Bsize) * stat.Bavail
	out.Bytes.Free = uint64(stat.Bsize) * stat.Bfree
	out.Bytes.Used = uint64(stat.Bsize) * (stat.Blocks - stat.Bfree)
	out.Bytes.Total = uint64(stat.Bsize) * stat.Blocks
	// inodes
	out.Inodes.Free = stat.Ffree
	out.Inodes.Used = stat.Files - stat.Ffree
	out.Inodes.Total = stat.Files

	return &out, nil
}

func rewriteMount(mount string) string {
	mount = strings.ReplaceAll(mount, "/", "_")
	if mount == "_" {
		return "root"
	}

	return fmt.Sprintf("root%s", mount)
}

func genMetricsDiskSpace(attrs []dto.MetricAttribute, diskSpace *DiskSpaceUsage) []dto.Metric {
	return []dto.Metric{
		{
			Name:       MetricBytesAvailable,
			Value:      diskSpace.Bytes.Available,
			Attributes: attrs,
		},
		{
			Name:       MetricBytesFree,
			Value:      diskSpace.Bytes.Free,
			Attributes: attrs,
		},
		{
			Name:       MetricBytesFreePercent,
			Value:      calculateDiskFreePercentage(diskSpace.Bytes),
			Attributes: attrs,
		},
		{
			Name:       MetricBytesUsed,
			Value:      diskSpace.Bytes.Used,
			Attributes: attrs,
		},
		{
			Name:       MetricBytesTotal,
			Value:      diskSpace.Bytes.Total,
			Attributes: attrs,
		},
		{
			Name:       MetricInodesFree,
			Value:      diskSpace.Inodes.Free,
			Attributes: attrs,
		},
		{
			Name:       MetricInodesFreePercent,
			Value:      calculateDiskFreeInodesPercentage(diskSpace.Inodes),
			Attributes: attrs,
		},
		{
			Name:       MetricInodesUsed,
			Value:      diskSpace.Inodes.Used,
			Attributes: attrs,
		},
		{
			Name:       MetricInodesTotal,
			Value:      diskSpace.Inodes.Total,
			Attributes: attrs,
		},
	}
}

func calculateDiskFreePercentage(info DiskSpaceBlockInfo) float64 {
	// prevent division by zero
	if float64(info.Used+info.Available) == 0 {
		return 0
	}

	return float64(info.Available*100) / float64(info.Used+info.Available)
}

func calculateDiskFreeInodesPercentage(info DiskSpaceInodeInfo) float64 {
	// prevent division by zero
	if float64(info.Total) == 0 {
		return 0
	}

	return float64(info.Free*100) / float64(info.Total)
}
