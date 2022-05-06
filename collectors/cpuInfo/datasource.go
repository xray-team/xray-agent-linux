package cpuInfo

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/xray-team/xray-agent-linux/dto"
	"github.com/xray-team/xray-agent-linux/reader"
)

type cpuInfoDataSource struct {
	filePath  string
	logPrefix string
}

// NewCPUInfoDataSource returns a new DataSource.
func NewCPUInfoDataSource(filePath, logPrefix string) *cpuInfoDataSource {
	if filePath == "" {
		return nil
	}

	return &cpuInfoDataSource{
		filePath:  filePath,
		logPrefix: logPrefix,
	}
}

func (ds *cpuInfoDataSource) GetData() (*dto.CPUInfo, error) {
	var out dto.CPUInfo
	out.CPU = make([]dto.CPU, 0)

	// read file to memory
	data, err := reader.ReadFile(ds.filePath, ds.logPrefix)
	if err != nil {
		return nil, err
	}

	processorStrings := strings.Split(strings.ReplaceAll(string(data), "processor", "|<>|processor"), "|<>|")

	for _, processor := range processorStrings {
		lines := strings.Split(processor, "\n")
		// Skip first empty element
		if len(lines) == 1 {
			continue
		}

		var cpuInfo = dto.CPU{}

		for _, line := range lines {
			columns := strings.Split(line, ":")
			// skip incorrect lines
			if len(columns) < 2 {
				continue
			}

			switch strings.TrimSpace(columns[0]) {
			case "processor": // Processor number
				cpuInfo.ProcessorNumber, err = strconv.Atoi(strings.TrimSpace(columns[1]))
				if err != nil {
					return nil, fmt.Errorf("error while trying to read file %s, processor line: %s: %s", ds.filePath, line, err.Error())
				}
			case "vendor_id":
				cpuInfo.VendorID = strings.TrimSpace(columns[1])
			case "model name":
				cpuInfo.ModelName = strings.TrimSpace(columns[1])
			case "cpu MHz":
				cpuInfo.MHz, err = strconv.ParseFloat(strings.TrimSpace(columns[1]), 64)
				if err != nil {
					return nil, fmt.Errorf("error while trying to read file %s, cpu MHz line: %s: %s", ds.filePath, line, err.Error())
				}
			case "cache size":
				cpuInfo.CacheSize = strings.TrimSpace(columns[1])
			case "physical id":
				cpuInfo.PhysicalID, err = strconv.Atoi(strings.TrimSpace(columns[1]))
				if err != nil {
					return nil, fmt.Errorf("error while trying to read file %s, physical id line: %s: %s", ds.filePath, line, err.Error())
				}
			case "core id":
				cpuInfo.CoreID, err = strconv.Atoi(strings.TrimSpace(columns[1]))
				if err != nil {
					return nil, fmt.Errorf("error while trying to read file %s, core id line: %s: %s", ds.filePath, line, err.Error())
				}
			case "cpu cores":
				cpuInfo.CPUCores, err = strconv.Atoi(strings.TrimSpace(columns[1]))
				if err != nil {
					return nil, fmt.Errorf("error while trying to read file %s, cpu cores line: %s: %s", ds.filePath, line, err.Error())
				}
			case "bugs":
				cpuInfo.Bugs = strings.TrimSpace(columns[1])
			}
		}

		out.CPU = append(out.CPU, cpuInfo)
	}

	return &out, nil
}
