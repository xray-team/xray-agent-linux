package proc

const (
	// Paths
	ProcPath = "/proc"
	// Relative Paths to proc files
	CPUInfoPath   = "/cpuinfo"
	LoadAvgPath   = "/loadavg"
	StatPath      = "/stat"
	MemInfoPath   = "/meminfo"
	NetArpPath    = "/net/arp"
	NetDevPath    = "/net/dev"
	NetStatPath   = "/net/netstat"
	NetSNMPPath   = "/net/snmp"
	NetSNMP6Path  = "/net/snmp6"
	UptimePath    = "/uptime"
	MountsPath    = "/mounts"
	DiskStatsPath = "/diskstats"
	PIDsLimit     = "/sys/kernel/pid_max"
	ThreadsLimit  = "/sys/kernel/threads-max"
)
