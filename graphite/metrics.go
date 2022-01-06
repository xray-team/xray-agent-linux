package graphite

import (
	"fmt"

	"github.com/xray-team/xray-agent-linux/dto"

	"github.com/crazygreenpenguin/graphite"
)

func genGraphiteTreeMetrics(telemetry dto.Telemetry) []graphite.Metric {
	var out = make([]graphite.Metric, 0, len(telemetry.Metrics))

	for _, metric := range telemetry.Metrics {
		out = append(out, graphite.Metric{
			Name:      fmt.Sprintf("%s.%s", telemetry.HostInfo.HostName, metric.GenGraphiteTreeName()),
			Value:     metric.Value,
			Timestamp: telemetry.HostInfo.Timestamp,
		})
	}

	return out
}

func genGraphiteTagsMetrics(telemetry dto.Telemetry) []graphite.Metric {
	var out = make([]graphite.Metric, 0, len(telemetry.Metrics))

	for _, metric := range telemetry.Metrics {
		metric.Attributes = append(metric.Attributes, dto.MetricAttribute{
			Name:  "Hostname",
			Value: telemetry.HostInfo.HostName,
		})

		metric.Attributes = append(metric.Attributes, telemetry.HostInfo.Attributes...)

		out = append(out, graphite.Metric{
			Name:      metric.GenGraphiteTagsName(),
			Value:     metric.Value,
			Timestamp: telemetry.HostInfo.Timestamp,
		})
	}

	return out
}
