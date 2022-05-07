package entropy

const (
	CollectorName = "Entropy"
	EntropyPath   = "/sys/kernel/random/entropy_avail"
)

// Metrics
const (
	ResourceName           = "Entropy"
	MetricEntropyAvailable = "Available"
)
