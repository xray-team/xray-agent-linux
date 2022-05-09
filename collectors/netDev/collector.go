package netDev

import (
	"strings"

	"github.com/xray-team/xray-agent-linux/dto"
	"github.com/xray-team/xray-agent-linux/logger"
)

type NetDevDataSource interface {
	GetData() (map[string]NetDevStatistics, error)
}

type ClassNetDataSource interface {
	GetData() (map[string]dto.ClassNet, error)
}

type Collector struct {
	Config             *Config
	DataSource         NetDevDataSource
	ClassNetDataSource ClassNetDataSource
}

// NewCollector returns a new collector object.
func NewCollector(config *Config, netDevDataSource NetDevDataSource, classNetDataSource ClassNetDataSource) dto.Collector {
	if config == nil || netDevDataSource == nil || classNetDataSource == nil {
		logger.Log.Error.Printf(logger.MessageInitCollectorError, CollectorName)

		return nil
	}

	// exit if collector disabled
	if !config.Enabled {
		return nil
	}

	return &Collector{
		Config:             config,
		DataSource:         netDevDataSource,
		ClassNetDataSource: classNetDataSource,
	}
}

// GetName returns the collector's name.
func (c *Collector) GetName() string {
	return CollectorName
}

// Collect collects and returns metrics.
func (c *Collector) Collect() ([]dto.Metric, error) {
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

func (c *Collector) filterNetDev(m map[string]dto.ClassNet) map[string]dto.ClassNet {
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

func genMetricsNetDevStatistics(ifName string, statistics NetDevStatistics) []dto.Metric {
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
			Name:       MetricStatisticsRxBytes,
			Value:      statistics.RxBytes,
			Attributes: attrs,
		},
		{
			Name:       MetricStatisticsRxPackets,
			Value:      statistics.RxPackets,
			Attributes: attrs,
		},
		{
			Name:       MetricStatisticsRxErrs,
			Value:      statistics.RxErrs,
			Attributes: attrs,
		},
		{
			Name:       MetricStatisticsRxDrop,
			Value:      statistics.RxDrop,
			Attributes: attrs,
		},
		{
			Name:       MetricStatisticsRxFifoErrs,
			Value:      statistics.RxFifoErrs,
			Attributes: attrs,
		},
		{
			Name:       MetricStatisticsRxFrameErrs,
			Value:      statistics.RxFrameErrs,
			Attributes: attrs,
		},
		{
			Name:       MetricStatisticsRxCompressed,
			Value:      statistics.RxCompressed,
			Attributes: attrs,
		},
		{
			Name:       MetricStatisticsMulticast,
			Value:      statistics.Multicast,
			Attributes: attrs,
		},
		{
			Name:       MetricStatisticsTxBytes,
			Value:      statistics.TxBytes,
			Attributes: attrs,
		},
		{
			Name:       MetricStatisticsTxPackets,
			Value:      statistics.TxPackets,
			Attributes: attrs,
		},
		{
			Name:       MetricStatisticsTxErrs,
			Value:      statistics.TxErrs,
			Attributes: attrs,
		},
		{
			Name:       MetricStatisticsTxDrop,
			Value:      statistics.TxDrop,
			Attributes: attrs,
		},
		{
			Name:       MetricStatisticsTxFifoErrs,
			Value:      statistics.TxFifoErrs,
			Attributes: attrs,
		},
		{
			Name:       MetricStatisticsCollisions,
			Value:      statistics.Collisions,
			Attributes: attrs,
		},
		{
			Name:       MetricStatisticsTxCarrierErrs,
			Value:      statistics.TxCarrierErrs,
			Attributes: attrs,
		},
		{
			Name:       MetricStatisticsTxCompressed,
			Value:      statistics.TxCompressed,
			Attributes: attrs,
		},
	}
}
