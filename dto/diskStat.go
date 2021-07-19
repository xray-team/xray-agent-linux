package dto

// DiskStat represents the content of the file /proc/diskstats
// https://www.kernel.org/doc/Documentation/ABI/testing/procfs-diskstats
// https://www.kernel.org/doc/Documentation/admin-guide/iostats.rst
// The "sectors" in question are the standard UNIX 512-byte sectors, not any device- or filesystem-specific block size.
type DiskStat struct {
	Major int64 // https://www.kernel.org/doc/Documentation/admin-guide/devices.txt
	Miner int64
	Dev   string
	// ReadsCompletedSuccessfully  is the total number of reads completed successfully.
	// (units:requests)
	ReadsCompletedSuccessfully uint64
	// ReadsMerged - reads  which are adjacent to each other may be merged for efficiency.
	// Thus two 4K reads may become one 8K read before it is ultimately handed to the disk, and so it will be counted (and queued)
	// as only one I/O.  This field lets you know how often this was done.
	// (units:requests)
	ReadsMerged uint64
	// SectorsRead  is the total number of sectors read successfully.
	// (units:sectors)
	SectorsRead uint64
	// TimeSpentReading  is the total number of milliseconds spent by all reads (as measured from __make_request() to end_that_request_last()).
	// (units:milliseconds)
	TimeSpentReading uint64 // total wait time for read requests
	//  WritesCompleted  is the total number of writes completed successfully (number of write I/Os processed).
	// (units:requests)
	WritesCompleted uint64
	// WritesMerged - writes which are adjacent to each other may be merged for efficiency.
	// Thus two 4K reads may become one 8K read before it is ultimately handed to the disk, and so it will be counted (and queued)
	// as only one I/O.  This field lets you know how often this was done.
	// (units:requests)
	WritesMerged uint64
	// This is the total number of sectors written successfully.
	// (units:sectors)
	SectorsWritten uint64
	// TimeSpentWriting  is the total number of milliseconds spent by all writes (as measured from __make_request() to end_that_request_last()).
	// (units:milliseconds)
	TimeSpentWriting uint64
	// IOsCurrentlyInProgress - number of I/Os currently in flight.
	// The only field that should go to zero. Incremented as requests are
	// given to appropriate struct request_queue and decremented as they finish.
	// (units:requests)
	IOsCurrentlyInProgress uint64
	// This field increases so long as field IOsCurrentlyInProgress is nonzero.
	// Since 5.0 this field counts jiffies when at least one request was started or completed.
	// If request runs more than 2 jiffies then some I/O time will not be accounted unless there are other requests.
	// (units:milliseconds)
	TimeSpentDoingIOs uint64
	// WeightedTimeSpentDoingIOs - total wait time for all requests. This field is incremented at each I/O start, I/O completion, I/O merge, or read of these stats by the number of I/Os in progress
	// (IOsCurrentlyInProgress) times the number of milliseconds spent doing I/O since the last update of this field.  This can provide an easy measure of both
	// I/O completion time and the backlog that may be accumulating.
	// (units:milliseconds)
	WeightedTimeSpentDoingIOs uint64
	// DiscardsCompletedSuccessfully is the total number of discards completed successfully (number of discard I/Os processed ).
	// (units:requests)		Kernel 4.18+
	DiscardsCompletedSuccessfully uint64
	// DiscardsMerged number of discard I/Os merged with in-queue I/O. See the description of field ReadsMerged
	// (units:requests)		Kernel 4.18+
	DiscardsMerged uint64
	// SectorsDiscarded is the total number of sectors discarded successfully
	// (units:sectors)		Kernel 4.18+
	SectorsDiscarded uint64
	// TimeSpentDiscarding is the total number of milliseconds spent by all discards (as measured from __make_request() to end_that_request_last()).
	// (units:milliseconds)	Kernel 4.18+
	TimeSpentDiscarding uint64
	// This is the total number of flush requests completed successfully.
	// Block layer combines flush requests and executes at most one at a time.
	// This counts flush requests executed by disk. Not tracked for partitions.
	// (units:requests)		Kernel 5.5+
	FlushRequestsCompletedSuccessfully uint64
	// This is the total number of milliseconds spent by all flush requests.
	// (units:milliseconds)	Kernel 5.5+
	TimeSpentFlushing uint64
}
