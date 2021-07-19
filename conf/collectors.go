package conf

import "xray-agent-linux/dto"

type CollectorsConf struct {
	RootPath          string            `json:"rootPath" validate:"required,startswith=/"`
	EnableSelfMetrics bool              `json:"enableSelfMetrics"`
	Uptime            *UptimeConf       `json:"uptime"`
	LoadAvg           *LoadAvgConf      `json:"loadAvg"`
	PS                *PSConf           `json:"ps"`
	PSStat            *PSStatConf       `json:"psStat"`
	Stat              *StatConf         `json:"stat"`
	CPUInfo           *CPUInfoConf      `json:"cpuInfo"`
	MemoryInfo        *MemoryInfoConf   `json:"memoryInfo"`
	DiskStat          *DiskStatConf     `json:"diskStat"`
	DiskSpace         *DiskSpaceConf    `json:"diskSpace"`
	NetDev            *NetDevConf       `json:"netDev"`
	NetDevStatus      *NetDevStatusConf `json:"netDevStatus"`
	NetARP            *NetARPConf       `json:"netARP"`
	NetStat           *NetStatConf      `json:"netStat"`
	NetSNMP           *NetSNMPConf      `json:"netSNMP"`
	NetSNMP6          *NetSNMP6Conf     `json:"netSNMP6"`
	MDStat            *MDStatConf       `json:"mdStat"`
	CMD               *CMDConf          `json:"cmd"`
	NginxStubStatus   *NginxStubStatus  `json:"nginxStubStatus"`
	Wireless          *WirelessConf     `json:"wireless"`
}

type UptimeConf struct {
	Enabled bool `json:"enabled"`
}

type LoadAvgConf struct {
	Enabled bool `json:"enabled"`
}

type PSConf struct {
	Enabled bool `json:"enabled"`
}

type PSStatConf struct {
	Enabled     bool     `json:"enabled"`
	ProcessList []string `json:"processList"`
}

type StatConf struct {
	Enabled bool `json:"enabled"`
}

type CPUInfoConf struct {
	Enabled bool `json:"enabled"`
}

type MemoryInfoConf struct {
	Enabled bool `json:"enabled"`
}

type DiskStatConf struct {
	Enabled bool `json:"enabled"`
	// MonitoredDiskTypes - block dev major numbers https://www.kernel.org/doc/Documentation/admin-guide/devices.txt
	MonitoredDiskTypes []int64  `json:"diskTypes" validate:"required,dive,min=0,max=259"`
	ExcludePartitions  bool     `json:"excludePartitions"`
	ExcludeByName      []string `json:"excludeByName"`
}

type DiskSpaceConf struct {
	Enabled bool `json:"enabled"`
	// MonitoredFileSystemTypes is used by procMounts
	MonitoredFileSystemTypes []string `json:"fsTypes" validate:"dive,oneof=ext4 ext3 ext2 btrfs xfs jfs ufs zfs vfat squashfs fuseblk ntfs msdos hfs hfsplus"`
}

type NetDevConf struct {
	Enabled          bool `json:"enabled"`
	ExcludeLoopbacks bool `json:"excludeLoopbacks"`
	ExcludeWireless  bool `json:"excludeWireless"`
	ExcludeBridges   bool `json:"excludeBridges"`
	// Virtual interfaces are network interfaces that are not associated with an any physical interface.
	// Virtual interface examples: loopback, bridge, tun, vlan, ...
	ExcludeVirtual     bool     `json:"excludeVirtual"`
	ExcludeByName      []string `json:"excludeByName"`
	ExcludeByOperState []string `json:"excludeByOperState" validate:"dive,oneof=unknown notpresent down lowerlayerdown testing dormant up"`
}

type NetDevStatusConf struct {
	Enabled         bool     `json:"enabled"`
	ExcludeWireless bool     `json:"excludeWireless"`
	ExcludeByName   []string `json:"excludeByName"`
}

type WirelessConf struct {
	Enabled            bool     `json:"enabled"`
	ExcludeByName      []string `json:"excludeByName"`
	ExcludeByOperState []string `json:"excludeByOperState" validate:"dive,oneof=unknown notpresent down lowerlayerdown testing dormant up"`
}

type NetARPConf struct {
	Enabled bool `json:"enabled"`
}

type NetStatConf struct {
	Enabled bool `json:"enabled"`
}

type NetSNMPConf struct {
	Enabled bool `json:"enabled"`
}

type NetSNMP6Conf struct {
	Enabled bool `json:"enabled"`
}

type MDStatConf struct {
	Enabled bool `json:"enabled"`
}

type CMDConf struct {
	Enabled bool            `json:"enabled"`
	Timeout int             `json:"timeout" validate:"required,min=1,max=120"`
	Metrics []CMDMetricConf `json:"metrics" validate:"dive"`
}

type CMDMetricConf struct {
	Names      []string              `json:"names" validate:"required,min=1"`
	Delimiter  string                `json:"delimiter" validate:"required"`
	Attributes []dto.MetricAttribute `json:"attributes" validate:"dive"`
	PipeLine   [][]string            `json:"pipeline" validate:"required,min=1,dive,required,min=1"`
}

type NginxStubStatus struct {
	Enabled  bool   `json:"enabled"`
	Endpoint string `json:"endpoint" validate:"required"`
	Timeout  int    `json:"timeout" validate:"required,min=1,max=120"`
}
