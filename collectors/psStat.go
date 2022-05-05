package collectors

import (
	"strconv"
	"strings"

	"github.com/xray-team/xray-agent-linux/conf"
	"github.com/xray-team/xray-agent-linux/dto"
	"github.com/xray-team/xray-agent-linux/logger"
)

type PSStatDataSource interface {
	GetData() (*dto.PSStat, error)
}

type PSStatCollector struct {
	Config     *conf.PSStatConf
	DataSource PSStatDataSource
}

// GetName returns the collector's name.
func (c *PSStatCollector) GetName() string {
	return dto.CollectorNamePSStat
}

// NewPSStatCollector returns a new collector object.
func NewPSStatCollector(cfg *conf.CollectorsConf, dataSource PSStatDataSource) dto.Collector {
	if cfg == nil || dataSource == nil {
		logger.Log.Error.Printf(logger.MessageInitCollectorError, dto.CollectorNamePSStat)

		return nil
	}

	// exit if collector disabled
	if cfg.PSStat == nil || !cfg.PS.Enabled || len(cfg.PSStat.ProcessList) == 0 {
		return nil
	}

	return &PSStatCollector{
		Config:     cfg.PSStat,
		DataSource: dataSource,
	}
}

// Collect collects and returns metrics.
func (c *PSStatCollector) Collect() ([]dto.Metric, error) {
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
							Value: dto.ResourcePSStat,
						},
						{
							Name:  dto.SetNamePSStatProcessName,
							Value: psName,
						},
						{
							Name:  dto.SetNamePSStatPID,
							Value: strconv.FormatInt(psStat.PID, 10),
						},
					}

					metrics = append(metrics,
						dto.Metric{
							Name:       dto.MetricPSStatUser,
							Value:      psStat.UTime + psStat.CuTime,
							Attributes: attrs,
						},
						dto.Metric{
							Name:       dto.MetricPSStatSystem,
							Value:      psStat.STime + psStat.CsTime,
							Attributes: attrs,
						},
						dto.Metric{
							Name:       dto.MetricPSStatGuest,
							Value:      psStat.GuestTime + psStat.CGuestTime,
							Attributes: attrs,
						},
						dto.Metric{
							Name:       dto.MetricPSStatTotal,
							Value:      psStat.UTime + psStat.CuTime + psStat.STime + psStat.CsTime + psStat.GuestTime + psStat.CGuestTime,
							Attributes: attrs,
						},
						dto.Metric{
							Name:       dto.MetricPSStatThreads,
							Value:      psStat.Threads,
							Attributes: attrs,
						},
						dto.Metric{
							Name:       dto.MetricPSStatVirtualMemorySize,
							Value:      psStat.VSize,
							Attributes: attrs,
						},
						dto.Metric{
							Name:       dto.MetricPSStatResidentMemorySize,
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
				Value: dto.ResourcePSStat,
			},
			{
				Name:  dto.SetNamePSStatProcessName,
				Value: psName,
			},
			{
				Name:  dto.SetNamePSStatPID,
				Value: dto.SetValuePSStatPIDTotal,
			},
		}

		metrics = append(metrics,
			dto.Metric{
				Name:       dto.MetricPSStatUser,
				Value:      st.uTime,
				Attributes: attrs,
			},
			dto.Metric{
				Name:       dto.MetricPSStatSystem,
				Value:      st.sTime,
				Attributes: attrs,
			},
			dto.Metric{
				Name:       dto.MetricPSStatGuest,
				Value:      st.guestTime,
				Attributes: attrs,
			},
			dto.Metric{
				Name:       dto.MetricPSStatTotal,
				Value:      st.totalTime,
				Attributes: attrs,
			},
			dto.Metric{
				Name:       dto.MetricPSStatVirtualMemorySize,
				Value:      st.vSize,
				Attributes: attrs,
			},
			dto.Metric{
				Name:       dto.MetricPSStatResidentMemorySize,
				Value:      st.rss * 4096, // 4096 - memory page size
				Attributes: attrs,
			},
			dto.Metric{
				Name:       dto.MetricPSStatThreads,
				Value:      st.threads,
				Attributes: attrs,
			},
			dto.Metric{
				Name:       dto.MetricPSStatProcesses,
				Value:      st.processCount,
				Attributes: attrs,
			},
		)
	}

	return metrics, nil
}
