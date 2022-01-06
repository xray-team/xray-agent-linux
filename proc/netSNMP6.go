package proc

import (
	"strconv"
	"strings"

	"github.com/xray-team/xray-agent-linux/dto"
	"github.com/xray-team/xray-agent-linux/reader"
)

type netSNMP6DataSource struct {
	filePath  string
	logPrefix string
}

func NewNetSNMP6DataSource(filePath, logPrefix string) *netSNMP6DataSource {
	if filePath == "" {
		return nil
	}

	return &netSNMP6DataSource{
		filePath:  filePath,
		logPrefix: logPrefix,
	}
}

// GetData parse file /proc/net/snmp6 (/proc/$PID/net/snmp6)
func (ds *netSNMP6DataSource) GetData() (*dto.NetSNMP6, error) {
	// read file to memory
	lines, err := reader.ReadMultilineFile(ds.filePath, ds.logPrefix)
	if err != nil {
		return nil, err
	}

	// Initialize map for results
	var out dto.NetSNMP6
	out.Counters = make(map[string]int64, len(lines))

	// loop for file lines
	for _, line := range lines {
		fields := strings.Fields(line)
		// skip incorrect lines
		if len(fields) != 2 {
			continue
		}

		i, err := strconv.ParseInt(fields[1], 10, 64)
		if err != nil {
			continue
		}

		out.Counters[fields[0]] = i
	}

	return &out, nil
}
