package stat

const (
	CollectorName = "Stat"
	StatPath      = "/proc/stat"
)

// Metrics
const (
	ResourceName = "CPU"

	// CPU (stat)
	SetNameCPUProcessor       = "CPU"
	SetValueCPUProcessorTotal = "Total"
	SetNameCPUSet             = "Set"
	MetricCPUCtxt             = "Ctxt"
	MetricCPUIntr             = "Intr"
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
	// CPU SoftIRQ
	SetValueCPUSetSoftIRQ   = "SoftIRQ"
	MetricCPUSoftIRQTotal   = "Total"
	MetricCPUSoftIRQHi      = "Hi"
	MetricCPUSoftIRQTimer   = "Timer"
	MetricCPUSoftIRQNetTx   = "NetTx"
	MetricCPUSoftIRQNetRx   = "NetRx"
	MetricCPUSoftIRQBlock   = "Block"
	MetricCPUSoftIRQIRQPoll = "IRQPoll"
	MetricCPUSoftIRQTasklet = "Tasklet"
	MetricCPUSoftIRQSched   = "Sched"
	MetricCPUSoftIRQHRTimer = "HRTimer"
	MetricCPUSoftIRQRCU     = "RCU"
	// CPU Count
	SetValueCPUSetCount = "Count"
	MetricCPUCountCPUs  = "CPUs"
)
