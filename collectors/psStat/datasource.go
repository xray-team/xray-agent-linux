package psStat

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/xray-team/xray-agent-linux/logger"
	"github.com/xray-team/xray-agent-linux/reader"
)

var psPidRe = regexp.MustCompile(`^\d+$`)

type psStatDataSource struct {
	filePath  string
	logPrefix string
}

// NewPSStatDataSource returns a new DataSource.
func NewPSStatDataSource(filePath, logPrefix string) *psStatDataSource {
	if filePath == "" {
		return nil
	}

	return &psStatDataSource{
		filePath:  filePath,
		logPrefix: logPrefix,
	}
}

func (ds *psStatDataSource) GetData() (*PSStat, error) {
	f, err := reader.ReadDir(ds.filePath, ds.logPrefix)
	if err != nil {
		return nil, err
	}

	out := PSStat{PS: make([]ProcessStat, 0)}

	for _, ff := range f {
		if ff.IsDir() && psPidRe.Match([]byte(ff.Name())) {
			stat, err := ds.readProcessStat(filepath.Join(ds.filePath, ff.Name(), "/stat"))
			if err != nil {
				continue
			}

			out.PS = append(out.PS, *stat)
		}
	}

	return &out, nil
}

func (ds *psStatDataSource) readProcessStat(filePath string) (*ProcessStat, error) {
	// read file to memory
	data, err := reader.ReadFile(filePath, ds.logPrefix)
	if err != nil {
		return nil, err
	}

	fields := strings.Fields(string(data))
	if len(fields) < 44 {
		return nil, fmt.Errorf("not valid stat file: %s", ds.filePath)
	}

	var out ProcessStat

	// PID
	out.PID, err = strconv.ParseInt(fields[0], 10, 64)
	if err != nil {
		logger.Log.Debug.Printf(logger.MessageReadFileFieldError, ds.logPrefix, ds.filePath, "PID", err.Error())

		return nil, err
	}

	// Name
	out.Name = fields[1]
	// State
	out.State = fields[2]

	// UTime
	out.UTime, err = strconv.ParseInt(fields[13], 10, 64)
	if err != nil {
		logger.Log.Debug.Printf(logger.MessageReadFileFieldError, ds.logPrefix, ds.filePath, "UTime", err.Error())

		return nil, err
	}

	// STime
	out.STime, err = strconv.ParseInt(fields[14], 10, 64)
	if err != nil {
		logger.Log.Debug.Printf(logger.MessageReadFileFieldError, ds.logPrefix, ds.filePath, "STime", err.Error())

		return nil, err
	}

	// CuTime
	out.CuTime, err = strconv.ParseInt(fields[15], 10, 64)
	if err != nil {
		logger.Log.Debug.Printf(logger.MessageReadFileFieldError, ds.logPrefix, ds.filePath, "CuTime", err.Error())

		return nil, err
	}

	// CsTime
	out.CsTime, err = strconv.ParseInt(fields[16], 10, 64)
	if err != nil {
		logger.Log.Debug.Printf(logger.MessageReadFileFieldError, ds.logPrefix, ds.filePath, "CsTime", err.Error())

		return nil, err
	}

	// GuestTime
	out.GuestTime, err = strconv.ParseInt(fields[42], 10, 64)
	if err != nil {
		logger.Log.Debug.Printf(logger.MessageReadFileFieldError, ds.logPrefix, ds.filePath, "GuestTime", err.Error())

		return nil, err
	}

	// CGuestTime
	out.CGuestTime, err = strconv.ParseInt(fields[43], 10, 64)
	if err != nil {
		logger.Log.Debug.Printf(logger.MessageReadFileFieldError, ds.logPrefix, ds.filePath, "CGuestTime", err.Error())

		return nil, err
	}

	// Threads
	out.Threads, err = strconv.ParseInt(fields[19], 10, 64)
	if err != nil {
		logger.Log.Debug.Printf(logger.MessageReadFileFieldError, ds.logPrefix, ds.filePath, "Threads", err.Error())

		return nil, err
	}

	// VSize
	out.VSize, err = strconv.ParseInt(fields[22], 10, 64)
	if err != nil {
		logger.Log.Debug.Printf(logger.MessageReadFileFieldError, ds.logPrefix, ds.filePath, "VSize", err.Error())

		return nil, err
	}

	// Rss
	out.Rss, err = strconv.ParseInt(fields[23], 10, 64)
	if err != nil {
		logger.Log.Debug.Printf(logger.MessageReadFileFieldError, ds.logPrefix, ds.filePath, "Rss", err.Error())

		return nil, err
	}

	return &out, nil
}
