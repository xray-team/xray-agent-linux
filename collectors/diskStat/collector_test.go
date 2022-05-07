package diskStat

import (
	"reflect"
	"testing"

	"github.com/xray-team/xray-agent-linux/conf"
	"github.com/xray-team/xray-agent-linux/dto"
)

func TestDiskStatCollector_excludeBlockDevByName(t *testing.T) {
	type fields struct {
		Config               *conf.DiskStatConf
		DataSource           DiskStatDataSource
		ClassBlockDataSource ClassBlockDataSource
	}
	type args struct {
		m map[string]dto.ClassBlock
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   map[string]dto.ClassBlock
	}{
		{
			name: "first",
			fields: fields{
				Config: &conf.DiskStatConf{
					ExcludeByName: []string{"sda1"},
				},
			},
			args: args{
				m: map[string]dto.ClassBlock{
					"sda1": {Uevent: &dto.BlockDevUevent{}},
					"sda2": {Uevent: &dto.BlockDevUevent{}},
					"sda3": {Uevent: &dto.BlockDevUevent{}},
				},
			},
			want: map[string]dto.ClassBlock{
				"sda2": {Uevent: &dto.BlockDevUevent{}},
				"sda3": {Uevent: &dto.BlockDevUevent{}},
			},
		},
		{
			name: "middle",
			fields: fields{
				Config: &conf.DiskStatConf{
					ExcludeByName: []string{"sda2"},
				},
			},
			args: args{
				m: map[string]dto.ClassBlock{
					"sda1": {Uevent: &dto.BlockDevUevent{}},
					"sda2": {Uevent: &dto.BlockDevUevent{}},
					"sda3": {Uevent: &dto.BlockDevUevent{}},
				},
			},
			want: map[string]dto.ClassBlock{
				"sda1": {Uevent: &dto.BlockDevUevent{}},
				"sda3": {Uevent: &dto.BlockDevUevent{}},
			},
		},
		{
			name: "last",
			fields: fields{
				Config: &conf.DiskStatConf{
					ExcludeByName: []string{"sda3"},
				},
			},
			args: args{
				m: map[string]dto.ClassBlock{
					"sda1": {Uevent: &dto.BlockDevUevent{}},
					"sda2": {Uevent: &dto.BlockDevUevent{}},
					"sda3": {Uevent: &dto.BlockDevUevent{}},
				},
			},
			want: map[string]dto.ClassBlock{
				"sda1": {Uevent: &dto.BlockDevUevent{}},
				"sda2": {Uevent: &dto.BlockDevUevent{}},
			},
		},
		{
			name: "nothing to exclude",
			fields: fields{
				Config: &conf.DiskStatConf{
					ExcludeByName: []string{"sdb"},
				},
			},
			args: args{
				m: map[string]dto.ClassBlock{
					"sda1": {Uevent: &dto.BlockDevUevent{}},
					"sda2": {Uevent: &dto.BlockDevUevent{}},
					"sda3": {Uevent: &dto.BlockDevUevent{}},
				},
			},
			want: map[string]dto.ClassBlock{
				"sda1": {Uevent: &dto.BlockDevUevent{}},
				"sda2": {Uevent: &dto.BlockDevUevent{}},
				"sda3": {Uevent: &dto.BlockDevUevent{}},
			},
		},
		{
			name: "nothing to exclude - 2",
			fields: fields{
				Config: &conf.DiskStatConf{
					ExcludeByName: []string{"sda1"},
				},
			},
			args: args{
				m: map[string]dto.ClassBlock{},
			},
			want: map[string]dto.ClassBlock{},
		},
		{
			name: "exclude all",
			fields: fields{
				Config: &conf.DiskStatConf{
					ExcludeByName: []string{"sda1", "sda3", "sda2"},
				},
			},
			args: args{
				m: map[string]dto.ClassBlock{
					"sda1": {Uevent: &dto.BlockDevUevent{}},
					"sda2": {Uevent: &dto.BlockDevUevent{}},
					"sda3": {Uevent: &dto.BlockDevUevent{}},
				},
			},
			want: map[string]dto.ClassBlock{},
		},
		{
			name: "exclude 2",
			fields: fields{
				Config: &conf.DiskStatConf{
					ExcludeByName: []string{"sda1", "sda3"},
				},
			},
			args: args{
				m: map[string]dto.ClassBlock{
					"sda1": {Uevent: &dto.BlockDevUevent{}},
					"sda2": {Uevent: &dto.BlockDevUevent{}},
					"sda3": {Uevent: &dto.BlockDevUevent{}},
				},
			},
			want: map[string]dto.ClassBlock{
				"sda2": {Uevent: &dto.BlockDevUevent{}},
			},
		},
	}

	for _, testCase := range tests {
		tt := testCase

		t.Run(tt.name, func(t *testing.T) {
			c := &DiskStatCollector{
				Config:               tt.fields.Config,
				DataSource:           tt.fields.DataSource,
				ClassBlockDataSource: tt.fields.ClassBlockDataSource,
			}
			if got := c.excludeBlockDevByName(tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("excludeBlockDevByName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDiskStatCollector_excludeBlockDevPartitions(t *testing.T) {
	type fields struct {
		Config               *conf.DiskStatConf
		DataSource           DiskStatDataSource
		ClassBlockDataSource ClassBlockDataSource
	}
	type args struct {
		m map[string]dto.ClassBlock
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   map[string]dto.ClassBlock
	}{
		{
			name: "exclude partitions",
			fields: fields{
				Config: &conf.DiskStatConf{
					ExcludePartitions: true,
				},
			},
			args: args{
				m: map[string]dto.ClassBlock{
					"sda":  {Uevent: &dto.BlockDevUevent{DevType: dto.BlockDevTypeDisk}},
					"sda1": {Uevent: &dto.BlockDevUevent{DevType: dto.BlockDevTypePartition}},
					"sda2": {Uevent: &dto.BlockDevUevent{DevType: dto.BlockDevTypePartition}},
					"sda3": {Uevent: &dto.BlockDevUevent{DevType: dto.BlockDevTypePartition}},
				},
			},
			want: map[string]dto.ClassBlock{
				"sda": {Uevent: &dto.BlockDevUevent{DevType: dto.BlockDevTypeDisk}},
			},
		},
		{
			name: "exclude partitions - 2",
			fields: fields{
				Config: &conf.DiskStatConf{
					ExcludePartitions: true,
				},
			},
			args: args{
				m: map[string]dto.ClassBlock{
					"sda":  {Uevent: &dto.BlockDevUevent{DevType: dto.BlockDevTypeDisk}},
					"sda1": {Uevent: &dto.BlockDevUevent{DevType: dto.BlockDevTypePartition}},
					"sdb":  {Uevent: &dto.BlockDevUevent{DevType: dto.BlockDevTypeDisk}},
					"sdb1": {Uevent: &dto.BlockDevUevent{DevType: dto.BlockDevTypePartition}},
				},
			},
			want: map[string]dto.ClassBlock{
				"sda": {Uevent: &dto.BlockDevUevent{DevType: dto.BlockDevTypeDisk}},
				"sdb": {Uevent: &dto.BlockDevUevent{DevType: dto.BlockDevTypeDisk}},
			},
		},
		{
			name: "not exclude partitions",
			fields: fields{
				Config: &conf.DiskStatConf{
					ExcludePartitions: false,
				},
			},
			args: args{
				m: map[string]dto.ClassBlock{
					"sda":  {Uevent: &dto.BlockDevUevent{DevType: dto.BlockDevTypeDisk}},
					"sda1": {Uevent: &dto.BlockDevUevent{DevType: dto.BlockDevTypePartition}},
					"sda2": {Uevent: &dto.BlockDevUevent{DevType: dto.BlockDevTypePartition}},
				},
			},
			want: map[string]dto.ClassBlock{
				"sda":  {Uevent: &dto.BlockDevUevent{DevType: dto.BlockDevTypeDisk}},
				"sda1": {Uevent: &dto.BlockDevUevent{DevType: dto.BlockDevTypePartition}},
				"sda2": {Uevent: &dto.BlockDevUevent{DevType: dto.BlockDevTypePartition}},
			},
		},
	}

	for _, testCase := range tests {
		tt := testCase

		t.Run(tt.name, func(t *testing.T) {
			c := &DiskStatCollector{
				Config:               tt.fields.Config,
				DataSource:           tt.fields.DataSource,
				ClassBlockDataSource: tt.fields.ClassBlockDataSource,
			}

			if got := c.excludeBlockDevPartitions(tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("excludeBlockDevPartitions() = %v, want %v", got, tt.want)
			}
		})
	}
}
