package dto_test

import (
	"fmt"
	"testing"

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
				Name: dto.MetricLoadAvgLast,
				Attributes: []dto.MetricAttribute{
					{
						Name:  dto.ResourceAttr,
						Value: dto.ResourceLoadAvg,
					},
				},
				Value: 1,
			},
			want: fmt.Sprintf("%s.%s", dto.ResourceLoadAvg, dto.MetricLoadAvgLast),
		},
		{
			name: "cpu",
			fields: fields{
				Name: dto.MetricCPUSoftIRQTimer,
				Attributes: []dto.MetricAttribute{
					{
						Name:  dto.ResourceAttr,
						Value: dto.ResourceStat,
					},
					{
						Name:  dto.SetNameCPUProcessor,
						Value: dto.SetValueCPUProcessorTotal,
					},
					{
						Name:  dto.SetNameCPUSet,
						Value: dto.SetValueCPUSetSoftIRQ,
					},
				},
				Value: 1,
			},
			want: fmt.Sprintf("%s.%s.%s.%s", dto.ResourceStat, dto.SetValueCPUProcessorTotal, dto.SetValueCPUSetSoftIRQ, dto.MetricCPUSoftIRQTimer),
		},
		{
			name: "netDev",
			fields: fields{
				Name: dto.MetricNetDevStatisticsRxBytes,
				Attributes: []dto.MetricAttribute{
					{
						Name:  dto.ResourceAttr,
						Value: dto.ResourceNetDev,
					},
					{
						Name:  dto.SetNameNetDevInterface,
						Value: "eth0",
					},
				},
				Value: 1,
			},
			want: fmt.Sprintf("%s.eth0.%s", dto.ResourceNetDev, dto.MetricNetDevStatisticsRxBytes),
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
