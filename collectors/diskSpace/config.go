package diskSpace

import (
	"encoding/json"

	"github.com/go-playground/validator"
)

type Config struct {
	Enabled bool `json:"enabled"`
	// MonitoredFileSystemTypes is used by procMounts
	MonitoredFileSystemTypes []string `json:"fsTypes" validate:"dive,oneof=ext4 ext3 ext2 btrfs xfs jfs ufs zfs vfat squashfs fuseblk ntfs msdos hfs hfsplus"`
}

// NewConfig returns Config with default values.
func NewConfig() *Config {
	return &Config{
		Enabled: false,
		MonitoredFileSystemTypes: []string{
			"ext4",
			"ext3",
			"ext2",
			"xfs",
			"jfs",
			"btrfs",
		},
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
