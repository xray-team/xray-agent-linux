package mdStat

import (
	"strconv"
	"strings"

	"github.com/xray-team/xray-agent-linux/dto"
	"github.com/xray-team/xray-agent-linux/logger"
)

type DataSource interface {
	GetData() (*MDStats, error)
}

type Collector struct {
	Config     *Config
	DataSource DataSource
}

// NewCollector returns a new collector object.
func NewCollector(config *Config, dataSource DataSource) dto.Collector {
	if config == nil || dataSource == nil {
		logger.Log.Error.Printf(logger.MessageInitCollectorError, CollectorName)
		return nil
	}

	// exit if collector disabled
	if !config.Enabled {
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
	mdStat, err := c.DataSource.GetData()
	if err != nil {
		return nil, err
	}

	metrics := make([]dto.Metric, 0, len(mdStat.Stats)*16)

	resourceAttr := dto.MetricAttribute{
		Name:  dto.ResourceAttr,
		Value: ResourceName,
	}

	for mdName, stat := range mdStat.Stats {
		mdNameAttr := dto.MetricAttribute{
			Name:  SetNameMD,
			Value: mdName,
		}

		// prepare level
		level, _ := strconv.Atoi(strings.TrimPrefix(strings.ToLower(stat.Level), "raid"))

		metrics = append(metrics,
			dto.Metric{
				Name:       MetricLevel,
				Value:      level,
				Attributes: []dto.MetricAttribute{resourceAttr, mdNameAttr},
			},
			dto.Metric{
				Name:       MetricNumDisks,
				Value:      stat.NumDisks,
				Attributes: []dto.MetricAttribute{resourceAttr, mdNameAttr},
			},
			dto.Metric{
				Name:       MetricArrayState,
				Value:      MDStatsArrayStates[stat.ArrayState],
				Attributes: []dto.MetricAttribute{resourceAttr, mdNameAttr},
			},
			dto.Metric{
				Name:       MetricArraySize,
				Value:      stat.ArraySizeKBytes,
				Attributes: []dto.MetricAttribute{resourceAttr, mdNameAttr},
			})
		// RAID with redundancy
		if stat.StatRaidWithRedundancy != nil {
			metrics = append(metrics,
				dto.Metric{
					Name:       MetricSyncAction,
					Value:      MDStatsSyncActions[stat.StatRaidWithRedundancy.SyncAction],
					Attributes: []dto.MetricAttribute{resourceAttr, mdNameAttr},
				},
				dto.Metric{
					Name:       MetricNumDegraded,
					Value:      stat.StatRaidWithRedundancy.NumDegraded,
					Attributes: []dto.MetricAttribute{resourceAttr, mdNameAttr},
				},
				dto.Metric{
					Name:       MetricMismatchCnt,
					Value:      stat.StatRaidWithRedundancy.MismatchCnt,
					Attributes: []dto.MetricAttribute{resourceAttr, mdNameAttr},
				},
				dto.Metric{
					Name:       MetricSyncCompletedSectors,
					Value:      stat.StatRaidWithRedundancy.SyncCompletedSectors,
					Attributes: []dto.MetricAttribute{resourceAttr, mdNameAttr},
				},
				dto.Metric{
					Name:       MetricNumSectors,
					Value:      stat.StatRaidWithRedundancy.NumSectors,
					Attributes: []dto.MetricAttribute{resourceAttr, mdNameAttr},
				},
				dto.Metric{
					Name:       MetricSyncSpeed,
					Value:      stat.StatRaidWithRedundancy.SyncSpeed,
					Attributes: []dto.MetricAttribute{resourceAttr, mdNameAttr},
				},
			)
		}
		// Dev Stats
		for devName, devStat := range stat.DevStats {
			devNameAttr := dto.MetricAttribute{
				Name:  SetNameDev,
				Value: devName,
			}

			// Slot
			//   "none" if device is spare
			if devStat.Slot == "none" {
				metrics = append(metrics,
					dto.Metric{
						Name:       MetricDevSlot,
						Value:      -1,
						Attributes: []dto.MetricAttribute{resourceAttr, mdNameAttr, devNameAttr},
					},
				)
			} else {
				metrics = append(metrics,
					dto.Metric{
						Name:       MetricDevSlot,
						Value:      devStat.Slot,
						Attributes: []dto.MetricAttribute{resourceAttr, mdNameAttr, devNameAttr},
					},
				)
			}

			metrics = append(metrics,
				// State
				dto.Metric{
					Name:       MetricDevState,
					Value:      MDStatsDevStates[devStat.State],
					Attributes: []dto.MetricAttribute{resourceAttr, mdNameAttr, devNameAttr},
				},
				dto.Metric{
					Name:       MetricDevErrors,
					Value:      devStat.Errors,
					Attributes: []dto.MetricAttribute{resourceAttr, mdNameAttr, devNameAttr},
				},
			)
		}
	}

	return metrics, nil
}
