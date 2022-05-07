package ps_test

import (
	"reflect"
	"testing"

	"github.com/xray-team/xray-agent-linux/collectors/ps"
	"github.com/xray-team/xray-agent-linux/logger"
)

func Test_psDataSource_GetData(t *testing.T) {
	logger.Init()

	tests := []struct {
		caseDescription string
		path            string
		want            *ps.PS
		wantErr         bool
	}{
		{
			caseDescription: "invalid",
			path:            "./testFiles/invalid",
			want:            nil,
			wantErr:         true,
		},
		{
			caseDescription: "proc1",
			path:            "./testFiles/proc1",
			want: &ps.PS{
				Count:            9,
				Limit:            32768,
				Threads:          22,
				ThreadsLimit:     126688,
				InStateRunning:   1,
				InStateIdle:      2,
				InStateSleeping:  2,
				InStateDiskSleep: 1,
				InStateStopped:   1,
				InStateZombie:    1,
				InStateDead:      1,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.caseDescription, func(t *testing.T) {
			psDataSource := ps.NewPSDataSource(tt.path, "")
			got, err := psDataSource.GetData()
			if (err != nil) != tt.wantErr {
				t.Errorf("ParsePS() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParsePS() got = %v, want %v", got, tt.want)
			}
		})
	}
}
