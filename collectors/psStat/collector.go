package psStat

import (
	"strconv"
	"strings"

	"github.com/xray-team/xray-agent-linux/dto"
	"github.com/xray-team/xray-agent-linux/logger"
)

type DataSource interface {
	GetData() (*PSStat, error)
}

type Collector struct {
	Config     *Config
	DataSource DataSource
}

// CreateCollector returns a new collector object.
func CreateCollector(rawConfig []byte) dto.Collector {
	config := NewConfig()

	err := config.Parse(rawConfig)
	if err != nil {
		logger.Log.Error.Printf(logger.MessageError, CollectorName, err.Error())

		return nil
	}

	return NewCollector(
		config,
		NewDataSource(ProcPath, CollectorName),
	)
}

// NewCollector returns a new collector object.
func NewCollector(config *Config, dataSource DataSource) dto.Collector {
	if config == nil || dataSource == nil {
		logger.Log.Error.Printf(logger.MessageInitCollectorError, CollectorName)

		return nil
	}

	// exit if collector disabled
	if !config.Enabled || len(config.ProcessList) == 0 {
		return nil
	}

	return &Collector{
		Config:     config,
		DataSource: dataSource,
	}
}

// GetName returns the collector's name.
func (c *Collector) GetName() string {
	return CollectorName
}

// Collect collects and returns metrics.
func (c *Collector) Collect() ([]dto.Metric, error) {
	psStats, err := c.DataSource.GetData()
	if err != nil {
		return nil, err
	}

	// Slice for results
	metrics := make([]dto.Metric, 0)
	// Temp map fot total values
	appStat := make(map[string]struct {
		processCount int64
		threads      int64
		uTime        int64
		sTime        int64
		guestTime    int64
		totalTime    int64
		vSize        int64
		rss          int64
	})

	for _, psStat := range psStats.PS {
		for _, psName := range c.Config.ProcessList {
			if psName == strings.TrimRight(strings.TrimLeft(psStat.Name, "("), ")") {
				// load
				st := appStat[psName]
				// update
				st.processCount++
				st.threads += psStat.Threads
				st.uTime += psStat.UTime + psStat.CuTime
				st.sTime += psStat.STime + psStat.CsTime
				st.guestTime += psStat.GuestTime + psStat.CGuestTime
				st.totalTime += psStat.UTime + psStat.CuTime + psStat.STime + psStat.CsTime + psStat.GuestTime + psStat.CGuestTime
				st.vSize += psStat.VSize
				st.rss += psStat.Rss
				// store
				appStat[psName] = st

				// Per PID Statistics
				if c.Config.CollectPerPidStat {
					attrs := []dto.MetricAttribute{
						{
							Name:  dto.ResourceAttr,
							Value: ResourceName,
						},
						{
							Name:  SetNameProcessName,
							Value: psName,
						},
						{
							Name:  SetNamePID,
							Value: strconv.FormatInt(psStat.PID, 10),
						},
					}

					metrics = append(metrics,
						dto.Metric{
							Name:       MetricUser,
							Value:      psStat.UTime + psStat.CuTime,
							Attributes: attrs,
						},
						dto.Metric{
							Name:       MetricSystem,
							Value:      psStat.STime + psStat.CsTime,
							Attributes: attrs,
						},
						dto.Metric{
							Name:       MetricGuest,
							Value:      psStat.GuestTime + psStat.CGuestTime,
							Attributes: attrs,
						},
						dto.Metric{
							Name:       MetricTotal,
							Value:      psStat.UTime + psStat.CuTime + psStat.STime + psStat.CsTime + psStat.GuestTime + psStat.CGuestTime,
							Attributes: attrs,
						},
						dto.Metric{
							Name:       MetricThreads,
							Value:      psStat.Threads,
							Attributes: attrs,
						},
						dto.Metric{
							Name:       MetricVirtualMemorySize,
							Value:      psStat.VSize,
							Attributes: attrs,
						},
						dto.Metric{
							Name:       MetricResidentMemorySize,
							Value:      psStat.Rss * 4096, // 4096 - memory page size
							Attributes: attrs,
						},
					)
				}
			}
		}
	}

	// Total
	for psName, st := range appStat {
		attrs := []dto.MetricAttribute{
			{
				Name:  dto.ResourceAttr,
				Value: ResourceName,
			},
			{
				Name:  SetNameProcessName,
				Value: psName,
			},
			{
				Name:  SetNamePID,
				Value: SetValuePIDTotal,
			},
		}

		metrics = append(metrics,
			dto.Metric{
				Name:       MetricUser,
				Value:      st.uTime,
				Attributes: attrs,
			},
			dto.Metric{
				Name:       MetricSystem,
				Value:      st.sTime,
				Attributes: attrs,
			},
			dto.Metric{
				Name:       MetricGuest,
				Value:      st.guestTime,
				Attributes: attrs,
			},
			dto.Metric{
				Name:       MetricTotal,
				Value:      st.totalTime,
				Attributes: attrs,
			},
			dto.Metric{
				Name:       MetricVirtualMemorySize,
				Value:      st.vSize,
				Attributes: attrs,
			},
			dto.Metric{
				Name:       MetricResidentMemorySize,
				Value:      st.rss * 4096, // 4096 - memory page size
				Attributes: attrs,
			},
			dto.Metric{
				Name:       MetricThreads,
				Value:      st.threads,
				Attributes: attrs,
			},
			dto.Metric{
				Name:       MetricProcesses,
				Value:      st.processCount,
				Attributes: attrs,
			},
		)
	}

	return metrics, nil
}
