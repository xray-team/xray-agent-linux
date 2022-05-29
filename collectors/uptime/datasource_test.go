package uptime_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/xray-team/xray-agent-linux/collectors/uptime"
	"github.com/xray-team/xray-agent-linux/logger"
)

func Test_uptimeDataSource_GetData(t *testing.T) {
	logger.Init()

	type args struct {
		filePath string
	}

	tests := []struct {
		name string
		args args
		want *uptime.Uptime
		err  error
	}{
		{
			name: "success",
			args: args{filePath: "./testFiles/uptime"},
			want: &uptime.Uptime{Uptime: 117332.61, Idle: 116056.85},
			err:  nil,
		},
		{
			name: "success 210 days",
			args: args{filePath: "./testFiles/uptime-210days"},
			want: &uptime.Uptime{Uptime: 18144238.28, Idle: 286741938.89},
			err:  nil,
		},
		{
			name: "Error no file",
			args: args{filePath: "./testFiles/nofile"},
			err:  fmt.Errorf("cannot read file uptime file ./testFiles/nofile. open testFiles/nofile: no such file or directory"),
		},
		{
			name: "Error not valid number of fields",
			args: args{filePath: "./testFiles/uptime_not_valid_number"},
			err:  fmt.Errorf("not valid number of fields in uptime file, needs 2: ./testFiles/uptime_not_valid_number"),
		},
		{
			name: "Error not valid format Uptime",
			args: args{filePath: "./testFiles/uptime_not_valid_format_uptime"},
			err:  fmt.Errorf("not format of Uptime field in uptime file, needs float: ./testFiles/uptime_not_valid_format_uptime"),
		},
		{
			name: "Error not valid format Idle",
			args: args{filePath: "./testFiles/uptime_not_valid_format_idle"},
			err:  fmt.Errorf("not format of Idle field in uptime file, needs float: ./testFiles/uptime_not_valid_format_idle"),
		},
	}

	for _, testCase := range tests {
		tt := testCase

		t.Run(tt.name, func(t *testing.T) {
			uptimeDataSource := uptime.NewDataSource(tt.args.filePath, "")
			got, err := uptimeDataSource.GetData()
			if fmt.Sprintf("%s", err) != fmt.Sprintf("%s", tt.err) {
				t.Errorf("ParseUptime() error = %v, wantErr %v", err, tt.err)

				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseUptime() = %v, want %v", got, tt.want)
			}
		})
	}
}
