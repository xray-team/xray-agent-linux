package diskSpace

const (
	CollectorName = "DiskSpace"
	MountsPath    = "/proc/mounts"
)

// Metrics
const (
	ResourceName      = "DiskSpace"
	SetNameMountPoint = "MountPoint"

	MetricBytesAvailable   = "BytesAvailable"
	MetricBytesFree        = "BytesFree"
	MetricBytesFreePercent = "BytesFreePercent"
	MetricBytesUsed        = "BytesUsed"
	MetricBytesTotal       = "BytesTotal"
	MetricInodesFree       = "InodesFree"
	MetricInodesUsed       = "InodesUsed"
	MetricInodesTotal      = "InodesTotal"
)