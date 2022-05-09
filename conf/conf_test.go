package conf_test

import (
	"reflect"
	"testing"

	"github.com/xray-team/xray-agent-linux/collectors/cpuInfo"
	"github.com/xray-team/xray-agent-linux/collectors/diskSpace"
	"github.com/xray-team/xray-agent-linux/collectors/diskStat"
	"github.com/xray-team/xray-agent-linux/collectors/entropy"
	"github.com/xray-team/xray-agent-linux/collectors/loadAvg"
	"github.com/xray-team/xray-agent-linux/collectors/mdStat"
	"github.com/xray-team/xray-agent-linux/collectors/memoryInfo"
	"github.com/xray-team/xray-agent-linux/collectors/netARP"
	"github.com/xray-team/xray-agent-linux/collectors/netDev"
	"github.com/xray-team/xray-agent-linux/collectors/netDevStatus"
	"github.com/xray-team/xray-agent-linux/collectors/netSNMP"
	"github.com/xray-team/xray-agent-linux/collectors/netSNMP6"
	"github.com/xray-team/xray-agent-linux/collectors/netStat"
	"github.com/xray-team/xray-agent-linux/collectors/ps"
	"github.com/xray-team/xray-agent-linux/collectors/psStat"
	"github.com/xray-team/xray-agent-linux/collectors/stat"
	"github.com/xray-team/xray-agent-linux/collectors/uptime"
	"github.com/xray-team/xray-agent-linux/collectors/wireless"
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
					EnableSelfMetrics: true,
					Uptime:            &uptime.Config{Enabled: true},
					LoadAvg:           &loadAvg.Config{Enabled: true},
					PS:                &ps.Config{Enabled: true},
					PSStat:            &psStat.Config{Enabled: true, ProcessList: []string{"xray-agent"}},
					Stat:              &stat.Config{Enabled: true},
					CPUInfo:           &cpuInfo.Config{Enabled: true},
					MemoryInfo:        &memoryInfo.Config{Enabled: true},
					NetARP:            &netARP.Config{Enabled: true},
					NetStat:           &netStat.Config{Enabled: true},
					NetSNMP:           &netSNMP.Config{Enabled: true},
					NetSNMP6:          &netSNMP6.Config{Enabled: true},
					Entropy:           &entropy.Config{Enabled: true},
					NetDev: &netDev.Config{
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
					NetDevStatus: &netDevStatus.Config{
						Enabled:         true,
						ExcludeWireless: true,
						ExcludeByName:   nil,
					},
					Wireless: &wireless.Config{
						Enabled:            true,
						ExcludeByName:      nil,
						ExcludeByOperState: nil,
					},
					DiskStat: &diskStat.Config{
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
					DiskSpace: &diskSpace.Config{
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
					MDStat: &mdStat.Config{Enabled: true},
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
