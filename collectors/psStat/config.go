package psStat

type Config struct {
	Enabled           bool     `json:"enabled"`
	CollectPerPidStat bool     `json:"collectPerPidStat"`
	ProcessList       []string `json:"processList"`
}
