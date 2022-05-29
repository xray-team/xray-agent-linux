package diskStat

import (
	"encoding/json"

	"github.com/go-playground/validator"
)

type Config struct {
	Enabled bool `json:"enabled"`
	// MonitoredDiskTypes - block dev major numbers https://www.kernel.org/doc/Documentation/admin-guide/devices.txt
	MonitoredDiskTypes []int64  `json:"diskTypes" validate:"required,dive,min=0,max=259"`
	ExcludePartitions  bool     `json:"excludePartitions"`
	ExcludeByName      []string `json:"excludeByName"`
}

// NewConfig returns Config with default values.
func NewConfig() *Config {
	return &Config{
		Enabled: false,
		MonitoredDiskTypes: []int64{
			8,  // SCSI disk devices (0-15)
			9,  // Metadisk (RAID) devices
			65, // SCSI disk devices (16-31)
			66, // SCSI disk devices (32-47)
			67, // SCSI disk devices (48-63)
			68, // SCSI disk devices (64-79)
		},
		ExcludePartitions: false,
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
