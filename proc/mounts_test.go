package proc_test

import (
	"reflect"
	"testing"

	"github.com/xray-team/xray-agent-linux/dto"
	"github.com/xray-team/xray-agent-linux/logger"
	"github.com/xray-team/xray-agent-linux/proc"
)

func Test_mountsDataSource_GetData(t *testing.T) {
	logger.Init()

	tests := []struct {
		caseDescription string
		filePath        string
		want            []dto.Mounts
		wantErr         bool
	}{
		{
			caseDescription: "no file",
			filePath:        "./testfiles/mounts/nofile",
			want:            []dto.Mounts{},
			wantErr:         true,
		},
		{
			caseDescription: "kernel4.15",
			filePath:        "./testfiles/mounts/mounts-kernel4.15.0-66-generic",
			want: []dto.Mounts{
				{
					Dev:            "sysfs",
					MountPoint:     "/sys",
					FileSystemType: "sysfs",
					MountOptions:   "rw,nosuid,nodev,noexec,relatime",
					Dump:           0,
					Pass:           0,
				},
				{
					Dev:            "proc",
					MountPoint:     "/proc",
					FileSystemType: "proc",
					MountOptions:   "rw,nosuid,nodev,noexec,relatime",
					Dump:           0,
					Pass:           0,
				},
				{
					Dev:            "udev",
					MountPoint:     "/dev",
					FileSystemType: "devtmpfs",
					MountOptions:   "rw,nosuid,relatime,size=981048k,nr_inodes=245262,mode=755",
					Dump:           0,
					Pass:           0,
				},
				{
					Dev:            "devpts",
					MountPoint:     "/dev/pts",
					FileSystemType: "devpts",
					MountOptions:   "rw,nosuid,noexec,relatime,gid=5,mode=620,ptmxmode=000",
					Dump:           0,
					Pass:           0,
				},
				{
					Dev:            "tmpfs",
					MountPoint:     "/run",
					FileSystemType: "tmpfs",
					MountOptions:   "rw,nosuid,noexec,relatime,size=199212k,mode=755",
					Dump:           0,
					Pass:           0,
				},
				{
					Dev:            "/dev/sda1",
					MountPoint:     "/",
					FileSystemType: "ext4",
					MountOptions:   "rw,relatime,errors=remount-ro,data=ordered",
					Dump:           0,
					Pass:           0,
				},
				{
					Dev:            "securityfs",
					MountPoint:     "/sys/kernel/security",
					FileSystemType: "securityfs",
					MountOptions:   "rw,nosuid,nodev,noexec,relatime",
					Dump:           0,
					Pass:           0,
				},
				{
					Dev:            "tmpfs",
					MountPoint:     "/dev/shm",
					FileSystemType: "tmpfs",
					MountOptions:   "rw,nosuid,nodev",
					Dump:           0,
					Pass:           0,
				},
				{
					Dev:            "tmpfs",
					MountPoint:     "/run/lock",
					FileSystemType: "tmpfs",
					MountOptions:   "rw,nosuid,nodev,noexec,relatime,size=5120k",
					Dump:           0,
					Pass:           0,
				},
				{
					Dev:            "tmpfs",
					MountPoint:     "/sys/fs/cgroup",
					FileSystemType: "tmpfs",
					MountOptions:   "ro,nosuid,nodev,noexec,mode=755",
					Dump:           0,
					Pass:           0,
				},
				{
					Dev:            "cgroup",
					MountPoint:     "/sys/fs/cgroup/unified",
					FileSystemType: "cgroup2",
					MountOptions:   "rw,nosuid,nodev,noexec,relatime",
					Dump:           0,
					Pass:           0,
				},
				{
					Dev:            "cgroup",
					MountPoint:     "/sys/fs/cgroup/systemd",
					FileSystemType: "cgroup",
					MountOptions:   "rw,nosuid,nodev,noexec,relatime,xattr,name=systemd",
					Dump:           0,
					Pass:           0,
				},
				{
					Dev:            "pstore",
					MountPoint:     "/sys/fs/pstore",
					FileSystemType: "pstore",
					MountOptions:   "rw,nosuid,nodev,noexec,relatime",
					Dump:           0,
					Pass:           0,
				},
				{
					Dev:            "cgroup",
					MountPoint:     "/sys/fs/cgroup/hugetlb",
					FileSystemType: "cgroup",
					MountOptions:   "rw,nosuid,nodev,noexec,relatime,hugetlb",
					Dump:           0,
					Pass:           0,
				},
				{
					Dev:            "cgroup",
					MountPoint:     "/sys/fs/cgroup/memory",
					FileSystemType: "cgroup",
					MountOptions:   "rw,nosuid,nodev,noexec,relatime,memory",
					Dump:           0,
					Pass:           0,
				},
				{
					Dev:            "cgroup",
					MountPoint:     "/sys/fs/cgroup/net_cls,net_prio",
					FileSystemType: "cgroup",
					MountOptions:   "rw,nosuid,nodev,noexec,relatime,net_cls,net_prio",
					Dump:           0,
					Pass:           0,
				},
				{
					Dev:            "cgroup",
					MountPoint:     "/sys/fs/cgroup/freezer",
					FileSystemType: "cgroup",
					MountOptions:   "rw,nosuid,nodev,noexec,relatime,freezer",
					Dump:           0,
					Pass:           0,
				},
				{
					Dev:            "cgroup",
					MountPoint:     "/sys/fs/cgroup/pids",
					FileSystemType: "cgroup",
					MountOptions:   "rw,nosuid,nodev,noexec,relatime,pids",
					Dump:           0,
					Pass:           0,
				},
				{
					Dev:            "cgroup",
					MountPoint:     "/sys/fs/cgroup/blkio",
					FileSystemType: "cgroup",
					MountOptions:   "rw,nosuid,nodev,noexec,relatime,blkio",
					Dump:           0,
					Pass:           0,
				},
				{
					Dev:            "cgroup",
					MountPoint:     "/sys/fs/cgroup/cpuset",
					FileSystemType: "cgroup",
					MountOptions:   "rw,nosuid,nodev,noexec,relatime,cpuset",
					Dump:           0,
					Pass:           0,
				},
				{
					Dev:            "cgroup",
					MountPoint:     "/sys/fs/cgroup/cpu,cpuacct",
					FileSystemType: "cgroup",
					MountOptions:   "rw,nosuid,nodev,noexec,relatime,cpu,cpuacct",
					Dump:           0,
					Pass:           0,
				},
				{
					Dev:            "cgroup",
					MountPoint:     "/sys/fs/cgroup/rdma",
					FileSystemType: "cgroup",
					MountOptions:   "rw,nosuid,nodev,noexec,relatime,rdma",
					Dump:           0,
					Pass:           0,
				},
				{
					Dev:            "cgroup",
					MountPoint:     "/sys/fs/cgroup/perf_event",
					FileSystemType: "cgroup",
					MountOptions:   "rw,nosuid,nodev,noexec,relatime,perf_event",
					Dump:           0,
					Pass:           0,
				},
				{
					Dev:            "cgroup",
					MountPoint:     "/sys/fs/cgroup/devices",
					FileSystemType: "cgroup",
					MountOptions:   "rw,nosuid,nodev,noexec,relatime,devices",
					Dump:           0,
					Pass:           0,
				},
				{
					Dev:            "systemd-1",
					MountPoint:     "/proc/sys/fs/binfmt_misc",
					FileSystemType: "autofs",
					MountOptions:   "rw,relatime,fd=27,pgrp=1,timeout=0,minproto=5,maxproto=5,direct,pipe_ino=11793",
					Dump:           0,
					Pass:           0,
				},
				{
					Dev:            "mqueue",
					MountPoint:     "/dev/mqueue",
					FileSystemType: "mqueue",
					MountOptions:   "rw,relatime",
					Dump:           0,
					Pass:           0,
				},
				{
					Dev:            "hugetlbfs",
					MountPoint:     "/dev/hugepages",
					FileSystemType: "hugetlbfs",
					MountOptions:   "rw,relatime,pagesize=2M",
					Dump:           0,
					Pass:           0,
				},
				{
					Dev:            "debugfs",
					MountPoint:     "/sys/kernel/debug",
					FileSystemType: "debugfs",
					MountOptions:   "rw,relatime",
					Dump:           0,
					Pass:           0,
				},
				{
					Dev:            "configfs",
					MountPoint:     "/sys/kernel/config",
					FileSystemType: "configfs",
					MountOptions:   "rw,relatime",
					Dump:           0,
					Pass:           0,
				},
				{
					Dev:            "fusectl",
					MountPoint:     "/sys/fs/fuse/connections",
					FileSystemType: "fusectl",
					MountOptions:   "rw,relatime",
					Dump:           0,
					Pass:           0,
				},
				{
					Dev:            "lxcfs",
					MountPoint:     "/var/lib/lxcfs",
					FileSystemType: "fuse.lxcfs",
					MountOptions:   "rw,nosuid,nodev,relatime,user_id=0,group_id=0,allow_other",
					Dump:           0,
					Pass:           0,
				},
				{
					Dev:            "overlay",
					MountPoint:     "/var/lib/docker/overlay2/9a146b0ad88863919f7a648350b575ddb8c8db8d085379571c016c40e128c467/merged",
					FileSystemType: "overlay",
					MountOptions:   "rw,relatime,lowerdir=/var/lib/docker/overlay2/l/4LZUQH3NOVAU4DFA2IRFKDLZTN:/var/lib/docker/overlay2/l/G4D3QANSYH5LXLHCDVBGJCSTNY:/var/lib/docker/overlay2/l/6FYSJWWHOGUIXE22T336QBY5SF:/var/lib/docker/overlay2/l/4GJJLJPUAF6CV6GIJN57RSZ2K4:/var/lib/docker/overlay2/l/QBD5CL7O56V2VASB6L734DSDVN:/var/lib/docker/overlay2/l/P7ZGRU3UAKYD6OFMYP2T5GMLM7:/var/lib/docker/overlay2/l/ZULR3Q762S5RPX2OJJ557MEA7H,upperdir=/var/lib/docker/overlay2/9a146b0ad88863919f7a648350b575ddb8c8db8d085379571c016c40e128c467/diff,workdir=/var/lib/docker/overlay2/9a146b0ad88863919f7a648350b575ddb8c8db8d085379571c016c40e128c467/work",
					Dump:           0,
					Pass:           0,
				},
				{
					Dev:            "shm",
					MountPoint:     "/var/lib/docker/containers/c63adfb53a8fd1ec5922c0f6584e4d149b98ca3e36bc70c0d0a397b28e805a29/mounts/shm",
					FileSystemType: "tmpfs",
					MountOptions:   "rw,nosuid,nodev,noexec,relatime,size=65536k",
					Dump:           0,
					Pass:           0,
				},
				{
					Dev:            "nsfs",
					MountPoint:     "/run/docker/netns/95cbecae489b",
					FileSystemType: "nsfs",
					MountOptions:   "rw",
					Dump:           0,
					Pass:           0,
				},
				{
					Dev:            "tmpfs",
					MountPoint:     "/var/lib/lxd/shmounts",
					FileSystemType: "tmpfs",
					MountOptions:   "rw,relatime,size=100k,mode=711",
					Dump:           0,
					Pass:           0,
				},
				{
					Dev:            "tmpfs",
					MountPoint:     "/var/lib/lxd/devlxd",
					FileSystemType: "tmpfs",
					MountOptions:   "rw,relatime,size=100k,mode=755",
					Dump:           0,
					Pass:           0,
				},
				{
					Dev:            "binfmt_misc",
					MountPoint:     "/proc/sys/fs/binfmt_misc",
					FileSystemType: "binfmt_misc",
					MountOptions:   "rw,relatime",
					Dump:           0,
					Pass:           0,
				},
				{
					Dev:            "tmpfs",
					MountPoint:     "/run/user/1000",
					FileSystemType: "tmpfs",
					MountOptions:   "rw,nosuid,nodev,relatime,size=199208k,mode=700,uid=1000,gid=1000",
					Dump:           0,
					Pass:           0,
				},
			},
			wantErr: false,
		},
	}

	for _, testCase := range tests {
		tt := testCase

		t.Run(tt.caseDescription, func(t *testing.T) {
			mountDataSource := proc.NewMountsDataSource(tt.filePath, "")
			got, err := mountDataSource.GetData()

			if (err != nil) != tt.wantErr {
				t.Errorf("ParseMounts() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseMounts() got  = %v, want = %v", got, tt.want)
			}
		})
	}
}
