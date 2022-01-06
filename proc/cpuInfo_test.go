package proc_test

import (
	"reflect"
	"testing"

	"github.com/xray-team/xray-agent-linux/dto"
	"github.com/xray-team/xray-agent-linux/logger"
	"github.com/xray-team/xray-agent-linux/proc"
)

func Test_cpuInfoDataSource_GetData(t *testing.T) {
	logger.Init("")

	tests := []struct {
		caseDescription string
		filePath        string
		want            *dto.CPUInfo
		wantErr         bool
	}{
		{
			caseDescription: "no file",
			filePath:        "./testfiles/cpuinfo/nofile",
			want:            nil,
			wantErr:         true,
		},
		{
			caseDescription: "kernel5.0",
			filePath:        "./testfiles/cpuinfo/cpuinfo-kernel5.0.0-32-generic",
			want: &dto.CPUInfo{
				CPU: []dto.CPU{
					{
						ProcessorNumber: 0,
						VendorID:        "GenuineIntel",
						ModelName:       "Intel(R) Core(TM) i5-7200U CPU @ 2.50GHz",
						MHz:             2095.591,
						CacheSize:       "3072 KB",
						PhysicalID:      0,
						CoreID:          0,
						CPUCores:        2,
						Bugs:            "cpu_meltdown spectre_v1 spectre_v2 spec_store_bypass l1tf mds swapgs",
					},
					{
						ProcessorNumber: 1,
						VendorID:        "GenuineIntel",
						ModelName:       "Intel(R) Core(TM) i5-7200U CPU @ 2.50GHz",
						MHz:             1839.038,
						CacheSize:       "3072 KB",
						PhysicalID:      0,
						CoreID:          1,
						CPUCores:        2,
						Bugs:            "cpu_meltdown spectre_v1 spectre_v2 spec_store_bypass l1tf mds swapgs",
					},
					{
						ProcessorNumber: 2,
						VendorID:        "GenuineIntel",
						ModelName:       "Intel(R) Core(TM) i5-7200U CPU @ 2.50GHz",
						MHz:             1870.994,
						CacheSize:       "3072 KB",
						PhysicalID:      0,
						CoreID:          0,
						CPUCores:        2,
						Bugs:            "cpu_meltdown spectre_v1 spectre_v2 spec_store_bypass l1tf mds swapgs",
					},
					{
						ProcessorNumber: 3,
						VendorID:        "GenuineIntel",
						ModelName:       "Intel(R) Core(TM) i5-7200U CPU @ 2.50GHz",
						MHz:             1600.185,
						CacheSize:       "3072 KB",
						PhysicalID:      0,
						CoreID:          1,
						CPUCores:        2,
						Bugs:            "cpu_meltdown spectre_v1 spectre_v2 spec_store_bypass l1tf mds swapgs",
					},
				},
			},
			wantErr: false,
		},
		{
			caseDescription: "2 cpu",
			filePath:        "./testfiles/cpuinfo/cpuinfo-2cpu-kernel5.4.0-48-generic",
			want: &dto.CPUInfo{
				CPU: []dto.CPU{
					{
						ProcessorNumber: 0,
						VendorID:        "GenuineIntel",
						ModelName:       "Intel(R) Xeon(R) CPU E5-2690 v2 @ 3.00GHz",
						MHz:             2999.936,
						CacheSize:       "25600 KB",
						PhysicalID:      0,
						CoreID:          0,
						CPUCores:        10,
						Bugs:            "cpu_meltdown spectre_v1 spectre_v2 spec_store_bypass l1tf mds swapgs itlb_multihit",
					},
					{
						ProcessorNumber: 1,
						VendorID:        "GenuineIntel",
						ModelName:       "Intel(R) Xeon(R) CPU E5-2690 v2 @ 3.00GHz",
						MHz:             3000.144,
						CacheSize:       "25600 KB",
						PhysicalID:      0,
						CoreID:          1,
						CPUCores:        10,
						Bugs:            "cpu_meltdown spectre_v1 spectre_v2 spec_store_bypass l1tf mds swapgs itlb_multihit",
					},
					{
						ProcessorNumber: 2,
						VendorID:        "GenuineIntel",
						ModelName:       "Intel(R) Xeon(R) CPU E5-2690 v2 @ 3.00GHz",
						MHz:             3000.125,
						CacheSize:       "25600 KB",
						PhysicalID:      0,
						CoreID:          2,
						CPUCores:        10,
						Bugs:            "cpu_meltdown spectre_v1 spectre_v2 spec_store_bypass l1tf mds swapgs itlb_multihit",
					},
					{
						ProcessorNumber: 3,
						VendorID:        "GenuineIntel",
						ModelName:       "Intel(R) Xeon(R) CPU E5-2690 v2 @ 3.00GHz",
						MHz:             2999.936,
						CacheSize:       "25600 KB",
						PhysicalID:      0,
						CoreID:          3,
						CPUCores:        10,
						Bugs:            "cpu_meltdown spectre_v1 spectre_v2 spec_store_bypass l1tf mds swapgs itlb_multihit",
					},
					{
						ProcessorNumber: 4,
						VendorID:        "GenuineIntel",
						ModelName:       "Intel(R) Xeon(R) CPU E5-2690 v2 @ 3.00GHz",
						MHz:             3000.101,
						CacheSize:       "25600 KB",
						PhysicalID:      0,
						CoreID:          4,
						CPUCores:        10,
						Bugs:            "cpu_meltdown spectre_v1 spectre_v2 spec_store_bypass l1tf mds swapgs itlb_multihit",
					},
					{
						ProcessorNumber: 5,
						VendorID:        "GenuineIntel",
						ModelName:       "Intel(R) Xeon(R) CPU E5-2690 v2 @ 3.00GHz",
						MHz:             2999.938,
						CacheSize:       "25600 KB",
						PhysicalID:      0,
						CoreID:          8,
						CPUCores:        10,
						Bugs:            "cpu_meltdown spectre_v1 spectre_v2 spec_store_bypass l1tf mds swapgs itlb_multihit",
					},
					{
						ProcessorNumber: 6,
						VendorID:        "GenuineIntel",
						ModelName:       "Intel(R) Xeon(R) CPU E5-2690 v2 @ 3.00GHz",
						MHz:             3000.125,
						CacheSize:       "25600 KB",
						PhysicalID:      0,
						CoreID:          9,
						CPUCores:        10,
						Bugs:            "cpu_meltdown spectre_v1 spectre_v2 spec_store_bypass l1tf mds swapgs itlb_multihit",
					},
					{
						ProcessorNumber: 7,
						VendorID:        "GenuineIntel",
						ModelName:       "Intel(R) Xeon(R) CPU E5-2690 v2 @ 3.00GHz",
						MHz:             2999.937,
						CacheSize:       "25600 KB",
						PhysicalID:      0,
						CoreID:          10,
						CPUCores:        10,
						Bugs:            "cpu_meltdown spectre_v1 spectre_v2 spec_store_bypass l1tf mds swapgs itlb_multihit",
					},
					{
						ProcessorNumber: 8,
						VendorID:        "GenuineIntel",
						ModelName:       "Intel(R) Xeon(R) CPU E5-2690 v2 @ 3.00GHz",
						MHz:             3001.773,
						CacheSize:       "25600 KB",
						PhysicalID:      0,
						CoreID:          11,
						CPUCores:        10,
						Bugs:            "cpu_meltdown spectre_v1 spectre_v2 spec_store_bypass l1tf mds swapgs itlb_multihit",
					},
					{
						ProcessorNumber: 9,
						VendorID:        "GenuineIntel",
						ModelName:       "Intel(R) Xeon(R) CPU E5-2690 v2 @ 3.00GHz",
						MHz:             3001.223,
						CacheSize:       "25600 KB",
						PhysicalID:      0,
						CoreID:          12,
						CPUCores:        10,
						Bugs:            "cpu_meltdown spectre_v1 spectre_v2 spec_store_bypass l1tf mds swapgs itlb_multihit",
					},
					{
						ProcessorNumber: 10,
						VendorID:        "GenuineIntel",
						ModelName:       "Intel(R) Xeon(R) CPU E5-2690 v2 @ 3.00GHz",
						MHz:             3299.945,
						CacheSize:       "25600 KB",
						PhysicalID:      1,
						CoreID:          0,
						CPUCores:        10,
						Bugs:            "cpu_meltdown spectre_v1 spectre_v2 spec_store_bypass l1tf mds swapgs itlb_multihit",
					},
					{
						ProcessorNumber: 11,
						VendorID:        "GenuineIntel",
						ModelName:       "Intel(R) Xeon(R) CPU E5-2690 v2 @ 3.00GHz",
						MHz:             3299.933,
						CacheSize:       "25600 KB",
						PhysicalID:      1,
						CoreID:          1,
						CPUCores:        10,
						Bugs:            "cpu_meltdown spectre_v1 spectre_v2 spec_store_bypass l1tf mds swapgs itlb_multihit",
					},
					{
						ProcessorNumber: 12,
						VendorID:        "GenuineIntel",
						ModelName:       "Intel(R) Xeon(R) CPU E5-2690 v2 @ 3.00GHz",
						MHz:             3299.931,
						CacheSize:       "25600 KB",
						PhysicalID:      1,
						CoreID:          2,
						CPUCores:        10,
						Bugs:            "cpu_meltdown spectre_v1 spectre_v2 spec_store_bypass l1tf mds swapgs itlb_multihit",
					},
					{
						ProcessorNumber: 13,
						VendorID:        "GenuineIntel",
						ModelName:       "Intel(R) Xeon(R) CPU E5-2690 v2 @ 3.00GHz",
						MHz:             3299.790,
						CacheSize:       "25600 KB",
						PhysicalID:      1,
						CoreID:          3,
						CPUCores:        10,
						Bugs:            "cpu_meltdown spectre_v1 spectre_v2 spec_store_bypass l1tf mds swapgs itlb_multihit",
					},
					{
						ProcessorNumber: 14,
						VendorID:        "GenuineIntel",
						ModelName:       "Intel(R) Xeon(R) CPU E5-2690 v2 @ 3.00GHz",
						MHz:             3299.934,
						CacheSize:       "25600 KB",
						PhysicalID:      1,
						CoreID:          4,
						CPUCores:        10,
						Bugs:            "cpu_meltdown spectre_v1 spectre_v2 spec_store_bypass l1tf mds swapgs itlb_multihit",
					},
					{
						ProcessorNumber: 15,
						VendorID:        "GenuineIntel",
						ModelName:       "Intel(R) Xeon(R) CPU E5-2690 v2 @ 3.00GHz",
						MHz:             3300.145,
						CacheSize:       "25600 KB",
						PhysicalID:      1,
						CoreID:          8,
						CPUCores:        10,
						Bugs:            "cpu_meltdown spectre_v1 spectre_v2 spec_store_bypass l1tf mds swapgs itlb_multihit",
					},
					{
						ProcessorNumber: 16,
						VendorID:        "GenuineIntel",
						ModelName:       "Intel(R) Xeon(R) CPU E5-2690 v2 @ 3.00GHz",
						MHz:             3298.458,
						CacheSize:       "25600 KB",
						PhysicalID:      1,
						CoreID:          9,
						CPUCores:        10,
						Bugs:            "cpu_meltdown spectre_v1 spectre_v2 spec_store_bypass l1tf mds swapgs itlb_multihit",
					},
					{
						ProcessorNumber: 17,
						VendorID:        "GenuineIntel",
						ModelName:       "Intel(R) Xeon(R) CPU E5-2690 v2 @ 3.00GHz",
						MHz:             3300.352,
						CacheSize:       "25600 KB",
						PhysicalID:      1,
						CoreID:          10,
						CPUCores:        10,
						Bugs:            "cpu_meltdown spectre_v1 spectre_v2 spec_store_bypass l1tf mds swapgs itlb_multihit",
					},
					{
						ProcessorNumber: 18,
						VendorID:        "GenuineIntel",
						ModelName:       "Intel(R) Xeon(R) CPU E5-2690 v2 @ 3.00GHz",
						MHz:             3299.930,
						CacheSize:       "25600 KB",
						PhysicalID:      1,
						CoreID:          11,
						CPUCores:        10,
						Bugs:            "cpu_meltdown spectre_v1 spectre_v2 spec_store_bypass l1tf mds swapgs itlb_multihit",
					},
					{
						ProcessorNumber: 19,
						VendorID:        "GenuineIntel",
						ModelName:       "Intel(R) Xeon(R) CPU E5-2690 v2 @ 3.00GHz",
						MHz:             3300.542,
						CacheSize:       "25600 KB",
						PhysicalID:      1,
						CoreID:          12,
						CPUCores:        10,
						Bugs:            "cpu_meltdown spectre_v1 spectre_v2 spec_store_bypass l1tf mds swapgs itlb_multihit",
					},
					{
						ProcessorNumber: 20,
						VendorID:        "GenuineIntel",
						ModelName:       "Intel(R) Xeon(R) CPU E5-2690 v2 @ 3.00GHz",
						MHz:             3003.762,
						CacheSize:       "25600 KB",
						PhysicalID:      0,
						CoreID:          0,
						CPUCores:        10,
						Bugs:            "cpu_meltdown spectre_v1 spectre_v2 spec_store_bypass l1tf mds swapgs itlb_multihit",
					},
					{
						ProcessorNumber: 21,
						VendorID:        "GenuineIntel",
						ModelName:       "Intel(R) Xeon(R) CPU E5-2690 v2 @ 3.00GHz",
						MHz:             3000.395,
						CacheSize:       "25600 KB",
						PhysicalID:      0,
						CoreID:          1,
						CPUCores:        10,
						Bugs:            "cpu_meltdown spectre_v1 spectre_v2 spec_store_bypass l1tf mds swapgs itlb_multihit",
					},
					{
						ProcessorNumber: 22,
						VendorID:        "GenuineIntel",
						ModelName:       "Intel(R) Xeon(R) CPU E5-2690 v2 @ 3.00GHz",
						MHz:             2999.977,
						CacheSize:       "25600 KB",
						PhysicalID:      0,
						CoreID:          2,
						CPUCores:        10,
						Bugs:            "cpu_meltdown spectre_v1 spectre_v2 spec_store_bypass l1tf mds swapgs itlb_multihit",
					},
					{
						ProcessorNumber: 23,
						VendorID:        "GenuineIntel",
						ModelName:       "Intel(R) Xeon(R) CPU E5-2690 v2 @ 3.00GHz",
						MHz:             2999.951,
						CacheSize:       "25600 KB",
						PhysicalID:      0,
						CoreID:          3,
						CPUCores:        10,
						Bugs:            "cpu_meltdown spectre_v1 spectre_v2 spec_store_bypass l1tf mds swapgs itlb_multihit",
					},
					{
						ProcessorNumber: 24,
						VendorID:        "GenuineIntel",
						ModelName:       "Intel(R) Xeon(R) CPU E5-2690 v2 @ 3.00GHz",
						MHz:             3000.122,
						CacheSize:       "25600 KB",
						PhysicalID:      0,
						CoreID:          4,
						CPUCores:        10,
						Bugs:            "cpu_meltdown spectre_v1 spectre_v2 spec_store_bypass l1tf mds swapgs itlb_multihit",
					},
					{
						ProcessorNumber: 25,
						VendorID:        "GenuineIntel",
						ModelName:       "Intel(R) Xeon(R) CPU E5-2690 v2 @ 3.00GHz",
						MHz:             2999.935,
						CacheSize:       "25600 KB",
						PhysicalID:      0,
						CoreID:          8,
						CPUCores:        10,
						Bugs:            "cpu_meltdown spectre_v1 spectre_v2 spec_store_bypass l1tf mds swapgs itlb_multihit",
					},
					{
						ProcessorNumber: 26,
						VendorID:        "GenuineIntel",
						ModelName:       "Intel(R) Xeon(R) CPU E5-2690 v2 @ 3.00GHz",
						MHz:             3000.435,
						CacheSize:       "25600 KB",
						PhysicalID:      0,
						CoreID:          9,
						CPUCores:        10,
						Bugs:            "cpu_meltdown spectre_v1 spectre_v2 spec_store_bypass l1tf mds swapgs itlb_multihit",
					},
					{
						ProcessorNumber: 27,
						VendorID:        "GenuineIntel",
						ModelName:       "Intel(R) Xeon(R) CPU E5-2690 v2 @ 3.00GHz",
						MHz:             2999.774,
						CacheSize:       "25600 KB",
						PhysicalID:      0,
						CoreID:          10,
						CPUCores:        10,
						Bugs:            "cpu_meltdown spectre_v1 spectre_v2 spec_store_bypass l1tf mds swapgs itlb_multihit",
					},
					{
						ProcessorNumber: 28,
						VendorID:        "GenuineIntel",
						ModelName:       "Intel(R) Xeon(R) CPU E5-2690 v2 @ 3.00GHz",
						MHz:             3000.114,
						CacheSize:       "25600 KB",
						PhysicalID:      0,
						CoreID:          11,
						CPUCores:        10,
						Bugs:            "cpu_meltdown spectre_v1 spectre_v2 spec_store_bypass l1tf mds swapgs itlb_multihit",
					},
					{
						ProcessorNumber: 29,
						VendorID:        "GenuineIntel",
						ModelName:       "Intel(R) Xeon(R) CPU E5-2690 v2 @ 3.00GHz",
						MHz:             2999.940,
						CacheSize:       "25600 KB",
						PhysicalID:      0,
						CoreID:          12,
						CPUCores:        10,
						Bugs:            "cpu_meltdown spectre_v1 spectre_v2 spec_store_bypass l1tf mds swapgs itlb_multihit",
					},
					{
						ProcessorNumber: 30,
						VendorID:        "GenuineIntel",
						ModelName:       "Intel(R) Xeon(R) CPU E5-2690 v2 @ 3.00GHz",
						MHz:             3299.930,
						CacheSize:       "25600 KB",
						PhysicalID:      1,
						CoreID:          0,
						CPUCores:        10,
						Bugs:            "cpu_meltdown spectre_v1 spectre_v2 spec_store_bypass l1tf mds swapgs itlb_multihit",
					},
					{
						ProcessorNumber: 31,
						VendorID:        "GenuineIntel",
						ModelName:       "Intel(R) Xeon(R) CPU E5-2690 v2 @ 3.00GHz",
						MHz:             3299.928,
						CacheSize:       "25600 KB",
						PhysicalID:      1,
						CoreID:          1,
						CPUCores:        10,
						Bugs:            "cpu_meltdown spectre_v1 spectre_v2 spec_store_bypass l1tf mds swapgs itlb_multihit",
					},
					{
						ProcessorNumber: 32,
						VendorID:        "GenuineIntel",
						ModelName:       "Intel(R) Xeon(R) CPU E5-2690 v2 @ 3.00GHz",
						MHz:             3299.967,
						CacheSize:       "25600 KB",
						PhysicalID:      1,
						CoreID:          2,
						CPUCores:        10,
						Bugs:            "cpu_meltdown spectre_v1 spectre_v2 spec_store_bypass l1tf mds swapgs itlb_multihit",
					},
					{
						ProcessorNumber: 33,
						VendorID:        "GenuineIntel",
						ModelName:       "Intel(R) Xeon(R) CPU E5-2690 v2 @ 3.00GHz",
						MHz:             3299.931,
						CacheSize:       "25600 KB",
						PhysicalID:      1,
						CoreID:          3,
						CPUCores:        10,
						Bugs:            "cpu_meltdown spectre_v1 spectre_v2 spec_store_bypass l1tf mds swapgs itlb_multihit",
					},
					{
						ProcessorNumber: 34,
						VendorID:        "GenuineIntel",
						ModelName:       "Intel(R) Xeon(R) CPU E5-2690 v2 @ 3.00GHz",
						MHz:             3299.601,
						CacheSize:       "25600 KB",
						PhysicalID:      1,
						CoreID:          4,
						CPUCores:        10,
						Bugs:            "cpu_meltdown spectre_v1 spectre_v2 spec_store_bypass l1tf mds swapgs itlb_multihit",
					},
					{
						ProcessorNumber: 35,
						VendorID:        "GenuineIntel",
						ModelName:       "Intel(R) Xeon(R) CPU E5-2690 v2 @ 3.00GHz",
						MHz:             3299.930,
						CacheSize:       "25600 KB",
						PhysicalID:      1,
						CoreID:          8,
						CPUCores:        10,
						Bugs:            "cpu_meltdown spectre_v1 spectre_v2 spec_store_bypass l1tf mds swapgs itlb_multihit",
					},
					{
						ProcessorNumber: 36,
						VendorID:        "GenuineIntel",
						ModelName:       "Intel(R) Xeon(R) CPU E5-2690 v2 @ 3.00GHz",
						MHz:             3298.145,
						CacheSize:       "25600 KB",
						PhysicalID:      1,
						CoreID:          9,
						CPUCores:        10,
						Bugs:            "cpu_meltdown spectre_v1 spectre_v2 spec_store_bypass l1tf mds swapgs itlb_multihit",
					},
					{
						ProcessorNumber: 37,
						VendorID:        "GenuineIntel",
						ModelName:       "Intel(R) Xeon(R) CPU E5-2690 v2 @ 3.00GHz",
						MHz:             3305.648,
						CacheSize:       "25600 KB",
						PhysicalID:      1,
						CoreID:          10,
						CPUCores:        10,
						Bugs:            "cpu_meltdown spectre_v1 spectre_v2 spec_store_bypass l1tf mds swapgs itlb_multihit",
					},
					{
						ProcessorNumber: 38,
						VendorID:        "GenuineIntel",
						ModelName:       "Intel(R) Xeon(R) CPU E5-2690 v2 @ 3.00GHz",
						MHz:             3299.965,
						CacheSize:       "25600 KB",
						PhysicalID:      1,
						CoreID:          11,
						CPUCores:        10,
						Bugs:            "cpu_meltdown spectre_v1 spectre_v2 spec_store_bypass l1tf mds swapgs itlb_multihit",
					},
					{
						ProcessorNumber: 39,
						VendorID:        "GenuineIntel",
						ModelName:       "Intel(R) Xeon(R) CPU E5-2690 v2 @ 3.00GHz",
						MHz:             3301.566,
						CacheSize:       "25600 KB",
						PhysicalID:      1,
						CoreID:          12,
						CPUCores:        10,
						Bugs:            "cpu_meltdown spectre_v1 spectre_v2 spec_store_bypass l1tf mds swapgs itlb_multihit",
					},
				},
			},
			wantErr: false,
		},
	}

	for _, testCase := range tests {
		tt := testCase

		t.Run(tt.caseDescription, func(t *testing.T) {
			cpuInfoDataSource := proc.NewCPUInfoDataSource(tt.filePath, "")
			got, err := cpuInfoDataSource.GetData()
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseCPUInfo() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseCPUInfo() got = %v, want %v", got, tt.want)
			}
		})
	}
}
