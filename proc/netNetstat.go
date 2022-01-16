package proc

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/xray-team/xray-agent-linux/dto"
	"github.com/xray-team/xray-agent-linux/reader"
)

type netStatDataSource struct {
	filePath  string
	logPrefix string
}

// NewNetStatDataSource returns a new DataSource.
func NewNetStatDataSource(filePath, logPrefix string) *netStatDataSource {
	if filePath == "" {
		return nil
	}

	return &netStatDataSource{
		filePath:  filePath,
		logPrefix: logPrefix,
	}
}

// GetData parse files /proc/net/netstat, /proc/net/snmp (/proc/$PID/net/netstat, /proc/$PID/net/snmp).
func (ds *netStatDataSource) GetData() (*dto.Netstat, error) {
	// Initialize map for results
	var out dto.Netstat
	out.Ext = make(map[string]map[string]int64)

	// read file to memory
	lines, err := reader.ReadMultilineFile(ds.filePath, ds.logPrefix)
	if err != nil {
		return nil, err
	}

	// Initialize temp map
	tempMap := make(map[string][]string)
	// loop for file lines
	for _, line := range lines {
		fields := strings.Fields(line)
		// Trim ':' from ext name
		ext := strings.TrimSuffix(fields[0], ":")
		// Cut ext name
		fields = fields[1:]

		params, ok := tempMap[ext]

		if ok {
			if len(fields) != len(params) {
				return nil, fmt.Errorf("can't parse file %s", ds.filePath)
			}

			// Initialize nested map for results
			out.Ext[ext] = make(map[string]int64)

			var ii int64

			for i, param := range params {
				ii, err = strconv.ParseInt(fields[i], 10, 64)
				if err != nil {
					return nil, err
				}

				out.Ext[ext][param] = ii
			}
		} else {
			tempMap[ext] = fields
		}
	}

	return &out, err
}
