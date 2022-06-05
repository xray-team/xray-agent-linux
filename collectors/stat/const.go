package stat

const (
	CollectorName = "Stat"
	StatPath      = "/proc/stat"
)

// Metrics
const (
	ResourceName = "CPU"

	// CPU (stat)
	SetNameProcessor       = "CPU"
	SetValueProcessorTotal = "Total"
	SetNameCPUSet          = "Set"
	MetricCtxt             = "Ctxt"
	MetricIntr             = "Intr"
	MetricProcesses        = "Processes"
	MetricProcessesRunning = "ProcessesRunning"
	MetricProcessesBlocked = "ProcessesBlocked"
	MetricBootTime         = "BootTime"
	// CPU Usage
	SetValueCPUSetUsage     = "Usage"
	MetricCPUUsageTotal     = "Total"
	MetricCPUUsageUser      = "User"
	MetricCPUUsageSystem    = "System"
	MetricCPUUsageNice      = "Nice"
	MetricCPUUsageIdle      = "Idle"
	MetricCPUUsageIOwait    = "IOwait"
	MetricCPUUsageIRQ       = "IRQ"
	MetricCPUUsageSoftIRQ   = "SoftIRQ"
	MetricCPUUsageSteal     = "Steal"
	MetricCPUUsageGuest     = "Guest"
	MetricCPUUsageGuestNice = "GuestNice"
	// SoftIRQ
	SetValueCPUSetSoftIRQ = "SoftIRQ"
	MetricSoftIRQTotal    = "Total"
	MetricSoftIRQHi       = "Hi"
	MetricSoftIRQTimer    = "Timer"
	MetricSoftIRQNetTx    = "NetTx"
	MetricSoftIRQNetRx    = "NetRx"
	MetricSoftIRQBlock    = "Block"
	MetricSoftIRQIRQPoll  = "IRQPoll"
	MetricSoftIRQTasklet  = "Tasklet"
	MetricSoftIRQSched    = "Sched"
	MetricSoftIRQHRTimer  = "HRTimer"
	MetricSoftIRQRCU      = "RCU"
	// CPU Count
	SetValueCPUSetCount = "Count"
	MetricCountCPUs     = "CPUs"
)
