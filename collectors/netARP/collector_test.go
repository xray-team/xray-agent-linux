package netARP_test

import (
	"reflect"
	"testing"

	"github.com/xray-team/xray-agent-linux/collectors/netARP"
	"github.com/xray-team/xray-agent-linux/dto"
	"github.com/xray-team/xray-agent-linux/logger"
)

func TestNetARPCollector_Collect(t *testing.T) {
	logger.Init()

	tests := []struct {
		name     string
		filePath string
		want     *dto.NetArp
		wantErr  bool
	}{
		{
			name:     "no file",
			filePath: "./testFiles/nofile",
			wantErr:  true,
		},
		{
			name:     "with-incomplets",
			filePath: "./testFiles/netarp-with-incomplets",
			want: &dto.NetArp{
				Entries: map[string]uint{
					"Total":  7,
					"wlp2s0": 4,
					"enp1s0": 3,
				},
				IncompleteEntries: map[string]uint{
					"Total":  4,
					"wlp2s0": 2,
					"enp1s0": 2,
				},
			},
			wantErr: false,
		},
	}

	for _, testCase := range tests {
		tt := testCase

		t.Run(tt.name, func(t *testing.T) {
			c := netARP.NetARPCollector{DataSource: netARP.NewNetARPDataSource(tt.filePath, "")}
			got, err := c.GetNetArp()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetNetArp() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetNetArp() got = %v, want %v", got, tt.want)
			}
		})
	}
}
