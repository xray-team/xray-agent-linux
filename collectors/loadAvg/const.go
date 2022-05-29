package loadAvg

const (
	CollectorName = "LoadAvg"
	LoadAvgPath   = "/proc/loadavg"
)

// Metrics
const (
	ResourceName = "LoadAvg"

	MetricLast                     = "Last"
	MetricLast5m                   = "Last5m"
	MetricLast15m                  = "Last15m"
	MetricKernelSchedulingEntities = "KernelSchedulingEntities"
)
