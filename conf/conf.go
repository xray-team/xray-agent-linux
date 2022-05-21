package conf

import (
	"encoding/json"

	"github.com/xray-team/xray-agent-linux/logger"
	"github.com/xray-team/xray-agent-linux/reader"

	"github.com/go-playground/validator"
)

// Config defines configuration object.
type Config struct {
	Agent      *AgentConf                 `json:"agent" validate:"required"`
	TSDB       *TSDBConf                  `json:"tsDB" validate:"required"`
	Collectors map[string]json.RawMessage `json:"collectors" validate:"required"`
}

type TSDBConf struct {
	Graphite *GraphiteConf `json:"graphite" validate:"required"`
}

func GetConfiguration(flags *Flags) (*Config, error) {
	cfg, err := ReadConfigFile(*flags.ConfigFilePath)
	if err != nil {
		return nil, err
	}

	cfg.Agent.Flags = flags

	err = cfg.Validate()

	return cfg, err
}

func ReadConfigFile(filePath string) (*Config, error) {
	data, err := reader.ReadFile(filePath, logger.TagConfig)
	if err != nil {
		return nil, err
	}

	var out Config

	err = json.Unmarshal(data, &out)
	if err != nil {
		return nil, err
	}

	return &out, nil
}

// Validate validates all Config fields.
func (cfg *Config) Validate() (err error) {
	validate := validator.New()

	return validate.Struct(cfg)
}
