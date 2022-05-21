package conf_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/xray-team/xray-agent-linux/conf"
	"github.com/xray-team/xray-agent-linux/dto"
	"github.com/xray-team/xray-agent-linux/logger"
)

func TestReadConfigFile(t *testing.T) {
	logger.Init()

	tests := []struct {
		name     string
		filePath string
		want     *conf.Config
		wantErr  bool
	}{
		{
			name:     "noFile",
			filePath: "./testFiles/noFile",
			want:     nil,
			wantErr:  true,
		},
		{
			name:     "fullConfig",
			filePath: "./testFiles/fullConfig.json",
			want: &conf.Config{
				Agent: &conf.AgentConf{
					GetStatIntervalSec: 60,
					EnableSelfMetrics:  true,
					HostAttributes: []dto.MetricAttribute{
						{
							Name:  "Source",
							Value: "xray",
						},
					},
					LogLevel: "default",
					LogOut:   "syslog",
				},
				Collectors: map[string]json.RawMessage{
					"uptime":  []byte(`{"enabled": true}`),
					"loadAvg": []byte("{\n      \"enabled\": true\n    }"),
					"psStat":  []byte(`{"enabled": true, "collectPerPidStat": false, "processList": ["xray-agent"]}`),
					"netDev":  []byte("{\n      \"enabled\": true,\n      \"excludeLoopbacks\": true,\n      \"excludeWireless\": false,\n      \"excludeBridges\": false,\n      \"excludeVirtual\": false,\n      \"excludeByName\": [\n        \"tun0\",\n        \"tun1\"\n      ]\n    }"),
				},
				TSDB: &conf.TSDBConf{
					Graphite: &conf.GraphiteConf{
						Servers: []conf.GraphiteServerConf{
							{
								Mode:     dto.GraphiteModeTree,
								Address:  "192.168.0.10:2003",
								Protocol: "tcp",
								Timeout:  10,
							},
							{
								Mode:     dto.GraphiteModeTags,
								Address:  "192.168.0.20:2003",
								Protocol: "tcp",
								Timeout:  10,
							},
						},
					},
				},
			},
			wantErr: false,
		},
	}

	for _, testCase := range tests {
		tt := testCase

		t.Run(tt.name, func(t *testing.T) {
			got, err := conf.ReadConfigFile(tt.filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadConfigFile() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadConfigFile() got = %v, want %v", got, tt.want)
			}
		})
	}
}
