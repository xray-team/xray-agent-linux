package proc_test

import (
	"reflect"
	"testing"

	"xray-agent-linux/dto"
	"xray-agent-linux/logger"
	"xray-agent-linux/proc"
)

func Test_loadAvgDataSource_GetData(t *testing.T) {
	logger.Init("")

	tests := []struct {
		caseDescription string
		filePath        string
		want            *dto.LoadAvg
		wantErr         bool
	}{
		{
			caseDescription: "no file",
			filePath:        "./testfiles/loadavg/nofile",
			want:            nil,
			wantErr:         true,
		},
		{
			caseDescription: "desktop",
			filePath:        "./testfiles/loadavg/loadavg-desktop",
			want: &dto.LoadAvg{
				Last:                     1.61,
				Last5m:                   2.11,
				Last15m:                  2.26,
				KernelSchedulingEntities: 1034,
			},
			wantErr: false,
		},
		{
			caseDescription: "server",
			filePath:        "./testfiles/loadavg/loadavg-server",
			want: &dto.LoadAvg{
				Last:                     43.07,
				Last5m:                   44.18,
				Last15m:                  48.25,
				KernelSchedulingEntities: 11846,
			},
			wantErr: false,
		},
		{
			caseDescription: "server-unused",
			filePath:        "./testfiles/loadavg/loadavg-server-unused",
			want: &dto.LoadAvg{
				Last:                     0.0,
				Last5m:                   0.0,
				Last15m:                  0.0,
				KernelSchedulingEntities: 191,
			},
			wantErr: false,
		},
	}

	for _, testCase := range tests {
		tt := testCase

		t.Run(tt.caseDescription, func(t *testing.T) {
			loadAvgDataSource := proc.NewLoadAvgDataSource(tt.filePath, "")
			got, err := loadAvgDataSource.GetData()

			if (err != nil) != tt.wantErr {
				t.Errorf("GetData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetData() got = %v, want %v", got, tt.want)
			}
		})
	}
}
