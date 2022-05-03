package conf

import "github.com/xray-team/xray-agent-linux/dto"

// AgentConf defines configuration object.
type AgentConf struct {
	Flags              *Flags
	TimeZoneOffset     int8
	TimeZoneName       string
	GetStatIntervalSec int                   `json:"getStatIntervalSec" validate:"required,min=5,max=3600"`
	HostAttributes     []dto.MetricAttribute `json:"hostAttributes" validate:"dive"`
}

// Flags defines configuration passed by flags.
type Flags struct {
	ConfigFilePath *string
	DryRun         *bool
}
