package mdStat_test

import (
	"reflect"
	"testing"

	"github.com/xray-team/xray-agent-linux/collectors/mdStat"
	"github.com/xray-team/xray-agent-linux/logger"
)

func Test_mdStatDataSource_GetData(t *testing.T) {
	logger.Init()

	tests := []struct {
		name    string
		path    string
		want    *mdStat.MDStats
		wantErr bool
	}{
		{
			name:    "no files",
			path:    "./",
			want:    nil,
			wantErr: true,
		},
		{
			name: "RAID1-create-missing",
			path: "./testFiles/RAID1-create-missing",
			want: &mdStat.MDStats{
				Stats: map[string]mdStat.MDStat{
					"md0": {
						Level:               "raid1",
						NumDisks:            2,
						ArrayState:          "clean",
						ArraySizeKBytes:     0,
						ComponentSizeKBytes: 7334912,
						StatRaidWithRedundancy: &mdStat.StatRaidWithRedundancy{
							SyncAction:           "idle",
							NumDegraded:          1,
							MismatchCnt:          0,
							SyncCompletedSectors: 0,
							NumSectors:           0,
							SyncSpeed:            0,
						},
						DevStats: map[string]mdStat.DevStats{
							"sdb1": {
								Slot:   "0",
								State:  "in_sync",
								Errors: 0,
							},
						},
					},
					"md1": {
						Level:               "raid1",
						NumDisks:            2,
						ArrayState:          "clean",
						ArraySizeKBytes:     0,
						ComponentSizeKBytes: 13621248,
						StatRaidWithRedundancy: &mdStat.StatRaidWithRedundancy{
							SyncAction:           "idle",
							NumDegraded:          1,
							MismatchCnt:          0,
							SyncCompletedSectors: 0,
							NumSectors:           0,
							SyncSpeed:            0,
						},
						DevStats: map[string]mdStat.DevStats{
							"sdb2": {
								Slot:   "0",
								State:  "in_sync",
								Errors: 0,
							},
						},
					},
				},
			},
		},
		{
			name: "RAID1-resync-DELAYED",
			path: "./testFiles/RAID1-resync-DELAYED",
			want: &mdStat.MDStats{
				Stats: map[string]mdStat.MDStat{
					"md0": {
						Level:               "raid1",
						NumDisks:            2,
						ArrayState:          "clean",
						ArraySizeKBytes:     0,
						ComponentSizeKBytes: 7334912,
						StatRaidWithRedundancy: &mdStat.StatRaidWithRedundancy{
							SyncAction:           "recover",
							NumDegraded:          1,
							MismatchCnt:          0,
							SyncCompletedSectors: 2750976,
							NumSectors:           14669824,
							SyncSpeed:            200246,
						},
						DevStats: map[string]mdStat.DevStats{
							"sdb1": {
								Slot:   "0",
								State:  "in_sync",
								Errors: 0,
							},
							"sdc1": {
								Slot:   "1",
								State:  "spare",
								Errors: 0,
							},
						},
					},
					"md1": {
						Level:               "raid1",
						NumDisks:            2,
						ArrayState:          "clean",
						ArraySizeKBytes:     0,
						ComponentSizeKBytes: 13621248,
						StatRaidWithRedundancy: &mdStat.StatRaidWithRedundancy{
							SyncAction:           "recover",
							NumDegraded:          1,
							MismatchCnt:          0,
							SyncCompletedSectors: 0,
							NumSectors:           0,
							SyncSpeed:            0,
						},
						DevStats: map[string]mdStat.DevStats{
							"sdb2": {
								Slot:   "0",
								State:  "in_sync",
								Errors: 0,
							},
							"sdc2": {
								Slot:   "1",
								State:  "spare",
								Errors: 0,
							},
						},
					},
				},
			},
		},
		{
			name: "RAID1-ok",
			path: "./testFiles/RAID1-ok",
			want: &mdStat.MDStats{
				Stats: map[string]mdStat.MDStat{
					"md0": {
						Level:               "raid1",
						NumDisks:            2,
						ArrayState:          "clean",
						ArraySizeKBytes:     0,
						ComponentSizeKBytes: 7334912,
						StatRaidWithRedundancy: &mdStat.StatRaidWithRedundancy{
							SyncAction:           "idle",
							NumDegraded:          0,
							MismatchCnt:          0,
							SyncCompletedSectors: 0,
							NumSectors:           0,
							SyncSpeed:            0,
						},
						DevStats: map[string]mdStat.DevStats{
							"sdb1": {
								Slot:   "0",
								State:  "in_sync",
								Errors: 0,
							},
							"sdc1": {
								Slot:   "1",
								State:  "in_sync",
								Errors: 0,
							},
						},
					},
					"md1": {
						Level:               "raid1",
						NumDisks:            2,
						ArrayState:          "clean",
						ArraySizeKBytes:     0,
						ComponentSizeKBytes: 13621248,
						StatRaidWithRedundancy: &mdStat.StatRaidWithRedundancy{
							SyncAction:           "idle",
							NumDegraded:          0,
							MismatchCnt:          0,
							SyncCompletedSectors: 0,
							NumSectors:           0,
							SyncSpeed:            0,
						},
						DevStats: map[string]mdStat.DevStats{
							"sdb2": {
								Slot:   "0",
								State:  "in_sync",
								Errors: 0,
							},
							"sdc2": {
								Slot:   "1",
								State:  "in_sync",
								Errors: 0,
							},
						},
					},
				},
			},
		},
		{
			name: "RAID1-with-spare",
			path: "./testFiles/RAID1-with-spare",
			want: &mdStat.MDStats{
				Stats: map[string]mdStat.MDStat{
					"md0": {
						Level:               "raid1",
						NumDisks:            2,
						ArrayState:          "clean",
						ArraySizeKBytes:     0,
						ComponentSizeKBytes: 7334912,
						StatRaidWithRedundancy: &mdStat.StatRaidWithRedundancy{
							SyncAction:           "idle",
							NumDegraded:          0,
							MismatchCnt:          0,
							SyncCompletedSectors: 0,
							NumSectors:           0,
							SyncSpeed:            0,
						},
						DevStats: map[string]mdStat.DevStats{
							"sdb1": {
								Slot:   "0",
								State:  "in_sync",
								Errors: 0,
							},
							"sdc1": {
								Slot:   "1",
								State:  "in_sync",
								Errors: 0,
							},
							"sdd1": {
								Slot:   "none",
								State:  "spare",
								Errors: 0,
							},
						},
					},
					"md1": {
						Level:               "raid1",
						NumDisks:            2,
						ArrayState:          "clean",
						ArraySizeKBytes:     0,
						ComponentSizeKBytes: 13621248,
						StatRaidWithRedundancy: &mdStat.StatRaidWithRedundancy{
							SyncAction:           "idle",
							NumDegraded:          0,
							MismatchCnt:          0,
							SyncCompletedSectors: 0,
							NumSectors:           0,
							SyncSpeed:            0,
						},
						DevStats: map[string]mdStat.DevStats{
							"sdb2": {
								Slot:   "0",
								State:  "in_sync",
								Errors: 0,
							},
							"sdc2": {
								Slot:   "1",
								State:  "in_sync",
								Errors: 0,
							},
							"sdd2": {
								Slot:   "none",
								State:  "spare",
								Errors: 0,
							},
						},
					},
				},
			},
		},
		{
			name: "RAID5-and-RAID0",
			path: "./testFiles/RAID5-and-RAID0",
			want: &mdStat.MDStats{
				Stats: map[string]mdStat.MDStat{
					"md0": {
						Level:               "raid5",
						NumDisks:            3,
						ArrayState:          "clean",
						ArraySizeKBytes:     0,
						ComponentSizeKBytes: 7334912,
						StatRaidWithRedundancy: &mdStat.StatRaidWithRedundancy{
							SyncAction:           "idle",
							NumDegraded:          0,
							MismatchCnt:          0,
							SyncCompletedSectors: 0,
							NumSectors:           0,
							SyncSpeed:            0,
						},
						DevStats: map[string]mdStat.DevStats{
							"sdb1": {
								Slot:   "0",
								State:  "in_sync",
								Errors: 0,
							},
							"sdc1": {
								Slot:   "1",
								State:  "in_sync",
								Errors: 0,
							},
							"sdd1": {
								Slot:   "2",
								State:  "in_sync",
								Errors: 0,
							},
						},
					},
					"md1": {
						Level:                  "raid0",
						NumDisks:               2,
						ArrayState:             "clean",
						ArraySizeKBytes:        0,
						ComponentSizeKBytes:    0,
						StatRaidWithRedundancy: nil,
						DevStats: map[string]mdStat.DevStats{
							"sdb2": {
								Slot:   "0",
								State:  "in_sync",
								Errors: 0,
							},
							"sdc2": {
								Slot:   "1",
								State:  "in_sync",
								Errors: 0,
							},
						},
					},
				},
			},
		},
		{
			name: "RAID6",
			path: "./testFiles/RAID6",
			want: &mdStat.MDStats{
				Stats: map[string]mdStat.MDStat{
					"md0": {
						Level:               "raid6",
						NumDisks:            4,
						ArrayState:          "clean",
						ArraySizeKBytes:     0,
						ComponentSizeKBytes: 10475520,
						StatRaidWithRedundancy: &mdStat.StatRaidWithRedundancy{
							SyncAction:           "idle",
							NumDegraded:          0,
							MismatchCnt:          0,
							SyncCompletedSectors: 0,
							NumSectors:           0,
							SyncSpeed:            0,
						},
						DevStats: map[string]mdStat.DevStats{
							"sdb1": {
								Slot:   "0",
								State:  "in_sync",
								Errors: 0,
							},
							"sdc1": {
								Slot:   "1",
								State:  "in_sync",
								Errors: 0,
							},
							"sdd1": {
								Slot:   "2",
								State:  "in_sync",
								Errors: 0,
							},
							"sde1": {
								Slot:   "3",
								State:  "in_sync",
								Errors: 0,
							},
						},
					},
				},
			},
		},
		{
			name: "RAID10",
			path: "./testFiles/RAID10",
			want: &mdStat.MDStats{
				Stats: map[string]mdStat.MDStat{
					"md0": {
						Level:               "raid10",
						NumDisks:            4,
						ArrayState:          "clean",
						ArraySizeKBytes:     0,
						ComponentSizeKBytes: 10475520,
						StatRaidWithRedundancy: &mdStat.StatRaidWithRedundancy{
							SyncAction:           "idle",
							NumDegraded:          0,
							MismatchCnt:          0,
							SyncCompletedSectors: 0,
							NumSectors:           0,
							SyncSpeed:            0,
						},
						DevStats: map[string]mdStat.DevStats{
							"sdb1": {
								Slot:   "0",
								State:  "in_sync",
								Errors: 0,
							},
							"sdc1": {
								Slot:   "1",
								State:  "in_sync",
								Errors: 0,
							},
							"sdd1": {
								Slot:   "2",
								State:  "in_sync",
								Errors: 0,
							},
							"sde1": {
								Slot:   "3",
								State:  "in_sync",
								Errors: 0,
							},
						},
					},
				},
			},
		},
		{
			name: "RAID1-fake",
			path: "./testFiles/RAID1-fake",
			want: &mdStat.MDStats{
				Stats: map[string]mdStat.MDStat{
					"md0": {
						Level:               "raid1",
						NumDisks:            2,
						ArrayState:          "clean",
						ArraySizeKBytes:     10000000,
						ComponentSizeKBytes: 7334912,
						StatRaidWithRedundancy: &mdStat.StatRaidWithRedundancy{
							SyncAction:           "idle",
							NumDegraded:          0,
							MismatchCnt:          42,
							SyncCompletedSectors: 0,
							NumSectors:           0,
							SyncSpeed:            0,
						},
						DevStats: map[string]mdStat.DevStats{
							"sdb1": {
								Slot:   "0",
								State:  "in_sync",
								Errors: 73,
							},
							"sdc1": {
								Slot:   "1",
								State:  "in_sync",
								Errors: 0,
							},
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mdStatDataSource := mdStat.NewMDStatDataSource(tt.path, "")
			got, err := mdStatDataSource.GetData()
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseMDStats() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseMDStats() = %v, want %v", got, tt.want)
			}
		})
	}
}
