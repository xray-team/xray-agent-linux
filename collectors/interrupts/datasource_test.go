package interrupts_test

import (
	"reflect"
	"testing"

	"github.com/xray-team/xray-agent-linux/collectors/interrupts"
	"github.com/xray-team/xray-agent-linux/logger"
)

func Test_interruptsDataSource_GetData(t *testing.T) {
	logger.Init()

	tests := []struct {
		caseDescription string
		filePath        string
		want            *interrupts.Interrupts
		wantErr         bool
	}{
		{
			caseDescription: "no file",
			filePath:        "./testFiles/nofile",
			want:            nil,
			wantErr:         true,
		},
		{
			caseDescription: "kernel5.15",
			filePath:        "./testFiles/interrupts-kernel5.15.0-76-generic",
			want: &interrupts.Interrupts{
				Total: interrupts.InterruptsStat{
					Total: 462516955,
				},
				PerCPU: map[int]interrupts.InterruptsStat{
					0: {Total: 49719585},
					1: {Total: 29309319},
					2: {Total: 36764937},
					3: {Total: 346723114},
				},
			},
		},
	}

	for _, testCase := range tests {
		tt := testCase

		t.Run(tt.caseDescription, func(t *testing.T) {
			interruptsDataSource := interrupts.NewDataSource(tt.filePath, "")
			got, err := interruptsDataSource.GetData()

			if (err != nil) != tt.wantErr {
				t.Errorf("ParseInterrupts() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseInterrupts() got = %v, want %v", got, tt.want)
			}
		})
	}
}
