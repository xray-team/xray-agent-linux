package netDevStatus

import (
	"encoding/json"

	"github.com/go-playground/validator"
)

type Config struct {
	Enabled         bool     `json:"enabled"`
	ExcludeWireless bool     `json:"excludeWireless"`
	ExcludeByName   []string `json:"excludeByName"`
}

// NewConfig returns Config with default values.
func NewConfig() *Config {
	return &Config{
		Enabled:         false,
		ExcludeWireless: true,
	}
}

// Validate validates all Config fields.
func (config *Config) Validate() error {
	validate := validator.New()

	return validate.Struct(config)
}

// Parse Config from raw json.
func (config *Config) Parse(data []byte) error {
	return json.Unmarshal(data, &config)
}
