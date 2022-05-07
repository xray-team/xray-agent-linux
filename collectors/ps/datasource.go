package ps

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/xray-team/xray-agent-linux/reader"
)

var psPidRe = regexp.MustCompile(`^\d+$`)

type psDataSource struct {
	filePath  string
	logPrefix string
}

// NewPSDataSource returns a new DataSource.
func NewPSDataSource(filePath, logPrefix string) *psDataSource {
	if filePath == "" {
		return nil
	}

	return &psDataSource{
		filePath:  filePath,
		logPrefix: logPrefix,
	}
}

func (ds *psDataSource) GetData() (*PS, error) {
	f, err := reader.ReadDir(ds.filePath, ds.logPrefix)
	if err != nil {
		return nil, err
	}

	var ps PS

	for _, ff := range f {
		if ff.IsDir() && psPidRe.Match([]byte(ff.Name())) {
			ps.Count++

			status, err := ds.readProcessStatus(filepath.Join(ds.filePath, ff.Name(), "/status"))
			if err != nil {
				continue
			}

			ps.Threads += status.Threads

			// State
			// Current state of the process. One of:
			//  "R (running)",
			//  "S (sleeping)",
			//  "I (idle)"
			//  "D (disk sleep)",
			//  "T (stopped)",
			//  "T (tracing stop)",
			//  "Z (zombie)",
			//  "X (dead)"
			switch status.State {
			case "R":
				ps.InStateRunning++
			case "S":
				ps.InStateSleeping++
			case "I":
				ps.InStateIdle++
			case "Z":
				ps.InStateZombie++
			case "T":
				ps.InStateStopped++
			case "X":
				ps.InStateDead++
			case "D":
				ps.InStateDiskSleep++
			}
		}
	}

	if ps.Count == 0 {
		return nil, fmt.Errorf("parsePS: no proc")
	}

	ps.Limit, _ = reader.ReadInt64File(filepath.Join(ds.filePath, PIDsLimit), ds.logPrefix)
	ps.ThreadsLimit, _ = reader.ReadInt64File(filepath.Join(ds.filePath, ThreadsLimit), ds.logPrefix)

	return &ps, nil
}

type processStatus struct {
	Name    string
	State   string
	PID     int64
	Threads int64
}

func (ds *psDataSource) readProcessStatus(filePath string) (*processStatus, error) {
	lines, err := reader.ReadMultilineFile(filePath, ds.logPrefix)
	if err != nil {
		return nil, err
	}

	var out processStatus

	for _, v := range lines {
		fields := strings.Fields(v)
		// skip incorrect lines
		if len(fields) < 2 {
			continue
		}

		switch fields[0] {
		case "Name:":
			out.Name = fields[1]
		case "State:":
			out.State = fields[1]
		case "Pid:":
			out.PID, err = strconv.ParseInt(fields[1], 10, 64)
			if err != nil {
				return nil, err
			}
		case "Threads:":
			out.Threads, err = strconv.ParseInt(fields[1], 10, 64)
			if err != nil {
				return nil, err
			}
		}
	}

	return &out, nil
}
