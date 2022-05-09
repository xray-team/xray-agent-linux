package diskStat

type Config struct {
	Enabled bool `json:"enabled"`
	// MonitoredDiskTypes - block dev major numbers https://www.kernel.org/doc/Documentation/admin-guide/devices.txt
	MonitoredDiskTypes []int64  `json:"diskTypes" validate:"required,dive,min=0,max=259"`
	ExcludePartitions  bool     `json:"excludePartitions"`
	ExcludeByName      []string `json:"excludeByName"`
}
