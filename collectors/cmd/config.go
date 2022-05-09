package cmd

import "github.com/xray-team/xray-agent-linux/dto"

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
