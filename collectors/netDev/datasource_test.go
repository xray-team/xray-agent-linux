package netDev_test

import (
	"reflect"
	"testing"

	"github.com/xray-team/xray-agent-linux/collectors/netDev"
	"github.com/xray-team/xray-agent-linux/logger"
)

func Test_netDevDataSource_GetData(t *testing.T) {
	logger.Init()

	tests := []struct {
		caseDescription string
		filePath        string
		want            map[string]netDev.NetDevStatistics
		wantErr         bool
	}{
		{
			caseDescription: "no file",
			filePath:        "./testFiles/nofile",
			want:            nil,
			wantErr:         true,
		},
		{
			caseDescription: "lo",
			filePath:        "./testFiles/lo",
			want: map[string]netDev.NetDevStatistics{
				"lo": {
					RxBytes:       1,
					RxPackets:     2,
					RxErrs:        3,
					RxDrop:        4,
					RxFifoErrs:    5,
					RxFrameErrs:   6,
					RxCompressed:  7,
					Multicast:     8,
					TxBytes:       9,
					TxPackets:     10,
					TxErrs:        11,
					TxDrop:        12,
					TxFifoErrs:    13,
					Collisions:    14,
					TxCarrierErrs: 15,
					TxCompressed:  16,
				},
			},
			wantErr: false,
		},
		{
			caseDescription: "lo",
			filePath:        "./testFiles/lo-eth0",
			want: map[string]netDev.NetDevStatistics{
				"lo": {
					RxBytes:       13641,
					RxPackets:     166,
					RxErrs:        0,
					RxDrop:        0,
					RxFifoErrs:    0,
					RxFrameErrs:   0,
					RxCompressed:  0,
					Multicast:     0,
					TxBytes:       13641,
					TxPackets:     166,
					TxErrs:        0,
					TxDrop:        0,
					TxFifoErrs:    0,
					Collisions:    0,
					TxCarrierErrs: 0,
					TxCompressed:  0,
				},
				"eth0": {
					RxBytes:       6725295,
					RxPackets:     18112,
					RxErrs:        0,
					RxDrop:        0,
					RxFifoErrs:    0,
					RxFrameErrs:   0,
					RxCompressed:  0,
					Multicast:     0,
					TxBytes:       2827329,
					TxPackets:     15859,
					TxErrs:        0,
					TxDrop:        0,
					TxFifoErrs:    0,
					Collisions:    0,
					TxCarrierErrs: 0,
					TxCompressed:  0,
				},
			},
			wantErr: false,
		},
	}

	for _, testCase := range tests {
		tt := testCase

		t.Run(tt.caseDescription, func(t *testing.T) {
			netDevDataSource := netDev.NewNetDevDataSource(tt.filePath, "")
			got, err := netDevDataSource.GetData()

			if (err != nil) != tt.wantErr {
				t.Errorf("ParseNetDev() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseNetDev() got = %v, want %v", got, tt.want)
			}
		})
	}
}
