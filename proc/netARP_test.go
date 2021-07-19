package proc_test

import (
	"reflect"
	"testing"

	"xray-agent-linux/dto"
	"xray-agent-linux/logger"
	"xray-agent-linux/proc"
)

func Test_netARPDataSource_GetData(t *testing.T) {
	logger.Init("")

	type args struct {
		filePath string
	}

	tests := []struct {
		name    string
		args    args
		want    []dto.ARPEntry
		wantErr bool
	}{
		/*
			IP address       HW type     Flags       HW address            Mask     Device
			172.31.1.1       0x1         0x2         d2:74:7f:6e:37:e3     *        eth0
			10.0.0.205       0x1         0x2         00:16:3e:41:52:5c     *        lxdbr0
		*/
		{
			name: "success",
			args: args{"./testfiles/netarp/netarp-ubuntu-2gb-nbg1-1_4.15.0-66-generic"},
			want: []dto.ARPEntry{
				{IP: "172.31.1.1", HWType: "0x1", Flags: "0x2", HWAddress: "d2:74:7f:6e:37:e3", Mask: "*", Device: "eth0"},
				{IP: "10.0.0.205", HWType: "0x1", Flags: "0x2", HWAddress: "00:16:3e:41:52:5c", Mask: "*", Device: "lxdbr0"},
			},
			wantErr: false,
		},
		{
			name: "success additional columns",
			args: args{"./testfiles/netarp/netarp-ubuntu-2gb-nbg1-1_4.15.0-66-generic-add-col"},
			want: []dto.ARPEntry{
				{IP: "172.31.1.1", HWType: "0x1", Flags: "0x2", HWAddress: "d2:74:7f:6e:37:e3", Mask: "*", Device: "eth0"},
				{IP: "10.0.0.205", HWType: "0x1", Flags: "0x2", HWAddress: "00:16:3e:41:52:5c", Mask: "*", Device: "lxdbr0"},
			},
			wantErr: false,
		},
		{
			name:    "wrong headers",
			args:    args{"./testfiles/netarp/netarp-wrong-headers"},
			wantErr: true,
		},
		{
			name:    "wrong number of columns",
			args:    args{"./testfiles/netarp/netarp-ubuntu-wrong-num-col"},
			wantErr: true,
		},
		{
			name:    "no file",
			args:    args{"./testfiles/netarp/nofile"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			netARPDataSource := proc.NewNetARPDataSource(tt.args.filePath, "")
			got, err := netARPDataSource.GetData()
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseNetArp() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			if (got == nil && tt.want != nil) || (got != nil && tt.want == nil) {
				t.Errorf("ParseNetArp() \n= %+v\nwant %+v", got, tt.want)

				return
			} else if got == nil && tt.want == nil {
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseNetArp() \n= %+v\nwant %+v", got, tt.want)
			}
		})
	}
}
