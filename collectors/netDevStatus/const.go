package netDevStatus

const (
	CollectorName = "NetDevStatus"
)

// Metrics
const (
	ResourceName     = "NetDevStatus"
	SetNameInterface = "Interface"

	MetricOperState = "OperState"
	MetricLinkFlaps = "LinkFlaps"
	MetricSpeed     = "Speed"
	MetricDuplex    = "Duplex"
	MetricMTU       = "MTU"
)

// RFC2863
var NetDevOperStates = map[string]int{
	"up":             1,
	"lowerlayerdown": 2,
	"dormant":        3,
	"down":           4,
	"unknown":        5,
	"testing":        6,
	"notpresent":     7,
}

var NetDevDuplexStates = map[string]int{
	"full":    1,
	"half":    2,
	"unknown": 3,
}
