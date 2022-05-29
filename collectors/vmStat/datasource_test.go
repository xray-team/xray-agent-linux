package vmStat_test

import (
	"reflect"
	"testing"

	"github.com/xray-team/xray-agent-linux/collectors/vmStat"
	"github.com/xray-team/xray-agent-linux/logger"
)

func Test_vmStatDataSource_GetData(t *testing.T) {
	logger.Init()

	tests := []struct {
		caseDescription string
		filePath        string
		want            *vmStat.VMStat
		wantErr         bool
	}{
		{
			caseDescription: "no file",
			filePath:        "./testFiles/nofile",
			want:            nil,
			wantErr:         true,
		},
		{
			caseDescription: "kernel5.4.0",
			filePath:        "./testFiles/vmstat-kernel5.4.0",
			want: &vmStat.VMStat{
				PgPgIn:         6048146,
				PgPgOut:        57268065,
				PSwpIn:         417,
				PSwpOut:        5566,
				PgFault:        308134150,
				PgMajFault:     25845,
				PgFree:         462668898,
				PgActivate:     10304971,
				PgDeactivate:   757885,
				PgLazyFree:     3044,
				PgLazyFreed:    1936,
				PgRefill:       825451,
				NumaHit:        455499789,
				NumaMiss:       2,
				NumaForeign:    3,
				NumaInterleave: 67183,
				NumaLocal:      455499789,
				NumaOther:      4,
				OOMKill:        1,
			},
			wantErr: false,
		},
	}

	for _, testCase := range tests {
		tt := testCase

		t.Run(tt.caseDescription, func(t *testing.T) {
			vmStatDataSource := vmStat.NewDataSource(tt.filePath, vmStat.CollectorName)
			got, err := vmStatDataSource.GetData()

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
