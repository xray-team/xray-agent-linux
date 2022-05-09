package netDevStatus

type Config struct {
	Enabled         bool     `json:"enabled"`
	ExcludeWireless bool     `json:"excludeWireless"`
	ExcludeByName   []string `json:"excludeByName"`
}
