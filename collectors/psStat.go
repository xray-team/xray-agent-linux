package collectors

import (
	"errors"
	"strconv"
	"strings"

	"xray-agent-linux/conf"
	"xray-agent-linux/dto"
	"xray-agent-linux/logger"
)

type PSStatDataSource interface {
	GetData() (*dto.PSStat, error)
}

type PSStatCollector struct {
	Config     *conf.PSStatConf
	DataSource PSStatDataSource
}

func (c *PSStatCollector) GetName() string {
	return dto.CollectorNamePSStat
}

func NewPSStatCollector(cfg *conf.CollectorsConf, dataSource PSStatDataSource) dto.Collector {
	if cfg == nil || dataSource == nil {
		logger.LogWarning(logger.CollectorInitPrefix, errors.New("PSStat collector init params error"))
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

func (c *PSStatCollector) Collect() ([]dto.Metric, error) {
	psStats, err := c.DataSource.GetData()
	if err != nil {
		return nil, err
	}

	// Slice for results
	metrics := make([]dto.Metric, 0)
	// Temp map
	processCount := make(map[string]int)

	for _, psStat := range psStats.PS {
		for _, psName := range c.Config.ProcessList {
			if psName == strings.TrimRight(strings.TrimLeft(psStat.Name, "("), ")") {
				processCount[psName] = processCount[psName] + 1

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
						Value:      psStat.Rss * 4096,
						Attributes: attrs,
					},
				)
			}
		}
	}

	for psName, count := range processCount {
		metrics = append(metrics,
			dto.Metric{
				Name:  dto.MetricPSStatProcesses,
				Value: count,
				Attributes: []dto.MetricAttribute{
					{
						Name:  dto.ResourceAttr,
						Value: dto.ResourcePSStat,
					},
					{
						Name:  dto.SetNamePSStatProcessName,
						Value: psName,
					},
				},
			})
	}

	return metrics, nil
}
