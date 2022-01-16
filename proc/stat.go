package proc

import (
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"

	"github.com/xray-team/xray-agent-linux/dto"
	"github.com/xray-team/xray-agent-linux/reader"
)

type statDataSource struct {
	filePath  string
	logPrefix string
}

// NewStatDataSource returns a new DataSource.
func NewStatDataSource(filePath, logPrefix string) *statDataSource {
	if filePath == "" {
		return nil
	}

	return &statDataSource{
		filePath:  filePath,
		logPrefix: logPrefix,
	}
}

func (ds *statDataSource) GetData() (*dto.Stat, error) {
	stats := dto.Stat{
		PerCPU: make(map[string]dto.CPUStats),
	}

	lines, err := reader.ReadMultilineFile(ds.filePath, ds.logPrefix)
	if err != nil {
		return nil, err
	}

	// RE for matching per CPU value
	cpuRE := regexp.MustCompile(`^cpu\d+$`)

	for _, line := range lines {
		fields := strings.Fields(line)

		// skip incorrect lines
		if len(fields) < 2 {
			continue
		}

		// PerCPU
		if cpuRE.Match([]byte(fields[0])) {
			var (
				cpuName  string
				cpuStats dto.CPUStats
			)

			cpuName, cpuStats, err = parseProcStatCPULine(line)
			if err != nil {
				return nil, err
			}

			stats.PerCPU[strings.TrimPrefix(cpuName, "cpu")] = cpuStats

			continue
		}

		switch fields[0] {
		case "cpu": // CPUTotal
			_, stats.CPU, err = parseProcStatCPULine(line)
			if err != nil {
				return nil, err
			}
		case "intr":
			stats.Intr, err = strconv.ParseUint(fields[1], 10, 64)
			if err != nil {
				return nil, err
			}
		case "ctxt":
			stats.Ctxt, err = strconv.ParseUint(fields[1], 10, 64)
			if err != nil {
				return nil, err
			}
		case "btime":
			stats.Btime, err = strconv.ParseUint(fields[1], 10, 64)
			if err != nil {
				return nil, err
			}
		case "processes":
			stats.Processes, err = strconv.ParseUint(fields[1], 10, 64)
			if err != nil {
				return nil, err
			}
		case "procs_running":
			stats.ProcessesRunning, err = strconv.ParseUint(fields[1], 10, 64)
			if err != nil {
				return nil, err
			}
		case "procs_blocked":
			stats.ProcessesBlocked, err = strconv.ParseUint(fields[1], 10, 64)
			if err != nil {
				return nil, err
			}
		case "softirq":
			stats.SoftIRQ, err = parseProcStatSoftIRQLine(line)
			if err != nil {
				return nil, err
			}
		}
	}

	return &stats, nil
}

func parseProcStatCPULine(line string) (string, dto.CPUStats, error) {
	var (
		out     dto.CPUStats
		cpuName string
	)

	count, err := fmt.Sscanf(line, "%s %d %d %d %d %d %d %d %d %d %d %d",
		&cpuName,
		&out.User,
		&out.Nice,
		&out.System,
		&out.Idle,
		&out.IOwait,
		&out.IRQ,
		&out.SoftIRQ,
		&out.Steal,
		&out.Guest,
		&out.GuestNice,
		&out.GuestNice,
	)

	if err != nil && err != io.EOF {
		return cpuName, dto.CPUStats{}, fmt.Errorf("can't parse cpu line: %s: %s", line, err.Error())
	}

	if count < 8 {
		return cpuName, dto.CPUStats{}, fmt.Errorf("can't parse cpu line: %s", line)
	}

	return cpuName, out, nil
}

func parseProcStatSoftIRQLine(line string) (dto.SoftIRQStat, error) {
	out := dto.SoftIRQStat{}

	_, err := fmt.Sscanf(line, "softirq %d %d %d %d %d %d %d %d %d %d %d",
		&out.Total,
		&out.Hi,
		&out.Timer,
		&out.NetTx,
		&out.NetRx,
		&out.Block,
		&out.IRQPoll,
		&out.Tasklet,
		&out.Sched,
		&out.HRTimer,
		&out.RCU,
	)

	if err != nil {
		return dto.SoftIRQStat{}, fmt.Errorf("can't parse softirq line: %s: %s", line, err.Error())
	}

	return out, nil
}
