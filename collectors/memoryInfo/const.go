package memoryInfo

const (
	CollectorName = "MemoryInfo"
	MemInfoPath   = "/proc/meminfo"
)

// Metrics
const (
	ResourceName = "Memory"

	MetricMemoryUsed      = "Used"
	MetricMemoryTotal     = "Total"
	MetricMemoryFree      = "Free"
	MetricMemoryAvailable = "Available"
	MetricMemoryBuffers   = "Buffers"
	MetricMemoryCached    = "Cached"
	MetricMemorySwapTotal = "SwapTotal"
	MetricMemorySwapFree  = "SwapFree"
)
