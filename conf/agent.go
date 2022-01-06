package conf

import "github.com/xray-team/xray-agent-linux/dto"

// Config defines configuration object.
type AgentConf struct {
	Flags              *Flags
	TimeZoneOffset     int8
	TimeZoneName       string
	GetStatIntervalSec int                   `json:"getStatIntervalSec" validate:"required,min=5,max=3600"`
	HostAttributes     []dto.MetricAttribute `json:"hostAttributes" validate:"dive"`
}

type Flags struct {
	ConfigFilePath *string
	DryRun         *bool
}
