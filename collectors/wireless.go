package collectors

import (
	"errors"
	"strings"

	"github.com/xray-team/xray-agent-linux/conf"
	"github.com/xray-team/xray-agent-linux/dto"
	"github.com/xray-team/xray-agent-linux/logger"
)

type WirelessDataSource interface {
	GetInterfaceData(ifName string) (*dto.Iwconfig, error)
}

type WirelessCollector struct {
	Config             *conf.WirelessConf
	DataSource         WirelessDataSource
	ClassNetDataSource ClassNetDataSource
}

// NewWirelessCollector returns a new collector object.
func NewWirelessCollector(cfg *conf.CollectorsConf, wirelessDataSource WirelessDataSource, classNetDataSource ClassNetDataSource) dto.Collector {
	if cfg == nil || wirelessDataSource == nil || classNetDataSource == nil {
		logger.LogWarning(logger.CollectorInitPrefix, errors.New("wireless collector init params error"))
		return nil
	}

	// exit if collector disabled
	if cfg.Wireless == nil || !cfg.Wireless.Enabled {
		return nil
	}

	return &WirelessCollector{
		Config:             cfg.Wireless,
		DataSource:         wirelessDataSource,
		ClassNetDataSource: classNetDataSource,
	}
}

// GetName returns the collector's name.
func (c *WirelessCollector) GetName() string {
	return dto.CollectorNameWireless
}

// Collect collects and returns metrics.
func (c *WirelessCollector) Collect() ([]dto.Metric, error) {
	// Inventory
	inventory, err := c.ClassNetDataSource.GetData()
	if err != nil {
		return nil, err
	}

	inventory = c.filterWireless(inventory)

	// Slice for results
	metrics := make([]dto.Metric, 0, 12*len(inventory))

	// fill out
	for ifName := range inventory {
		wireless, err := c.DataSource.GetInterfaceData(ifName)
		if err != nil {
			return nil, err
		}

		metrics = append(metrics, genMetricsNetDevWireless(strings.ReplaceAll(ifName, ".", "_"), wireless)...)
	}

	return metrics, nil
}

func (c *WirelessCollector) filterWireless(m map[string]dto.ClassNet) map[string]dto.ClassNet {
	out := make(map[string]dto.ClassNet)

	for devName, dev := range m {
		// Exclude not Wireless
		if !dev.IsWireless() {
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

func genMetricsNetDevWireless(ifName string, iwconfig *dto.Iwconfig) []dto.Metric {
	attrs := []dto.MetricAttribute{
		{
			Name:  dto.ResourceAttr,
			Value: dto.ResourceWireless,
		},
		{
			Name:  dto.SetNameWirelessInterface,
			Value: ifName,
		},
	}

	return []dto.Metric{
		{
			Name:       dto.MetricWirelessFrequency,
			Value:      iwconfig.Frequency,
			Attributes: attrs,
		},
		{
			Name:       dto.MetricWirelessBitRate,
			Value:      iwconfig.BitRate,
			Attributes: attrs,
		},
		{
			Name:       dto.MetricWirelessTxPower,
			Value:      iwconfig.TxPower,
			Attributes: attrs,
		},
		{
			Name:       dto.MetricWirelessLinkQuality,
			Value:      iwconfig.LinkQuality,
			Attributes: attrs,
		},
		{
			Name:       dto.MetricWirelessLinkQualityLimit,
			Value:      iwconfig.LinkQualityLimit,
			Attributes: attrs,
		},
		{
			Name:       dto.MetricWirelessSignalLevel,
			Value:      iwconfig.SignalLevel,
			Attributes: attrs,
		},
		{
			Name:       dto.MetricWirelessRxInvalidNwid,
			Value:      iwconfig.RxInvalidNwid,
			Attributes: attrs,
		},
		{
			Name:       dto.MetricWirelessRxInvalidCrypt,
			Value:      iwconfig.RxInvalidCrypt,
			Attributes: attrs,
		},
		{
			Name:       dto.MetricWirelessRxInvalidFrag,
			Value:      iwconfig.RxInvalidFrag,
			Attributes: attrs,
		},
		{
			Name:       dto.MetricWirelessTxExcessiveRetries,
			Value:      iwconfig.TxExcessiveRetries,
			Attributes: attrs,
		},
		{
			Name:       dto.MetricWirelessInvalidMisc,
			Value:      iwconfig.InvalidMisc,
			Attributes: attrs,
		},
		{
			Name:       dto.MetricWirelessMissedBeacon,
			Value:      iwconfig.MissedBeacon,
			Attributes: attrs,
		},
	}
}
