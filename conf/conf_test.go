package conf_test

import (
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
					HostAttributes: []dto.MetricAttribute{
						{
							Name:  "Source",
							Value: "xray",
						},
					},
					LogLevel: "default",
					LogOut:   "syslog",
				},
				Collectors: &conf.CollectorsConf{
					RootPath:          "/",
					EnableSelfMetrics: true,
					Uptime:            &conf.UptimeConf{Enabled: true},
					LoadAvg:           &conf.LoadAvgConf{Enabled: true},
					PS:                &conf.PSConf{Enabled: true},
					PSStat:            &conf.PSStatConf{Enabled: true, ProcessList: []string{"xray-agent"}},
					Stat:              &conf.StatConf{Enabled: true},
					CPUInfo:           &conf.CPUInfoConf{Enabled: true},
					MemoryInfo:        &conf.MemoryInfoConf{Enabled: true},
					NetARP:            &conf.NetARPConf{Enabled: true},
					NetStat:           &conf.NetStatConf{Enabled: true},
					NetSNMP:           &conf.NetSNMPConf{Enabled: true},
					NetSNMP6:          &conf.NetSNMP6Conf{Enabled: true},
					Entropy:           &conf.EntropyConf{Enabled: true},
					NetDev: &conf.NetDevConf{
						Enabled:          true,
						ExcludeLoopbacks: true,
						ExcludeWireless:  false,
						ExcludeBridges:   false,
						ExcludeVirtual:   false,
						ExcludeByName: []string{
							"tun0",
							"tun1",
						},
					},
					NetDevStatus: &conf.NetDevStatusConf{
						Enabled:         true,
						ExcludeWireless: true,
						ExcludeByName:   nil,
					},
					Wireless: &conf.WirelessConf{
						Enabled:            true,
						ExcludeByName:      nil,
						ExcludeByOperState: nil,
					},
					DiskStat: &conf.DiskStatConf{
						Enabled: true,
						MonitoredDiskTypes: []int64{
							dto.BlockDevMajorTypeSCSI, // SCSI disk devices  (sda*, sdb*, sdc*, ...,)
							dto.BlockDevMajorTypeMD,   // Metadisk (RAID) devices (md0, md1, ...)
						},
						ExcludeByName: []string{
							"sde",
							"sde1",
						},
					},
					DiskSpace: &conf.DiskSpaceConf{
						Enabled: true,
						MonitoredFileSystemTypes: []string{
							"ext4",
							"ext3",
							"ext2",
							"xfs",
							"jfs",
							"btrfs",
						},
					},
					MDStat: &conf.MDStatConf{Enabled: true},
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
