package dto

// https://www.kernel.org/doc/Documentation/ABI/testing/sysfs-class-net
type ClassNet struct {
	Uevent *NetDevUevent
	// ProtocolType indicates the interface protocol type as a decimal value.
	// Examples:
	// 1 - Ethernet
	// 32 - InfiniBand
	// 772 - Loopback device
	ProtocolType int64
	// Speed indicates the interface latest or current speed value.
	// Value is	an integer representing the link speed in Mbits/sec.
	// Examples:
	//  100 - Fast Ethernet
	//  1000 - Gigabit Ethernet
	//  10000 - 10 Gigabit Ethernet
	//  -1  - port is Down
	//  0 - not applied for interface
	Speed int64
	MTU   int64
	// CarrierChanges
	// awk '{ split(FILENAME, array, "/"); print array[5] ": " $1 }' $(find /sys/class/net/*/carrier_changes)
	CarrierChanges int64
	// Duplex
	// https://www.kernel.org/doc/Documentation/ABI/testing/sysfs-class-net
	// Possible values are: "full", "half", "unknown"
	Duplex string
	// OperState indicates the interface RFC2863 operational state as a string.
	// https://tools.ietf.org/html/rfc2863#section-3.1.14
	// Possible values are: "unknown", "notpresent", "down", "lowerlayerdown", "testing", "dormant", "up".
	// awk '{ split(FILENAME, array, "/"); print array[5] ": " $1 }' $(find /sys/class/net/*/operstate)
	OperState string
	// MACAddress - hardware address currently assigned to this interface.
	// Format is a string, e.g: 00:11:22:33:44:55 for an Ethernet MAC address.
	MACAddress string
	Lower      []string
	Upper      []string
	Device     bool
	Virtual    bool
	BondMaster bool
	BondSlave  bool
}

func (netDev *ClassNet) IsVirtual() bool {
	return netDev.Virtual
}

func (netDev *ClassNet) IsDevice() bool {
	return netDev.Device
}

func (netDev *ClassNet) IsBond() bool {
	return netDev.HasDevType(NetDevTypeBond)
}

func (netDev *ClassNet) IsBridge() bool {
	return netDev.HasDevType(NetDevTypeBridge)
}

func (netDev *ClassNet) IsVlan() bool {
	return netDev.HasDevType(NetDevTypeVlan)
}

func (netDev *ClassNet) IsWireless() bool {
	return netDev.HasDevType(NetDevTypeWireless)
}

func (netDev *ClassNet) HasDevType(devType string) bool {
	return netDev.Uevent.DevType == devType
}

func (netDev *ClassNet) HasParent(parent string) bool {
	for _, lower := range netDev.Lower {
		if lower == parent {
			return true
		}
	}

	return false
}

func (netDev *ClassNet) HasChild(child string) bool {
	for _, upper := range netDev.Upper {
		if upper == child {
			return true
		}
	}

	return false
}

func (netDev *ClassNet) IsLoopback() bool {
	return netDev.ProtocolType == NetDevProtocolTypeLoopback
}

type NetDevUevent struct {
	Interface string
	// DevType never set for ethernet devices
	// Possible values are: "", "vlan", "wlan", "bridge", "bond" ...
	DevType string
	IfIndex int64
}

type ClassBlock struct {
	Uevent *BlockDevUevent
}

// BlockDevUevent - describes /sys/dev/block/$DEV/uevent or /sys/class/block/$DEV/uevent fields
type BlockDevUevent struct {
	Major      int64
	Minor      int64
	DevName    string
	DevType    string // disk | partition
	PartNumber string
	PartName   string
}
