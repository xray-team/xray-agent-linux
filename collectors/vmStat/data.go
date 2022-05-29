package vmStat

// VMStat describes the content of the file /proc/vmstat
// https://man7.org/linux/man-pages/man5/proc.5.html
type VMStat struct {
	PgPgIn         uint64
	PgPgOut        uint64
	PSwpIn         uint64
	PSwpOut        uint64
	PgFault        uint64
	PgMajFault     uint64
	PgFree         uint64
	PgActivate     uint64
	PgDeactivate   uint64
	PgLazyFree     uint64
	PgLazyFreed    uint64
	PgRefill       uint64
	NumaHit        uint64
	NumaMiss       uint64
	NumaForeign    uint64
	NumaInterleave uint64
	NumaLocal      uint64
	NumaOther      uint64
	OOMKill        uint64
}
