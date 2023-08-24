package interrupts

import (
	"errors"
	"strconv"
	"strings"

	"github.com/xray-team/xray-agent-linux/reader"
)

type interruptsDataSource struct {
	filePath  string
	logPrefix string
}

// NewDataSource returns a new DataSource.
func NewDataSource(filePath, logPrefix string) *interruptsDataSource {
	if filePath == "" {
		return nil
	}

	return &interruptsDataSource{
		filePath:  filePath,
		logPrefix: logPrefix,
	}
}

func (ds *interruptsDataSource) GetData() (*Interrupts, error) {
	var (
		data Interrupts
		err  error
	)

	lines, err := reader.ReadMultilineFile(ds.filePath, ds.logPrefix)
	if err != nil {
		return nil, err
	}

	if len(lines) < 2 {
		return nil, errors.New("invalid interrupts datasource format")
	}

	coresCount := len(strings.Fields(lines[0]))
	if coresCount <= 0 {
		return nil, errors.New("cant calculate cores total number")
	}

	data.PerCPU = make(map[int]int64)
	for _, line := range lines[1:] {
		fields := strings.Fields(line)

		var irqTotalsPerCPU []string

		// Not all lines in /proc/interrupts have the same number of fields
		// This condition makes parsing correct.
		if len(fields) <= 1+coresCount {
			irqTotalsPerCPU = fields[1:]
		} else {
			irqTotalsPerCPU = fields[1 : 1+coresCount]
		}

		// Summation of interrupts for each cpu core
		for i, v := range irqTotalsPerCPU {
			irqPerCoreTotal, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				return nil, err
			}

			data.PerCPU[i] += irqPerCoreTotal
		}
	}

	// Total interrupts calculation
	for _, v := range data.PerCPU {
		data.Total += v
	}

	return &data, nil
}
