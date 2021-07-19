package proc

import (
	"strconv"
	"strings"

	"xray-agent-linux/dto"
	"xray-agent-linux/logger"
	"xray-agent-linux/reader"
)

type diskStatsDataSource struct {
	filePath  string
	logPrefix string
}

func NewBlockDevDataSource(filePath, logPrefix string) *diskStatsDataSource {
	if filePath == "" {
		return nil
	}

	return &diskStatsDataSource{
		filePath:  filePath,
		logPrefix: logPrefix,
	}
}

func (ds *diskStatsDataSource) GetData() ([]dto.DiskStat, error) {
	out := make([]dto.DiskStat, 0)
	// read file to memory
	lines, err := reader.ReadMultilineFile(ds.filePath, ds.logPrefix)
	if err != nil {
		return nil, err
	}

	for _, v := range lines {
		fields := strings.Fields(v)

		var disk dto.DiskStat

		// before Kernel 4.18
		if len(fields) >= 14 {
			// Major
			disk.Major, err = strconv.ParseInt(fields[0], 10, 64)
			if err != nil {
				logger.LogReadFileFieldError(ds.logPrefix, ds.filePath, "Major", err)

				continue
			}
			// Miner
			disk.Miner, err = strconv.ParseInt(fields[1], 10, 64)
			if err != nil {
				logger.LogReadFileFieldError(ds.logPrefix, ds.filePath, "Minor", err)

				continue
			}
			// Dev
			disk.Dev = fields[2]
			// ReadsCompletedSuccessfully
			disk.ReadsCompletedSuccessfully, err = strconv.ParseUint(fields[3], 10, 64)
			if err != nil {
				logger.LogReadFileFieldError(ds.logPrefix, ds.filePath, "ReadsCompletedSuccessfully", err)

				continue
			}
			// ReadsMerged
			disk.ReadsMerged, err = strconv.ParseUint(fields[4], 10, 64)
			if err != nil {
				logger.LogReadFileFieldError(ds.logPrefix, ds.filePath, "ReadsMerged", err)

				continue
			}
			// SectorsRead
			disk.SectorsRead, err = strconv.ParseUint(fields[5], 10, 64)
			if err != nil {
				logger.LogReadFileFieldError(ds.logPrefix, ds.filePath, "SectorsRead", err)

				continue
			}
			// TimeSpentReading
			disk.TimeSpentReading, err = strconv.ParseUint(fields[6], 10, 64)
			if err != nil {
				logger.LogReadFileFieldError(ds.logPrefix, ds.filePath, "TimeSpentReading", err)

				continue
			}
			// WritesCompleted
			disk.WritesCompleted, err = strconv.ParseUint(fields[7], 10, 64)
			if err != nil {
				logger.LogReadFileFieldError(ds.logPrefix, ds.filePath, "WritesCompleted", err)

				continue
			}
			// WritesMerged
			disk.WritesMerged, err = strconv.ParseUint(fields[8], 10, 64)
			if err != nil {
				logger.LogReadFileFieldError(ds.logPrefix, ds.filePath, "WritesMerged", err)

				continue
			}
			// SectorsWritten
			disk.SectorsWritten, err = strconv.ParseUint(fields[9], 10, 64)
			if err != nil {
				logger.LogReadFileFieldError(ds.logPrefix, ds.filePath, "SectorsWritten", err)

				continue
			}
			// TimeSpentWriting
			disk.TimeSpentWriting, err = strconv.ParseUint(fields[10], 10, 64)
			if err != nil {
				logger.LogReadFileFieldError(ds.logPrefix, ds.filePath, "TimeSpentWriting", err)

				continue
			}
			// IOsCurrentlyInProgress
			disk.IOsCurrentlyInProgress, err = strconv.ParseUint(fields[11], 10, 64)
			if err != nil {
				logger.LogReadFileFieldError(ds.logPrefix, ds.filePath, "IOsCurrentlyInProgress", err)

				continue
			}
			// TimeSpentDoingIOs
			disk.TimeSpentDoingIOs, err = strconv.ParseUint(fields[12], 10, 64)
			if err != nil {
				logger.LogReadFileFieldError(ds.logPrefix, ds.filePath, "TimeSpentDoingIOs", err)

				continue
			}
			// WeightedTimeSpentDoingIOs
			disk.WeightedTimeSpentDoingIOs, err = strconv.ParseUint(fields[13], 10, 64)
			if err != nil {
				logger.LogReadFileFieldError(ds.logPrefix, ds.filePath, "WeightedTimeSpentDoingIOs", err)

				continue
			}
		}

		// after Kernel 4.18
		if len(fields) >= 18 {
			// DiscardsCompletedSuccessfully
			disk.DiscardsCompletedSuccessfully, err = strconv.ParseUint(fields[14], 10, 64)
			if err != nil {
				logger.LogReadFileFieldError(ds.logPrefix, ds.filePath, "DiscardsCompletedSuccessfully", err)

				continue
			}
			// DiscardsMerged
			disk.DiscardsMerged, err = strconv.ParseUint(fields[15], 10, 64)
			if err != nil {
				logger.LogReadFileFieldError(ds.logPrefix, ds.filePath, "DiscardsMerged", err)

				continue
			}
			// SectorsDiscarded
			disk.SectorsDiscarded, err = strconv.ParseUint(fields[16], 10, 64)
			if err != nil {
				logger.LogReadFileFieldError(ds.logPrefix, ds.filePath, "SectorsDiscarded", err)

				continue
			}
			// TimeSpentDiscarding
			disk.TimeSpentDiscarding, err = strconv.ParseUint(fields[17], 10, 64)
			if err != nil {
				logger.LogReadFileFieldError(ds.logPrefix, ds.filePath, "TimeSpentDiscarding", err)

				continue
			}
		}

		// after Kernel 5.5
		if len(fields) == 20 {
			// FlushRequestsCompletedSuccessfully
			disk.FlushRequestsCompletedSuccessfully, err = strconv.ParseUint(fields[16], 10, 64)
			if err != nil {
				logger.LogReadFileFieldError(ds.logPrefix, ds.filePath, "FlushRequestsCompletedSuccessfully", err)

				continue
			}
			// TimeSpentFlushing
			disk.TimeSpentFlushing, err = strconv.ParseUint(fields[17], 10, 64)
			if err != nil {
				logger.LogReadFileFieldError(ds.logPrefix, ds.filePath, "TimeSpentFlushing", err)

				continue
			}
		}

		out = append(out, disk)
	}

	return out, err
}
