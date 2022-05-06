package mdStat

import (
	"strconv"
	"strings"

	"github.com/xray-team/xray-agent-linux/conf"
	"github.com/xray-team/xray-agent-linux/dto"
	"github.com/xray-team/xray-agent-linux/logger"
)

type MDStatDataSource interface {
	GetData() (*dto.MDStats, error)
}

type MDStatCollector struct {
	Config     *conf.MDStatConf
	DataSource MDStatDataSource
}

// NewMDStatCollector returns a new collector object.
func NewMDStatCollector(cfg *conf.CollectorsConf, dataSource MDStatDataSource) dto.Collector {
	if cfg == nil || dataSource == nil {
		logger.Log.Error.Printf(logger.MessageInitCollectorError, dto.CollectorNameMDStat)
		return nil
	}

	// exit if collector disabled
	if cfg.MDStat == nil || !cfg.MDStat.Enabled {
		return nil
	}

	return &MDStatCollector{
		Config:     cfg.MDStat,
		DataSource: dataSource,
	}
}

// GetName returns the collector's name.
func (c *MDStatCollector) GetName() string {
	return dto.CollectorNameMDStat
}

// Collect collects and returns metrics.
func (c *MDStatCollector) Collect() ([]dto.Metric, error) {
	mdStat, err := c.DataSource.GetData()
	if err != nil {
		return nil, err
	}

	metrics := make([]dto.Metric, 0, len(mdStat.Stats)*16)

	resourceAttr := dto.MetricAttribute{
		Name:  dto.ResourceAttr,
		Value: dto.ResourceMDStat,
	}

	for mdName, stat := range mdStat.Stats {
		mdNameAttr := dto.MetricAttribute{
			Name:  dto.SetNameMDStatMD,
			Value: mdName,
		}

		// prepare level
		level, _ := strconv.Atoi(strings.TrimPrefix(strings.ToLower(stat.Level), "raid"))

		metrics = append(metrics,
			dto.Metric{
				Name:       dto.MetricMDStatLevel,
				Value:      level,
				Attributes: []dto.MetricAttribute{resourceAttr, mdNameAttr},
			},
			dto.Metric{
				Name:       dto.MetricMDStatNumDisks,
				Value:      stat.NumDisks,
				Attributes: []dto.MetricAttribute{resourceAttr, mdNameAttr},
			},
			dto.Metric{
				Name:       dto.MetricMDStatArrayState,
				Value:      MDStatsArrayStates[stat.ArrayState],
				Attributes: []dto.MetricAttribute{resourceAttr, mdNameAttr},
			},
			dto.Metric{
				Name:       dto.MetricMDStatArraySize,
				Value:      stat.ArraySizeKBytes,
				Attributes: []dto.MetricAttribute{resourceAttr, mdNameAttr},
			})
		// RAID with redundancy
		if stat.StatRaidWithRedundancy != nil {
			metrics = append(metrics,
				dto.Metric{
					Name:       dto.MetricMDStatSyncAction,
					Value:      MDStatsSyncActions[stat.StatRaidWithRedundancy.SyncAction],
					Attributes: []dto.MetricAttribute{resourceAttr, mdNameAttr},
				},
				dto.Metric{
					Name:       dto.MetricMDStatNumDegraded,
					Value:      stat.StatRaidWithRedundancy.NumDegraded,
					Attributes: []dto.MetricAttribute{resourceAttr, mdNameAttr},
				},
				dto.Metric{
					Name:       dto.MetricMDStatMismatchCnt,
					Value:      stat.StatRaidWithRedundancy.MismatchCnt,
					Attributes: []dto.MetricAttribute{resourceAttr, mdNameAttr},
				},
				dto.Metric{
					Name:       dto.MetricMDStatSyncCompletedSectors,
					Value:      stat.StatRaidWithRedundancy.SyncCompletedSectors,
					Attributes: []dto.MetricAttribute{resourceAttr, mdNameAttr},
				},
				dto.Metric{
					Name:       dto.MetricMDStatNumSectors,
					Value:      stat.StatRaidWithRedundancy.NumSectors,
					Attributes: []dto.MetricAttribute{resourceAttr, mdNameAttr},
				},
				dto.Metric{
					Name:       dto.MetricMDStatSyncSpeed,
					Value:      stat.StatRaidWithRedundancy.SyncSpeed,
					Attributes: []dto.MetricAttribute{resourceAttr, mdNameAttr},
				},
			)
		}
		// Dev Stats
		for devName, devStat := range stat.DevStats {
			devNameAttr := dto.MetricAttribute{
				Name:  dto.SetNameMDStatDev,
				Value: devName,
			}

			// Slot
			//   "none" if device is spare
			if devStat.Slot == "none" {
				metrics = append(metrics,
					dto.Metric{
						Name:       dto.MetricMDStatDevSlot,
						Value:      -1,
						Attributes: []dto.MetricAttribute{resourceAttr, mdNameAttr, devNameAttr},
					},
				)
			} else {
				metrics = append(metrics,
					dto.Metric{
						Name:       dto.MetricMDStatDevSlot,
						Value:      devStat.Slot,
						Attributes: []dto.MetricAttribute{resourceAttr, mdNameAttr, devNameAttr},
					},
				)
			}

			metrics = append(metrics,
				// State
				dto.Metric{
					Name:       dto.MetricMDStatDevState,
					Value:      MDStatsDevStates[devStat.State],
					Attributes: []dto.MetricAttribute{resourceAttr, mdNameAttr, devNameAttr},
				},
				dto.Metric{
					Name:       dto.MetricMDStatDevErrors,
					Value:      devStat.Errors,
					Attributes: []dto.MetricAttribute{resourceAttr, mdNameAttr, devNameAttr},
				},
			)
		}
	}

	return metrics, nil
}
