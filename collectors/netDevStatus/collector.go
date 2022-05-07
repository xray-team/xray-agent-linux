package netDevStatus

import (
	"strings"

	"github.com/xray-team/xray-agent-linux/conf"
	"github.com/xray-team/xray-agent-linux/dto"
	"github.com/xray-team/xray-agent-linux/logger"
)

type ClassNetStatusDataSource interface {
	GetData() (map[string]dto.ClassNet, error)
}

type NetDevStatusCollector struct {
	Config             *conf.NetDevStatusConf
	ClassNetDataSource ClassNetStatusDataSource
}

// NewNetDevStatusCollector returns a new collector object.
func NewNetDevStatusCollector(cfg *conf.CollectorsConf, classNetDataSource ClassNetStatusDataSource) dto.Collector {
	if cfg == nil || classNetDataSource == nil {
		logger.Log.Error.Printf(logger.MessageInitCollectorError, CollectorName)

		return nil
	}

	// exit if collector disabled
	if cfg.NetDevStatus == nil || !cfg.NetDevStatus.Enabled {
		return nil
	}

	return &NetDevStatusCollector{
		Config:             cfg.NetDevStatus,
		ClassNetDataSource: classNetDataSource,
	}
}

// GetName returns the collector's name.
func (c *NetDevStatusCollector) GetName() string {
	return CollectorName
}

// Collect collects and returns metrics.
func (c *NetDevStatusCollector) Collect() ([]dto.Metric, error) {
	// Net Dev Inventory
	inventory, err := c.ClassNetDataSource.GetData()
	if err != nil {
		return nil, err
	}

	// Inventory Filters
	inventory = c.filterNetDev(inventory)

	// Slice for results
	metrics := make([]dto.Metric, 0, 64)

	// fill out
	for ifName, invItem := range inventory {
		if invItem.IsDevice() {
			status := &NetDevStatus{
				OperState:      invItem.OperState,
				Duplex:         invItem.Duplex,
				Speed:          invItem.Speed,
				MTU:            invItem.MTU,
				CarrierChanges: invItem.CarrierChanges,
			}

			metrics = append(metrics, genMetricsNetDevStatus(strings.ReplaceAll(ifName, ".", "_"), status)...)
		}
	}

	return metrics, nil
}

func (c *NetDevStatusCollector) filterNetDev(m map[string]dto.ClassNet) map[string]dto.ClassNet {
	out := make(map[string]dto.ClassNet)

	for devName, dev := range m {
		// Exclude Wireless
		if c.Config.ExcludeWireless && dev.IsWireless() {
			continue
		}

		out[devName] = dev
	}

	// Exclude by dev name
	for _, devName := range c.Config.ExcludeByName {
		delete(out, devName)
	}

	return out
}

func genMetricsNetDevStatus(ifName string, status *NetDevStatus) []dto.Metric {
	attrs := []dto.MetricAttribute{
		{
			Name:  dto.ResourceAttr,
			Value: ResourceName,
		},
		{
			Name:  SetNameInterface,
			Value: ifName,
		},
	}

	return []dto.Metric{
		{
			Name:       MetricOperState,
			Value:      NetDevOperStates[strings.ToLower(status.OperState)],
			Attributes: attrs,
		},
		{
			Name:       MetricSpeed,
			Value:      status.Speed,
			Attributes: attrs,
		},
		{
			Name:       MetricLinkFlaps,
			Value:      convertCarrierChangesToLinkFlaps(status.CarrierChanges),
			Attributes: attrs,
		},
		{
			Name:       MetricDuplex,
			Value:      NetDevDuplexStates[strings.ToLower(status.Duplex)],
			Attributes: attrs,
		},
		{
			Name:       MetricMTU,
			Value:      status.MTU,
			Attributes: attrs,
		},
	}
}

func convertCarrierChangesToLinkFlaps(cc int64) int64 {
	if cc <= 3 {
		return 0
	}

	return cc/2 - 1
}
