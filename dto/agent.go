package dto

import (
	"time"
)

type Collector interface {
	Collect() ([]Metric, error)
	GetName() string
}

type AgentSummary struct {
	Duration      time.Duration
	MetricsNumber int
}

type CollectorSummary struct {
	CollectorName string
	MetricsNumber int
	Duration      time.Duration
	// 0 - unknown
	// 1 - success
	// 2 - error
	Status int
}
