package dto

// NetDevs describes the content of the file /proc/netdev
// https://github.com/torvalds/linux/blob/master/include/linux/netdevice.h
// https://github.com/torvalds/linux/blob/master/net/core/net-procfs.c
// Physical are all the network interfaces that are listed in /proc/net/dev, but do not exist in /sys/devices/virtual/net
type NetDevStatistics struct {
	RxBytes       uint64
	RxPackets     uint64
	RxErrs        uint64
	RxDrop        uint64
	RxFifoErrs    uint64
	RxFrameErrs   uint64 // rx_length_errors + rx_over_errors + rx_crc_errors + rx_frame_errors
	RxCompressed  uint64
	Multicast     uint64
	TxBytes       uint64
	TxPackets     uint64
	TxErrs        uint64
	TxDrop        uint64
	TxFifoErrs    uint64
	Collisions    uint64
	TxCarrierErrs uint64 // tx_carrier_errors + tx_aborted_errors + tx_window_errors + tx_heartbeat_errors
	TxCompressed  uint64
}
