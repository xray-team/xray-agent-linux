package memoryInfo_test

import (
	"reflect"
	"testing"

	"github.com/xray-team/xray-agent-linux/collectors/memoryInfo"
	"github.com/xray-team/xray-agent-linux/dto"
	"github.com/xray-team/xray-agent-linux/logger"
)

func Test_memoryDataSource_GetData(t *testing.T) {
	logger.Init()

	tests := []struct {
		caseDescription string
		filePath        string
		want            *dto.MemoryInfo
		wantErr         bool
	}{
		{
			caseDescription: "no file",
			filePath:        "./testFiles/nofile",
			want:            nil,
			wantErr:         true,
		},
		{
			caseDescription: "mint19.2",
			filePath:        "./testFiles/meminfo-Mint19.2-linux5.0.0-32-generic",
			want: &dto.MemoryInfo{
				MemTotal:     16316304,
				MemFree:      6823340,
				MemAvailable: 11863620,
				Buffers:      475280,
				Cached:       5032104,
				SwapTotal:    16671740,
				SwapFree:     16671730,
			},
			wantErr: false,
		},
		{
			caseDescription: "debian7.11",
			filePath:        "./testFiles/meminfo-debian7.11-linux2.6.32-openvz",
			want: &dto.MemoryInfo{
				MemTotal:     264115680,
				MemFree:      22077932,
				MemAvailable: 414156072,
				Buffers:      13972328,
				Cached:       158145996,
				SwapTotal:    16762876,
				SwapFree:     16647008,
			},
			wantErr: false,
		},
	}

	for _, testCase := range tests {
		tt := testCase

		t.Run(tt.caseDescription, func(t *testing.T) {
			memInfoDataSource := memoryInfo.NewMemoryDataSource(tt.filePath, "")
			got, err := memInfoDataSource.GetData()

			if (err != nil) != tt.wantErr {
				t.Errorf("ParseMemInfo() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseMemInfo() got = %v, want %v", got, tt.want)
			}
		})
	}
}
