package nginx

import (
	"encoding/json"

	"github.com/go-playground/validator"
)

type Config struct {
	Enabled  bool   `json:"enabled"`
	Endpoint string `json:"endpoint" validate:"required"`
	Timeout  int    `json:"timeout" validate:"required,min=1,max=120"`
}

// NewConfig returns Config with default values.
func NewConfig() *Config {
	return &Config{
		Enabled:  false,
		Endpoint: "http://127.0.0.1/basic_status",
		Timeout:  5,
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
