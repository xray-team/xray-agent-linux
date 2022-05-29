package entropy

const (
	CollectorName = "Entropy"
	EntropyPath   = "/proc/sys/kernel/random/entropy_avail"
)

// Metrics
const (
	ResourceName           = "Entropy"
	MetricEntropyAvailable = "Available"
)
