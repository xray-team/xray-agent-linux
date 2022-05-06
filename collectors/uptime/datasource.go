package uptime

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/xray-team/xray-agent-linux/dto"
	"github.com/xray-team/xray-agent-linux/reader"
)

/*
/proc/uptime
	The first value represents the total number of seconds the system has been up.
	The second value is the sum of how much time each core has spent idle, in seconds.
	Consequently, the second value may be greater than the overall system uptime on systems with multiple cores.
*/

type uptimeDataSource struct {
	filePath  string
	logPrefix string
}

// NewUptimeDataSource returns a new DataSource.
func NewUptimeDataSource(filePath, logPrefix string) *uptimeDataSource {
	if filePath == "" {
		return nil
	}

	return &uptimeDataSource{
		filePath:  filePath,
		logPrefix: logPrefix,
	}
}

func (ds *uptimeDataSource) GetData() (*dto.Uptime, error) {
	var uptime dto.Uptime

	// read file to memory
	data, err := reader.ReadFile(ds.filePath, ds.logPrefix)
	if err != nil {
		return nil, fmt.Errorf("cannot read file uptime file %s. %s", ds.filePath, err)
	}

	fields := strings.Fields(string(data))
	if len(fields) != 2 {
		return nil, fmt.Errorf("not valid number of fields in uptime file, needs 2: %s", ds.filePath)
	}

	uptime.Uptime, err = strconv.ParseFloat(fields[0], 64)
	if err != nil {
		return nil, fmt.Errorf("not format of Uptime field in uptime file, needs float: %s", ds.filePath)
	}

	uptime.Idle, err = strconv.ParseFloat(fields[1], 64)
	if err != nil {
		return nil, fmt.Errorf("not format of Idle field in uptime file, needs float: %s", ds.filePath)
	}

	return &uptime, nil
}
