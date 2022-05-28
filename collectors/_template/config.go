package template

import (
	"encoding/json"
	"github.com/go-playground/validator"
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

// Validate validates all Config fields.
func (config *Config) Validate() error {
	validate := validator.New()

	return validate.Struct(config)
}

// Parse Config from raw json.
func (config *Config) Parse(data []byte) error {
	return json.Unmarshal(data, &config)
}
