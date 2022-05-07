package wireless

// refer to man iwconfig
type Iwconfig struct {
	SSID      string
	Frequency float64
	BitRate   float64
	// An  address equal to "00:00:00:00:00:00" means that the card failed to associate with an Access Point
	// "Not-Associated" means that the device is off
	AccessPoint        string
	TxPower            int64
	LinkQuality        int64
	LinkQualityLimit   int64
	SignalLevel        int64
	RxInvalidNwid      uint64
	RxInvalidCrypt     uint64
	RxInvalidFrag      uint64
	TxExcessiveRetries uint64
	InvalidMisc        uint64
	MissedBeacon       uint64
}
