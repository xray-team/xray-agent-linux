package dto

// Stat partially describes the content of the file /proc/stat
type Stat struct {
	CPU              CPUStats
	PerCPU           map[string]CPUStats
	Intr             uint64 // counts of interrupts serviced since boot time (the total of all interrupts serviced)
	Ctxt             uint64 // the total number of context switches across all CPUs
	Btime            uint64 // the time at which the system booted, in seconds since the Unix epoch
	Processes        uint64 // number of forks since boot
	ProcessesRunning uint64 // number of processes currently running
	ProcessesBlocked uint64 // number of processes currently blocked. Blocked are processes that are willing to enter the CPU, but they cannot, e.g. because they wait for disk activity.
	SoftIRQ          SoftIRQStat
}

type CPUStats struct {
	User      uint64 // normal processes executing in user mode
	Nice      uint64 // niced processes executing in user mode
	System    uint64 // processes executing in kernel mode
	Idle      uint64 // time spent in the idle task
	IOwait    uint64 // waiting for I/O to complete (since Linux 2.5.41)
	IRQ       uint64 // servicing interrupts (since Linux 2.6.0)
	SoftIRQ   uint64 // servicing softirqs (since Linux 2.6.0
	Steal     uint64 // stolen time, which is the time spent in other operating systems when running in a virtualized environment (since Linux 2.6.11)
	Guest     uint64 // time spent running a virtual CPU for guest operating systems under the control of the Linux	kernel (since Linux 2.6.24)
	GuestNice uint64 // time spent running a niced guest (virtual CPU for guest operating systems under the control of the Linux kernel) (since Linux 2.6.33)
}

type SoftIRQStat struct {
	// /proc/softirqs - per cpu statistics
	Total   uint64
	Hi      uint64 // HI
	Timer   uint64 // TIMER
	NetTx   uint64 // NET_TX
	NetRx   uint64 // NET_RX
	Block   uint64 // BLOCK
	IRQPoll uint64 // IRQ_POLL
	Tasklet uint64 // TASKLET
	Sched   uint64 // SCHED
	HRTimer uint64 // HRTIMER
	RCU     uint64 // RCU
}
