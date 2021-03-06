package psStat

const (
	CollectorName = "PSStat"
	ProcPath      = "/proc"
)

// Metrics
const (
	ResourceName = "PSStat"

	SetNameProcessName       = "ProcessName"
	SetNamePID               = "PID"
	SetValuePIDTotal         = "Total"
	MetricUser               = "User"
	MetricSystem             = "System"
	MetricGuest              = "Guest"
	MetricTotal              = "Total"
	MetricThreads            = "Threads"
	MetricProcesses          = "Processes"
	MetricVirtualMemorySize  = "VirtualMemorySize"
	MetricResidentMemorySize = "ResidentMemorySize"
)
