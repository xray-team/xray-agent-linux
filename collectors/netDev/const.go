package netDev

const (
	CollectorName = "NetDev"
	NetDevPath    = "/net/dev"
)

// Metrics
const (
	ResourceName     = "NetDev"
	SetNameInterface = "Interface"

	MetricStatisticsRxBytes       = "RxBytes"
	MetricStatisticsRxPackets     = "RxPackets"
	MetricStatisticsRxErrs        = "RxErrs"
	MetricStatisticsRxDrop        = "RxDrop"
	MetricStatisticsRxFifoErrs    = "RxFifoErrs"
	MetricStatisticsRxFrameErrs   = "RxFrameErrs"
	MetricStatisticsRxCompressed  = "RxCompressed"
	MetricStatisticsMulticast     = "Multicast"
	MetricStatisticsTxBytes       = "TxBytes"
	MetricStatisticsTxPackets     = "TxPackets"
	MetricStatisticsTxErrs        = "TxErrs"
	MetricStatisticsTxDrop        = "TxDrop"
	MetricStatisticsTxFifoErrs    = "TxFifoErrs"
	MetricStatisticsCollisions    = "Collisions"
	MetricStatisticsTxCarrierErrs = "TxCarrierErrs"
	MetricStatisticsTxCompressed  = "TxCompressed"
)
