package proc_test

import (
	"reflect"
	"testing"

	"github.com/xray-team/xray-agent-linux/dto"
	"github.com/xray-team/xray-agent-linux/logger"
	"github.com/xray-team/xray-agent-linux/proc"
)

func Test_statDataSource_GetData(t *testing.T) {
	logger.Init("")

	tests := []struct {
		caseDescription string
		filePath        string
		want            *dto.Stat
		wantErr         bool
	}{
		{
			caseDescription: "no file",
			filePath:        "./testfiles/stat/nofile",
			want:            nil,
			wantErr:         true,
		},
		{
			caseDescription: "kernel5.0",
			filePath:        "./testfiles/stat/stat-kernel5.0.0-32-generic",
			want: &dto.Stat{
				CPU: dto.CPUStats{
					User:      1029062,
					Nice:      1454,
					System:    269719,
					Idle:      10005830,
					IOwait:    31528,
					IRQ:       5,
					SoftIRQ:   64688,
					Steal:     23,
					Guest:     25,
					GuestNice: 27,
				},
				PerCPU: map[string]dto.CPUStats{
					"0": {
						User:      257829,
						Nice:      162,
						System:    66105,
						Idle:      2517673,
						IOwait:    10220,
						IRQ:       2,
						SoftIRQ:   27554,
						Steal:     10,
						Guest:     11,
						GuestNice: 12,
					},
					"1": {
						User:      257026,
						Nice:      756,
						System:    68168,
						Idle:      2458616,
						IOwait:    8186,
						IRQ:       3,
						SoftIRQ:   17546,
						Steal:     13,
						Guest:     14,
						GuestNice: 15,
					},
					"2": {
						User:      254606,
						Nice:      173,
						System:    69896,
						Idle:      2515573,
						IOwait:    7074,
						IRQ:       0,
						SoftIRQ:   11811,
						Steal:     0,
						Guest:     0,
						GuestNice: 0,
					},
					"3": {
						User:      259600,
						Nice:      361,
						System:    65549,
						Idle:      2513967,
						IOwait:    6046,
						IRQ:       0,
						SoftIRQ:   7776,
						Steal:     0,
						Guest:     0,
						GuestNice: 0,
					},
				},
				Ctxt:             80500006,
				Intr:             45387578,
				Btime:            1574064476,
				Processes:        31992,
				ProcessesRunning: 3,
				ProcessesBlocked: 1,
				SoftIRQ: dto.SoftIRQStat{
					Total:   46493915,
					Hi:      4993583,
					Timer:   14688474,
					NetTx:   5779,
					NetRx:   5571,
					Block:   1279116,
					IRQPoll: 100,
					Tasklet: 21102,
					Sched:   15431535,
					HRTimer: 1052,
					RCU:     10067703,
				},
			},
			wantErr: false,
		},
	}

	for _, testCase := range tests {
		tt := testCase

		t.Run(tt.caseDescription, func(t *testing.T) {
			statDataSource := proc.NewStatDataSource(tt.filePath, "")
			got, err := statDataSource.GetData()
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseStat() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseStat() got = %v, want %v", got, tt.want)
			}
		})
	}
}
