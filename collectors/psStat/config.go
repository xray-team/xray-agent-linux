package psStat

import (
	"encoding/json"
)

type Config struct {
	Enabled           bool     `json:"enabled"`
	CollectPerPidStat bool     `json:"collectPerPidStat"`
	ProcessList       []string `json:"processList"`
}

// NewConfig returns Config with default values.
func NewConfig() *Config {
	return &Config{
		Enabled:           false,
		CollectPerPidStat: false,
	}
}

// Parse Config from raw json.
func (config *Config) Parse(data []byte) error {
	return json.Unmarshal(data, &config)
}
