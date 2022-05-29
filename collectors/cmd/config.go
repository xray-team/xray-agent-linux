package cmd

import (
	"encoding/json"

	"github.com/go-playground/validator"

	"github.com/xray-team/xray-agent-linux/dto"
)

type Config struct {
	Enabled bool           `json:"enabled"`
	Timeout int            `json:"timeout" validate:"required,min=1,max=120"`
	Metrics []MetricConfig `json:"metrics" validate:"dive"`
}

type MetricConfig struct {
	Names      []string              `json:"names" validate:"required,min=1"`
	Delimiter  string                `json:"delimiter" validate:"required"`
	Attributes []dto.MetricAttribute `json:"attributes" validate:"dive"`
	PipeLine   [][]string            `json:"pipeline" validate:"required,min=1,dive,required,min=1"`
}

// NewConfig returns Config with default values.
func NewConfig() *Config {
	return &Config{
		Enabled: false,
		Timeout: 10,
		Metrics: nil,
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
