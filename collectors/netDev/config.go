package netDev

import (
	"encoding/json"

	"github.com/go-playground/validator"
)

type Config struct {
	Enabled          bool `json:"enabled"`
	ExcludeLoopbacks bool `json:"excludeLoopbacks"`
	ExcludeWireless  bool `json:"excludeWireless"`
	ExcludeBridges   bool `json:"excludeBridges"`
	// Virtual interfaces are network interfaces that are not associated with an any physical interface.
	// Virtual interface examples: loopback, bridge, tun, vlan, ...
	ExcludeVirtual     bool     `json:"excludeVirtual"`
	ExcludeByName      []string `json:"excludeByName"`
	ExcludeByOperState []string `json:"excludeByOperState" validate:"dive,oneof=unknown notpresent down lowerlayerdown testing dormant up"`
}

// NewConfig returns Config with default values.
func NewConfig() *Config {
	return &Config{
		Enabled:          false,
		ExcludeLoopbacks: true,
		ExcludeWireless:  false,
		ExcludeBridges:   false,
		ExcludeVirtual:   false,
	}
}

// Validate validates all Config fields.
func (config *Config) Validate() error {
	validate := validator.New()

	return validate.Struct(config)
}

// Parse Config from raw json
func (config *Config) Parse(data []byte) error {
	return json.Unmarshal(data, &config)
}
