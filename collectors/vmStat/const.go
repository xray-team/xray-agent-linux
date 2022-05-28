package vmStat

const (
	CollectorName = "VMStat"
	VMStatPath    = "/proc/vmstat"
)

// Metrics
const (
	ResourceName = "VMStat"

	MetricPgPgIn         = "PgPgIn"
	MetricPgPgOut        = "PgPgOut"
	MetricPSwpIn         = "PSwpIn"
	MetricPSwpOut        = "PSwpOut"
	MetricPgFault        = "PgFault"
	MetricPgMajFault     = "PgMajFault"
	MetricPgFree         = "PgFree"
	MetricPgActivate     = "PgActivate"
	MetricPgDeactivate   = "PgDeactivate"
	MetricPgLazyFree     = "PgLazyFree"
	MetricPgLazyFreed    = "PgLazyFreed"
	MetricPgRefill       = "PgRefill"
	MetricNumaHit        = "NumaHit"
	MetricNumaMiss       = "NumaMiss"
	MetricNumaForeign    = "NumaForeign"
	MetricNumaInterleave = "NumaInterleave"
	MetricNumaLocal      = "NumaLocal"
	MetricNumaOther      = "NumaOther"
	MetricOOMKill        = "OOMKill"
)
