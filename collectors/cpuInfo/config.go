package cpuInfo

import (
	"encoding/json"
)

type Config struct {
	Enabled bool `json:"enabled"`
}

// NewConfig returns Config with default values.
func NewConfig() *Config {
	return &Config{
		Enabled: false,
	}
}

// Parse Config from raw json.
func (config *Config) Parse(data []byte) error {
	return json.Unmarshal(data, &config)
}
