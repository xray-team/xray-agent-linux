package dto

import "fmt"

type Telemetry struct {
	HostInfo *HostInfo
	Metrics  []Metric
}

type HostInfo struct {
	HostName   string
	Timestamp  int64
	Attributes []MetricAttribute
}

type MetricAttribute struct {
	Name  string `json:"name" validate:"required"`
	Value string `json:"value" validate:"required"`
}

type Metric struct {
	Name       string
	Value      interface{}
	Attributes []MetricAttribute
}

func (m *Metric) GenGraphiteTreeName() string {
	return fmt.Sprintf("%s%s", GenGraphiteTreeString(m.Attributes), m.Name)
}

func GenGraphiteTreeString(attrs []MetricAttribute) string {
	var out string

	for _, attr := range attrs {
		out = fmt.Sprintf("%s%s.", out, attr.Value)
	}

	return out
}

func (m *Metric) GenGraphiteTagsName() string {
	out := m.Name

	for _, attr := range m.Attributes {
		out = fmt.Sprintf("%s;%s=%s", out, attr.Name, attr.Value)
	}

	return out
}
