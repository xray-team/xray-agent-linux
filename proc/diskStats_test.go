package proc_test

import (
	"reflect"
	"testing"

	"xray-agent-linux/dto"
	"xray-agent-linux/logger"
	"xray-agent-linux/proc"
)

func Test_diskStatsDataSource_GetData(t *testing.T) {
	logger.Init("")

	tests := []struct {
		caseDescription string
		filePath        string
		want            []dto.DiskStat
		wantErr         bool
	}{
		{
			caseDescription: "no file",
			filePath:        "./testfiles/diskstat/nofile",
			want:            nil,
			wantErr:         true,
		},
		{
			caseDescription: "kernel 4.15",
			filePath:        "./testfiles/diskstat/diskstats-kernel4.15.0-66-generic",
			want: []dto.DiskStat{
				{
					Major:                         7,
					Miner:                         0,
					Dev:                           "loop0",
					ReadsCompletedSuccessfully:    5,
					ReadsMerged:                   0,
					SectorsRead:                   16,
					TimeSpentReading:              0,
					WritesCompleted:               0,
					WritesMerged:                  0,
					SectorsWritten:                0,
					TimeSpentWriting:              0,
					IOsCurrentlyInProgress:        0,
					TimeSpentDoingIOs:             0,
					WeightedTimeSpentDoingIOs:     0,
					DiscardsCompletedSuccessfully: 0,
					DiscardsMerged:                0,
					SectorsDiscarded:              0,
					TimeSpentDiscarding:           0,
				},
				{
					Major:                         7,
					Miner:                         1,
					Dev:                           "loop1",
					ReadsCompletedSuccessfully:    0,
					ReadsMerged:                   0,
					SectorsRead:                   0,
					TimeSpentReading:              0,
					WritesCompleted:               0,
					WritesMerged:                  0,
					SectorsWritten:                0,
					TimeSpentWriting:              0,
					IOsCurrentlyInProgress:        0,
					TimeSpentDoingIOs:             0,
					WeightedTimeSpentDoingIOs:     0,
					DiscardsCompletedSuccessfully: 0,
					DiscardsMerged:                0,
					SectorsDiscarded:              0,
					TimeSpentDiscarding:           0,
				},
				{
					Major:                         8,
					Miner:                         0,
					Dev:                           "sda",
					ReadsCompletedSuccessfully:    31592,
					ReadsMerged:                   0,
					SectorsRead:                   4307174,
					TimeSpentReading:              293288,
					WritesCompleted:               770177,
					WritesMerged:                  266072,
					SectorsWritten:                55271602,
					TimeSpentWriting:              577676,
					IOsCurrentlyInProgress:        1,
					TimeSpentDoingIOs:             133132,
					WeightedTimeSpentDoingIOs:     355240,
					DiscardsCompletedSuccessfully: 0,
					DiscardsMerged:                0,
					SectorsDiscarded:              0,
					TimeSpentDiscarding:           0,
				},
				{
					Major:                         8,
					Miner:                         1,
					Dev:                           "sda1",
					ReadsCompletedSuccessfully:    31563,
					ReadsMerged:                   2,
					SectorsRead:                   4305054,
					TimeSpentReading:              293272,
					WritesCompleted:               768141,
					WritesMerged:                  266072,
					SectorsWritten:                55271602,
					TimeSpentWriting:              575256,
					IOsCurrentlyInProgress:        3,
					TimeSpentDoingIOs:             128912,
					WeightedTimeSpentDoingIOs:     349104,
					DiscardsCompletedSuccessfully: 0,
					DiscardsMerged:                0,
					SectorsDiscarded:              0,
					TimeSpentDiscarding:           0,
				},
			},
			wantErr: false,
		},
		{
			caseDescription: "kernel 5.0",
			filePath:        "./testfiles/diskstat/diskstats-kernel5.0.0-32-generic",
			want: []dto.DiskStat{
				{
					Major:                         8,
					Miner:                         0,
					Dev:                           "sda",
					ReadsCompletedSuccessfully:    338,
					ReadsMerged:                   13,
					SectorsRead:                   19431,
					TimeSpentReading:              79,
					WritesCompleted:               1,
					WritesMerged:                  0,
					SectorsWritten:                8,
					TimeSpentWriting:              0,
					IOsCurrentlyInProgress:        0,
					TimeSpentDoingIOs:             100,
					WeightedTimeSpentDoingIOs:     0,
					DiscardsCompletedSuccessfully: 0,
					DiscardsMerged:                0,
					SectorsDiscarded:              0,
					TimeSpentDiscarding:           0,
				},
				{
					Major:                         8,
					Miner:                         1,
					Dev:                           "sda1",
					ReadsCompletedSuccessfully:    73,
					ReadsMerged:                   0,
					SectorsRead:                   6296,
					TimeSpentReading:              19,
					WritesCompleted:               0,
					WritesMerged:                  0,
					SectorsWritten:                0,
					TimeSpentWriting:              0,
					IOsCurrentlyInProgress:        0,
					TimeSpentDoingIOs:             52,
					WeightedTimeSpentDoingIOs:     0,
					DiscardsCompletedSuccessfully: 0,
					DiscardsMerged:                0,
					SectorsDiscarded:              0,
					TimeSpentDiscarding:           0,
				},
				{
					Major:                         8,
					Miner:                         16,
					Dev:                           "sdb",
					ReadsCompletedSuccessfully:    2546747,
					ReadsMerged:                   12350,
					SectorsRead:                   352373968,
					TimeSpentReading:              1610594,
					WritesCompleted:               50740,
					WritesMerged:                  53841,
					SectorsWritten:                5454098,
					TimeSpentWriting:              83299,
					IOsCurrentlyInProgress:        0,
					TimeSpentDoingIOs:             828176,
					WeightedTimeSpentDoingIOs:     115700,
					DiscardsCompletedSuccessfully: 1,
					DiscardsMerged:                2,
					SectorsDiscarded:              3,
					TimeSpentDiscarding:           4,
				},
				{
					Major:                         8,
					Miner:                         18,
					Dev:                           "sdb1",
					ReadsCompletedSuccessfully:    2546302,
					ReadsMerged:                   12347,
					SectorsRead:                   352347402,
					TimeSpentReading:              1609751,
					WritesCompleted:               49210,
					WritesMerged:                  53841,
					SectorsWritten:                5454096,
					TimeSpentWriting:              82882,
					IOsCurrentlyInProgress:        0,
					TimeSpentDoingIOs:             827988,
					WeightedTimeSpentDoingIOs:     115236,
					DiscardsCompletedSuccessfully: 0,
					DiscardsMerged:                0,
					SectorsDiscarded:              0,
					TimeSpentDiscarding:           0,
				},
			},
			wantErr: false,
		},
	}

	for _, testCase := range tests {
		tt := testCase

		t.Run(tt.caseDescription, func(t *testing.T) {
			blockDevDataSource := proc.NewBlockDevDataSource(tt.filePath, "")
			got, err := blockDevDataSource.GetData()

			if (err != nil) != tt.wantErr {
				t.Errorf("ParseDiskStats() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseDiskStats() got = %v, want %v", got, tt.want)
			}
		})
	}
}
