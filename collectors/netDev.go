package collectors

import (
	"errors"
	"fmt"
	"strings"

	"xray-agent-linux/conf"
	"xray-agent-linux/dto"
	"xray-agent-linux/logger"
)

type NetDevDataSource interface {
	GetData() (map[string]dto.NetDevStatistics, error)
}

type ClassNetDataSource interface {
	GetData() (map[string]dto.ClassNet, error)
}

type NetDevCollector struct {
	Config             *conf.NetDevConf
	DataSource         NetDevDataSource
	ClassNetDataSource ClassNetDataSource
}

func NewNetDevCollector(cfg *conf.CollectorsConf, netDevDataSource NetDevDataSource, classNetDataSource ClassNetDataSource) dto.Collector {
	if cfg == nil || netDevDataSource == nil || classNetDataSource == nil {
		logger.LogWarning(logger.CollectorInitPrefix, errors.New(fmt.Sprintf("%s collector init params error", dto.CollectorNameNetDev)))
		return nil
	}

	// exit if collector disabled
	if cfg.NetDev == nil || !cfg.NetDev.Enabled {
		return nil
	}

	return &NetDevCollector{
		Config:             cfg.NetDev,
		DataSource:         netDevDataSource,
		ClassNetDataSource: classNetDataSource,
	}
}

func (c *NetDevCollector) GetName() string {
	return dto.CollectorNameNetDev
}

func (c *NetDevCollector) Collect() ([]dto.Metric, error) {
	// Net Dev Inventory
	inventory, err := c.ClassNetDataSource.GetData()
	if err != nil {
		return nil, err
	}

	// Inventory Filters
	inventory = c.filterNetDev(inventory)

	// Net Dev Data
	stat, err := c.DataSource.GetData()
	if err != nil {
		return nil, err
	}

	// Slice for results
	metrics := make([]dto.Metric, 0, 64)

	// fill out
	for ifName, st := range stat {
		statistics := st

		_, ok := inventory[ifName]
		// skip excluded interfaces (absent into inventory)
		if !ok {
			continue
		}

		metrics = append(metrics, genMetricsNetDevStatistics(strings.ReplaceAll(ifName, ".", "_"), statistics)...)
	}

	return metrics, nil
}

func (c *NetDevCollector) filterNetDev(m map[string]dto.ClassNet) map[string]dto.ClassNet {
	out := make(map[string]dto.ClassNet)

	for devName, dev := range m {
		// Exclude Loopback
		if c.Config.ExcludeLoopbacks && dev.IsLoopback() {
			continue
		}

		// Exclude Wireless
		if c.Config.ExcludeWireless && dev.IsWireless() {
			continue
		}

		// Exclude Bridge
		if c.Config.ExcludeBridges && dev.IsBridge() {
			continue
		}

		// Exclude Virtual
		if c.Config.ExcludeVirtual && dev.IsVirtual() {
			continue
		}

		// Exclude by OperState
		if func() bool {
			for _, operState := range c.Config.ExcludeByOperState {
				if dev.OperState == operState {
					return true
				}
			}
			return false
		}() {
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

func convertCarrierChangesToLinkFlaps(cc int64) int64 {
	if cc <= 3 {
		return 0
	}

	return cc/2 - 1
}

func genMetricsNetDevStatistics(ifName string, statistics dto.NetDevStatistics) []dto.Metric {
	attrs := []dto.MetricAttribute{
		{
			Name:  dto.ResourceAttr,
			Value: dto.ResourceNetDev,
		},
		{
			Name:  dto.SetNameNetDevInterface,
			Value: ifName,
		},
	}

	return []dto.Metric{
		{
			Name:       dto.MetricNetDevStatisticsRxBytes,
			Value:      statistics.RxBytes,
			Attributes: attrs,
		},
		{
			Name:       dto.MetricNetDevStatisticsRxPackets,
			Value:      statistics.RxPackets,
			Attributes: attrs,
		},
		{
			Name:       dto.MetricNetDevStatisticsRxErrs,
			Value:      statistics.RxErrs,
			Attributes: attrs,
		},
		{
			Name:       dto.MetricNetDevStatisticsRxDrop,
			Value:      statistics.RxDrop,
			Attributes: attrs,
		},
		{
			Name:       dto.MetricNetDevStatisticsRxFifoErrs,
			Value:      statistics.RxFifoErrs,
			Attributes: attrs,
		},
		{
			Name:       dto.MetricNetDevStatisticsRxFrameErrs,
			Value:      statistics.RxFrameErrs,
			Attributes: attrs,
		},
		{
			Name:       dto.MetricNetDevStatisticsRxCompressed,
			Value:      statistics.RxCompressed,
			Attributes: attrs,
		},
		{
			Name:       dto.MetricNetDevStatisticsMulticast,
			Value:      statistics.Multicast,
			Attributes: attrs,
		},
		{
			Name:       dto.MetricNetDevStatisticsTxBytes,
			Value:      statistics.TxBytes,
			Attributes: attrs,
		},
		{
			Name:       dto.MetricNetDevStatisticsTxPackets,
			Value:      statistics.TxPackets,
			Attributes: attrs,
		},
		{
			Name:       dto.MetricNetDevStatisticsTxErrs,
			Value:      statistics.TxErrs,
			Attributes: attrs,
		},
		{
			Name:       dto.MetricNetDevStatisticsTxDrop,
			Value:      statistics.TxDrop,
			Attributes: attrs,
		},
		{
			Name:       dto.MetricNetDevStatisticsTxFifoErrs,
			Value:      statistics.TxFifoErrs,
			Attributes: attrs,
		},
		{
			Name:       dto.MetricNetDevStatisticsCollisions,
			Value:      statistics.Collisions,
			Attributes: attrs,
		},
		{
			Name:       dto.MetricNetDevStatisticsTxCarrierErrs,
			Value:      statistics.TxCarrierErrs,
			Attributes: attrs,
		},
		{
			Name:       dto.MetricNetDevStatisticsTxCompressed,
			Value:      statistics.TxCompressed,
			Attributes: attrs,
		},
	}
}
