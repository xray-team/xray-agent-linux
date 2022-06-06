package dto_test

import (
	"fmt"
	"testing"

	"github.com/xray-team/xray-agent-linux/collectors/loadAvg"
	"github.com/xray-team/xray-agent-linux/collectors/netDev"
	"github.com/xray-team/xray-agent-linux/collectors/stat"
	"github.com/xray-team/xray-agent-linux/dto"
)

func TestMetric_GenGraphiteTreeName(t *testing.T) {
	type fields struct {
		Name       string
		Attributes []dto.MetricAttribute
		Value      interface{}
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "loadAvg",
			fields: fields{
				Name: loadAvg.MetricLast,
				Attributes: []dto.MetricAttribute{
					{
						Name:  dto.ResourceAttr,
						Value: loadAvg.ResourceName,
					},
				},
				Value: 1,
			},
			want: fmt.Sprintf("%s.%s", loadAvg.ResourceName, loadAvg.MetricLast),
		},
		{
			name: "cpu",
			fields: fields{
				Name: stat.MetricSoftIRQTimer,
				Attributes: []dto.MetricAttribute{
					{
						Name:  dto.ResourceAttr,
						Value: stat.ResourceName,
					},
					{
						Name:  stat.SetNameProcessor,
						Value: stat.SetValueProcessorTotal,
					},
					{
						Name:  stat.SetNameCPUSet,
						Value: stat.SetValueCPUSetSoftIRQ,
					},
				},
				Value: 1,
			},
			want: fmt.Sprintf("%s.%s.%s.%s", stat.ResourceName, stat.SetValueProcessorTotal, stat.SetValueCPUSetSoftIRQ, stat.MetricSoftIRQTimer),
		},
		{
			name: "netDev",
			fields: fields{
				Name: netDev.MetricStatisticsRxBytes,
				Attributes: []dto.MetricAttribute{
					{
						Name:  dto.ResourceAttr,
						Value: netDev.ResourceName,
					},
					{
						Name:  netDev.SetNameInterface,
						Value: "eth0",
					},
				},
				Value: 1,
			},
			want: fmt.Sprintf("%s.eth0.%s", netDev.ResourceName, netDev.MetricStatisticsRxBytes),
		},
	}

	for _, testCase := range tests {
		tt := testCase

		t.Run(tt.name, func(t *testing.T) {
			m := &dto.Metric{
				Name:       tt.fields.Name,
				Attributes: tt.fields.Attributes,
				Value:      tt.fields.Value,
			}

			if got := m.GenGraphiteTreeName(); got != tt.want {
				t.Errorf("GenGraphiteTreeName() = %v, want %v", got, tt.want)
			}
		})
	}
}
