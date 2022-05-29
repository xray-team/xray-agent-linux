package ps

const (
	CollectorName = "PS"

	ProcPath     = "/proc"
	PIDsLimit    = "/sys/kernel/pid_max"
	ThreadsLimit = "/sys/kernel/threads-max"
)

// Metrics
const (
	ResourceName = "Processes"

	MetricProcessesCount            = "Count"
	MetricProcessesLimit            = "Limit"
	MetricProcessesInStateRunning   = "InStateRunning"
	MetricProcessesInStateSleeping  = "InStateSleeping"
	MetricProcessesInStateDiskSleep = "InStateDiskSleep"
	MetricProcessesInStateIdle      = "InStateIdle"
	MetricProcessesInStateStopped   = "InStateStopped"
	MetricProcessesInStateZombie    = "InStateZombie"
	MetricProcessesInStateDead      = "InStateDead"
	MetricProcessesThreadsCount     = "ThreadsCount"
	MetricProcessesThreadsLimit     = "ThreadsLimit"
)
