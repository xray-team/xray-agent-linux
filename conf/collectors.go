package conf

import (
	"github.com/xray-team/xray-agent-linux/collectors/cmd"
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
	"github.com/xray-team/xray-agent-linux/collectors/nginx"
	"github.com/xray-team/xray-agent-linux/collectors/ps"
	"github.com/xray-team/xray-agent-linux/collectors/psStat"
	"github.com/xray-team/xray-agent-linux/collectors/stat"
	"github.com/xray-team/xray-agent-linux/collectors/uptime"
	"github.com/xray-team/xray-agent-linux/collectors/wireless"
)

type CollectorsConf struct {
	EnableSelfMetrics bool                 `json:"enableSelfMetrics"`
	Uptime            *uptime.Config       `json:"uptime"`
	LoadAvg           *loadAvg.Config      `json:"loadAvg"`
	PS                *ps.Config           `json:"ps"`
	PSStat            *psStat.Config       `json:"psStat"`
	Stat              *stat.Config         `json:"stat"`
	CPUInfo           *cpuInfo.Config      `json:"cpuInfo"`
	MemoryInfo        *memoryInfo.Config   `json:"memoryInfo"`
	DiskStat          *diskStat.Config     `json:"diskStat"`
	DiskSpace         *diskSpace.Config    `json:"diskSpace"`
	NetDev            *netDev.Config       `json:"netDev"`
	NetDevStatus      *netDevStatus.Config `json:"netDevStatus"`
	NetARP            *netARP.Config       `json:"netARP"`
	NetStat           *netStat.Config      `json:"netStat"`
	NetSNMP           *netSNMP.Config      `json:"netSNMP"`
	NetSNMP6          *netSNMP6.Config     `json:"netSNMP6"`
	MDStat            *mdStat.Config       `json:"mdStat"`
	CMD               *cmd.Config          `json:"cmd"`
	NginxStubStatus   *nginx.Config        `json:"nginxStubStatus"`
	Wireless          *wireless.Config     `json:"wireless"`
	Entropy           *entropy.Config      `json:"entropy"`
}
