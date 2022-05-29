package diskStat

const (
	CollectorName = "DiskStat"
	DiskStatsPath = "/proc/diskstats"
)

// Metrics
const (
	ResourceName = "DiskStat"
	SetNameDev   = "Dev"

	MetricReadsCompletedSuccessfully         = "ReadsCompletedSuccessfully"
	MetricReadsMerged                        = "ReadsMerged"
	MetricSectorsRead                        = "SectorsRead"
	MetricTimeSpentReading                   = "TimeSpentReading"
	MetricWritesCompleted                    = "WritesCompleted"
	MetricWritesMerged                       = "WritesMerged"
	MetricSectorsWritten                     = "SectorsWritten"
	MetricTimeSpentWriting                   = "TimeSpentWriting"
	MetricIOsCurrentlyInProgress             = "IOsCurrentlyInProgress"
	MetricTimeSpentDoingIOs                  = "TimeSpentDoingIOs"
	MetricWeightedTimeSpentDoingIOs          = "WeightedTimeSpentDoingIOs"
	MetricDiscardsCompletedSuccessfully      = "DiscardsCompletedSuccessfully"
	MetricDiscardsMerged                     = "DiscardsMerged"
	MetricSectorsDiscarded                   = "SectorsDiscarded"
	MetricTimeSpentDiscarding                = "TimeSpentDiscarding"
	MetricFlushRequestsCompletedSuccessfully = "FlushRequestsCompletedSuccessfully"
	MetricTimeSpentFlushing                  = "TimeSpentFlushing"
)
