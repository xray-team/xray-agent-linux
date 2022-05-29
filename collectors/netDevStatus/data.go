package netDevStatus

type NetDevStatus struct {
	// OperState
	//   https://tools.ietf.org/html/rfc2863#section-3.1.14
	//   "up", "down", "dormant", "notPresent", "lowerLayerDown", "unknown" or "testing".
	OperState string // awk '{ split(FILENAME, array, "/"); print array[5] ": " $1 }' $(find /sys/class/net/*/operstate)
	// Duplex
	//   https://www.kernel.org/doc/Documentation/ABI/testing/sysfs-class-net
	//   "full", "half", "unknown"
	Duplex         string
	Speed          int64
	MTU            int64
	CarrierChanges int64 // awk '{ split(FILENAME, array, "/"); print array[5] ": " $1 }' $(find /sys/class/net/*/carrier_changes)
}
