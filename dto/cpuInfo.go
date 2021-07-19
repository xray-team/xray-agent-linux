package dto

// CPUInfo describes the content of the file /proc/cpuinfo
type CPUInfo struct {
	CPU []CPU
}

type CPU struct {
	ProcessorNumber int     // processor
	VendorID        string  // vendor_id
	ModelName       string  // model name
	MHz             float64 // cpu MHz
	CacheSize       string  // cache size
	PhysicalID      int     // physical id
	CoreID          int     // core id
	CPUCores        int     // cpu cores
	Bugs            string  // info about cpu vulnerabilities such as meltdown, spectre
}
