package dto

const (
	// Collectors Names
	CollectorNameLoadAvg      = "LoadAvg"
	CollectorNameUptime       = "Uptime"
	CollectorNamePS           = "PS"
	CollectorNamePSStat       = "PSStat"
	CollectorNameStat         = "Stat"
	CollectorNameCPUInfo      = "CPUInfo"
	CollectorNameMemoryInfo   = "MemoryInfo"
	CollectorNameDiskStat     = "DiskStat"
	CollectorNameDiskSpace    = "DiskSpace"
	CollectorNameNetDev       = "NetDev"
	CollectorNameNetDevStatus = "NetDevStatus"
	CollectorNameWireless     = "Wireless"
	CollectorNameNetARP       = "NetARP"
	CollectorNameNetStat      = "NetStat"
	CollectorNameNetSNMP      = "NetSNMP"
	CollectorNameNetSNMP6     = "NetSNMP6"
	CollectorNameMDStat       = "MDStat"
	CollectorNameCMD          = "CMD"
	CollectorNameNginx        = "Nginx"
	CollectorNameEntropy      = "Entropy"

	// Graphite working modes.
	GraphiteModeTree = "tree"
	GraphiteModeTags = "tags"

	// NetDev Types
	//   "/sys/class/net/<iface>/uevent"
	NetDevTypeGeneric  = "generic"
	NetDevTypeBond     = "bond"
	NetDevTypeBridge   = "bridge"
	NetDevTypeVlan     = "vlan"
	NetDevTypeWireless = "wlan"

	// NetDev Protocol Types
	//  "/sys/class/net/<iface>/type"
	NetDevProtocolTypeEthernet = 1
	NetDevProtocolTypeLoopback = 772

	// BlockDev types
	// "/sys/class/block/<device>/uevent"
	BlockDevTypeDisk      = "disk"
	BlockDevTypePartition = "partition"

	// Major numbers
	//   https://www.kernel.org/doc/Documentation/admin-guide/devices.txt
	BlockDevMajorTypeRAMDisk  = 1
	BlockDevMajorTypeFloppy   = 2
	BlockDevMajorTypeIDE      = 3
	BlockDevMajorTypeLoopback = 7
	BlockDevMajorTypeSCSI     = 8 // SCSI disk devices  (sda*, sdb*, sdc*, ...,)
	BlockDevMajorTypeMD       = 9 // Metadisk (RAID) devices (md0, md1, ...)
)

// RFC2863
var NetDevOperStates = map[string]int{
	"up":             1,
	"lowerlayerdown": 2,
	"dormant":        3,
	"down":           4,
	"unknown":        5,
	"testing":        6,
	"notpresent":     7,
}

var NetDevDuplexStates = map[string]int{
	"full":    1,
	"half":    2,
	"unknown": 3,
}
