package dto

const (
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
