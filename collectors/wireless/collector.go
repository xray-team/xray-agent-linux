package wireless

import (
	"github.com/xray-team/xray-agent-linux/run"
	"github.com/xray-team/xray-agent-linux/sys"
	"strings"

	"github.com/xray-team/xray-agent-linux/dto"
	"github.com/xray-team/xray-agent-linux/logger"
)

type WirelessDataSource interface {
	GetInterfaceData(ifName string) (*Iwconfig, error)
}

type ClassNetDataSource interface {
	GetData() (map[string]dto.ClassNet, error)
}

type Collector struct {
	Config             *Config
	DataSource         WirelessDataSource
	ClassNetDataSource ClassNetDataSource
}

// CreateCollector returns a new collector object.
func CreateCollector(rawConfig []byte) dto.Collector {
	config := NewConfig()

	err := config.Parse(rawConfig)
	if err != nil {
		logger.Log.Error.Printf(logger.MessageError, CollectorName, err.Error())

		return nil
	}

	err = config.Validate()
	if err != nil {
		logger.Log.Error.Printf(logger.MessageError, CollectorName, err.Error())

		return nil
	}

	return NewCollector(
		config,
		NewIwconfigDataSource(run.NewCmdRunner(CollectorName)),
		sys.NewClassNetDataSource(sys.ClassNetDir, CollectorName),
	)
}

// NewCollector returns a new collector object.
func NewCollector(config *Config, wirelessDataSource WirelessDataSource, classNetDataSource ClassNetDataSource) dto.Collector {
	if config == nil || wirelessDataSource == nil || classNetDataSource == nil {
		logger.Log.Error.Printf(logger.MessageInitCollectorError, CollectorName)

		return nil
	}

	// exit if collector disabled
	if !config.Enabled {
		return nil
	}

	return &Collector{
		Config:             config,
		DataSource:         wirelessDataSource,
		ClassNetDataSource: classNetDataSource,
	}
}

// GetName returns the collector's name.
func (c *Collector) GetName() string {
	return CollectorName
}

// Collect collects and returns metrics.
func (c *Collector) Collect() ([]dto.Metric, error) {
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

func (c *Collector) filterWireless(m map[string]dto.ClassNet) map[string]dto.ClassNet {
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

func genMetricsNetDevWireless(ifName string, iwconfig *Iwconfig) []dto.Metric {
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
			Name:       MetricWirelessFrequency,
			Value:      iwconfig.Frequency,
			Attributes: attrs,
		},
		{
			Name:       MetricWirelessBitRate,
			Value:      iwconfig.BitRate,
			Attributes: attrs,
		},
		{
			Name:       MetricWirelessTxPower,
			Value:      iwconfig.TxPower,
			Attributes: attrs,
		},
		{
			Name:       MetricWirelessLinkQuality,
			Value:      iwconfig.LinkQuality,
			Attributes: attrs,
		},
		{
			Name:       MetricWirelessLinkQualityLimit,
			Value:      iwconfig.LinkQualityLimit,
			Attributes: attrs,
		},
		{
			Name:       MetricWirelessSignalLevel,
			Value:      iwconfig.SignalLevel,
			Attributes: attrs,
		},
		{
			Name:       MetricWirelessRxInvalidNwid,
			Value:      iwconfig.RxInvalidNwid,
			Attributes: attrs,
		},
		{
			Name:       MetricWirelessRxInvalidCrypt,
			Value:      iwconfig.RxInvalidCrypt,
			Attributes: attrs,
		},
		{
			Name:       MetricWirelessRxInvalidFrag,
			Value:      iwconfig.RxInvalidFrag,
			Attributes: attrs,
		},
		{
			Name:       MetricWirelessTxExcessiveRetries,
			Value:      iwconfig.TxExcessiveRetries,
			Attributes: attrs,
		},
		{
			Name:       MetricWirelessInvalidMisc,
			Value:      iwconfig.InvalidMisc,
			Attributes: attrs,
		},
		{
			Name:       MetricWirelessMissedBeacon,
			Value:      iwconfig.MissedBeacon,
			Attributes: attrs,
		},
	}
}
