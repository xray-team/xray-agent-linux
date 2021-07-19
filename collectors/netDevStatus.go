package collectors

import (
	"errors"
	"fmt"
	"strings"

	"xray-agent-linux/conf"
	"xray-agent-linux/dto"
	"xray-agent-linux/logger"
)

type ClassNetStatusDataSource interface {
	GetData() (map[string]dto.ClassNet, error)
}

type NetDevStatusCollector struct {
	Config             *conf.NetDevStatusConf
	ClassNetDataSource ClassNetDataSource
}

func NewNetDevStatusCollector(cfg *conf.CollectorsConf, classNetDataSource ClassNetDataSource) dto.Collector {
	if cfg == nil || classNetDataSource == nil {
		logger.LogWarning(logger.CollectorInitPrefix, errors.New(fmt.Sprintf("%s collector init params error", dto.CollectorNameNetDevStatus)))
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

func (c *NetDevStatusCollector) GetName() string {
	return dto.CollectorNameNetDevStatus
}

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
			status := &dto.NetDevStatus{
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

func genMetricsNetDevStatus(ifName string, status *dto.NetDevStatus) []dto.Metric {
	attrs := []dto.MetricAttribute{
		{
			Name:  dto.ResourceAttr,
			Value: dto.ResourceNetDevStatus,
		},
		{
			Name:  dto.SetNameNetDevInterface,
			Value: ifName,
		},
	}

	return []dto.Metric{
		{
			Name:       dto.MetricNetDevStatusOperState,
			Value:      dto.NetDevOperStates[strings.ToLower(status.OperState)],
			Attributes: attrs,
		},
		{
			Name:       dto.MetricNetDevStatusSpeed,
			Value:      status.Speed,
			Attributes: attrs,
		},
		{
			Name:       dto.MetricNetDevStatusLinkFlaps,
			Value:      convertCarrierChangesToLinkFlaps(status.CarrierChanges),
			Attributes: attrs,
		},
		{
			Name:       dto.MetricNetDevStatusDuplex,
			Value:      dto.NetDevDuplexStates[strings.ToLower(status.Duplex)],
			Attributes: attrs,
		},
		{
			Name:       dto.MetricNetDevStatusMTU,
			Value:      status.MTU,
			Attributes: attrs,
		},
	}
}
