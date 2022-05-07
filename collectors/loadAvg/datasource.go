package loadAvg

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/xray-team/xray-agent-linux/logger"
	"github.com/xray-team/xray-agent-linux/reader"
)

/*
http://man7.org/linux/man-pages/man5/proc.5.html
	The first three fields in this file are load average figures giving the number of jobs in the run queue (state R)
	or waiting for disk I/O (state D) averaged over 1, 5, and 15 minutes.
	They are the same as the load average numbers given by uptime(1) and other programs.
	The fourth field consists of two numbers separated by a slash (/).  The first of these is
	the number of currently runnable kernel scheduling entities (processes, threads).
	The value after the slash is the number of kernel scheduling entities that currently exist on the system.
	The fifth field is the PID of the process that was most recently created on the system.

	https://github.com/torvalds/linux/blob/master/include/linux/sched/loadavg.h#L21-L23
	https://github.com/torvalds/linux/blob/master/kernel/sched/loadavg.c
*/

type loadAvgDataSource struct {
	filePath  string
	logPrefix string
}

// NewLoadAvgDataSource returns a new DataSource.
func NewLoadAvgDataSource(filePath, logPrefix string) *loadAvgDataSource {
	if filePath == "" {
		return nil
	}

	return &loadAvgDataSource{
		filePath:  filePath,
		logPrefix: logPrefix,
	}
}

func (ds *loadAvgDataSource) GetData() (*LoadAvg, error) {
	var la = LoadAvg{}

	// read file to memory
	data, err := reader.ReadFile(ds.filePath, ds.logPrefix)
	if err != nil {
		return nil, err
	}

	fields := strings.Fields(string(data))
	if len(fields) != 5 {
		return nil, fmt.Errorf("not valid loadavg file: %s", ds.filePath)
	}

	// Last
	la.Last, err = strconv.ParseFloat(fields[0], 64)
	if err != nil {
		logger.Log.Debug.Printf(logger.MessageReadFileFieldError, ds.logPrefix, ds.filePath, "Last", err.Error())

		return nil, err
	}

	// Last5m
	la.Last5m, err = strconv.ParseFloat(fields[1], 64)
	if err != nil {
		logger.Log.Debug.Printf(logger.MessageReadFileFieldError, ds.logPrefix, ds.filePath, "Last5m", err.Error())

		return nil, err
	}

	// Last15m
	la.Last15m, err = strconv.ParseFloat(fields[2], 64)
	if err != nil {
		logger.Log.Debug.Printf(logger.MessageReadFileFieldError, ds.logPrefix, ds.filePath, "Last15m", err.Error())

		return nil, err
	}

	// KernelSchedulingEntities
	kse := strings.Split(fields[3], "/")

	if len(kse) != 2 {
		return nil, fmt.Errorf("not valid loadavg file: %s", ds.filePath)
	}

	la.KernelSchedulingEntities, err = strconv.ParseInt(kse[1], 10, 64)
	if err != nil {
		logger.Log.Debug.Printf(logger.MessageReadFileFieldError, ds.logPrefix, ds.filePath, "KernelSchedulingEntities", err.Error())

		return nil, err
	}

	return &la, nil
}
