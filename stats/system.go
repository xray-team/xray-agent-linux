package stats

import (
	"net/http"
	"time"

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

	"github.com/xray-team/xray-agent-linux/conf"
	"github.com/xray-team/xray-agent-linux/dto"
	"github.com/xray-team/xray-agent-linux/logger"
	"github.com/xray-team/xray-agent-linux/run"
	"github.com/xray-team/xray-agent-linux/sys"
)

type Stat struct {
	cfg           *conf.Config
	telemetryChan chan<- *dto.Telemetry
	stopChan      chan bool
}

func New(cfg *conf.Config, telemetryChan chan<- *dto.Telemetry) *Stat {
	return &Stat{cfg: cfg, telemetryChan: telemetryChan, stopChan: make(chan bool)}
}

func (s *Stat) Start() {
	ticker := time.NewTicker(time.Duration(s.cfg.Agent.GetStatIntervalSec) * time.Second)
	defer func() {
		ticker.Stop()
		close(s.telemetryChan)
		close(s.stopChan)
	}()

	// start first time before ticker
	stats, err := s.getStat()
	if err != nil {
		logger.Log.Error.Printf(logger.MessageCollectError, logger.TagAgent, err.Error())
	}
	s.telemetryChan <- stats

	for {
		select {
		case <-ticker.C:
			stats, err := s.getStat()
			if err != nil {
				logger.Log.Error.Printf(logger.MessageCollectError, logger.TagAgent, err.Error())

				continue
			}
			s.telemetryChan <- stats
		case <-s.stopChan:
			return
		}
	}
}

func (s *Stat) Stop() {
	s.stopChan <- true
}

func (s *Stat) Title() string {
	return "stat getter"
}

func (s *Stat) getStat() (*dto.Telemetry, error) {
	var (
		metrics    []dto.Metric
		numMetrics int
		startTime  = time.Now()
		cols       = s.initCollectors()
	)

	hostInfo, err := getHostInfo(s.cfg.Agent)
	if err != nil {
		return nil, err
	}

	for _, collector := range cols {
		// if collector is not enable
		if collector == nil {
			continue
		}

		m, err := s.Collect(collector)
		if err != nil {
			logger.Log.Error.Printf(logger.MessageCollectError, collector.GetName(), err.Error())
		}

		numMetrics += len(m)

		metrics = append(metrics, m...)
	}

	// append self metrics (agent scope)
	if s.cfg.Collectors.EnableSelfMetrics {
		agentSummaryMetrics := s.agentSummaryToMetrics(dto.AgentSummary{
			Duration:      time.Since(startTime),
			MetricsNumber: len(metrics),
		})

		metrics = append(metrics, agentSummaryMetrics...)
	}

	return &dto.Telemetry{
		HostInfo: hostInfo,
		Metrics:  metrics,
	}, nil
}

func (s *Stat) Collect(collector dto.Collector) ([]dto.Metric, error) {
	var (
		m                  []dto.Metric
		err                error
		collectorStartTime = time.Now()
		summary            = dto.CollectorSummary{CollectorName: collector.GetName(), Status: 1}
	)

	func() {
		defer func() {
			if s.cfg.Collectors.EnableSelfMetrics {
				summary.Duration = time.Since(collectorStartTime)
				summary.MetricsNumber = len(m)
				// append self metrics (collector scope)
				m = append(m, s.collectorSummaryToMetrics(summary)...)
			}
		}()

		m, err = collector.Collect()
		if err != nil {
			summary.Status = 2
		}
	}()

	return m, err
}

func (s *Stat) agentSummaryToMetrics(as dto.AgentSummary) []dto.Metric {
	attrs := []dto.MetricAttribute{
		{Name: dto.ResourceAttr, Value: dto.ResourceXraySelf},
		{Name: dto.SetNameSelfScope, Value: dto.SetValueSelfScopeAgent},
	}

	metrics := []dto.Metric{
		{
			Name:       dto.MetricSelfDurationNs,
			Attributes: attrs,
			Value:      as.Duration.Nanoseconds(),
		},
	}

	metrics = append(metrics, dto.Metric{
		Name:       dto.MetricSelfMetricsNumber,
		Attributes: attrs,
		Value:      as.MetricsNumber + len(metrics) + 1,
	})

	return metrics
}

func (s *Stat) collectorSummaryToMetrics(cs dto.CollectorSummary) []dto.Metric {
	attrs := []dto.MetricAttribute{
		{Name: dto.ResourceAttr, Value: dto.ResourceXraySelf},
		{Name: dto.SetNameSelfScope, Value: dto.SetValueSelfScopeCollector},
		{Name: dto.SetNameSelfCollectorName, Value: cs.CollectorName},
	}

	return []dto.Metric{
		// collector status:
		// 	0 - unknown
		//	1 - success
		//	2 - error
		{
			Name:       dto.MetricSelfCollectorState,
			Attributes: attrs,
			Value:      cs.Status,
		},
		{
			Name:       dto.MetricSelfMetricsNumber,
			Attributes: attrs,
			Value:      cs.MetricsNumber,
		},
		{
			Name:       dto.MetricSelfDurationNs,
			Attributes: attrs,
			Value:      cs.Duration.Nanoseconds(),
		},
	}
}

func (s *Stat) initCollectors() []dto.Collector {
	return []dto.Collector{
		// /proc/uptime
		uptime.NewCollector(s.cfg.Collectors.Uptime,
			uptime.NewDataSource(uptime.UptimePath, uptime.CollectorName)),
		// /proc/loadavg
		loadAvg.NewCollector(s.cfg.Collectors.LoadAvg,
			loadAvg.NewDataSource(loadAvg.LoadAvgPath, loadAvg.CollectorName)),
		// PS
		ps.NewCollector(s.cfg.Collectors.PS,
			ps.NewDataSource(ps.ProcPath, ps.CollectorName)),
		// PS stat
		psStat.NewCollector(s.cfg.Collectors.PSStat,
			psStat.NewDataSource(psStat.ProcPath, psStat.CollectorName)),
		// /proc/stat
		stat.NewCollector(s.cfg.Collectors.Stat,
			stat.NewDataSource(stat.StatPath, stat.CollectorName)),
		// /proc/cpuinfo
		cpuInfo.NewCollector(s.cfg.Collectors.CPUInfo,
			cpuInfo.NewDataSource(cpuInfo.CPUInfoPath, cpuInfo.CollectorName)),
		// /proc/meminfo
		memoryInfo.NewCollector(s.cfg.Collectors.MemoryInfo,
			memoryInfo.NewDataSource(memoryInfo.MemInfoPath, memoryInfo.CollectorName)),
		// /proc/diskstat
		diskStat.NewCollector(
			s.cfg.Collectors.DiskStat,
			diskStat.NewBlockDevDataSource(diskStat.DiskStatsPath, diskStat.CollectorName),
			sys.NewClassBlockDataSource(sys.ClassBlockDir, diskStat.CollectorName),
		),
		// disk space
		diskSpace.NewCollector(s.cfg.Collectors.DiskSpace, diskSpace.NewMountsDataSource(diskSpace.MountsPath, diskSpace.CollectorName)),
		// /proc/net/dev
		netDev.NewCollector(
			s.cfg.Collectors.NetDev,
			netDev.NewNetDevDataSource(netDev.NetDevPath, netDev.CollectorName),
			sys.NewClassNetDataSource(sys.ClassNetDir, netDev.CollectorName),
		),
		// /sys/class/net
		netDevStatus.NewCollector(s.cfg.Collectors.NetDevStatus,
			sys.NewClassNetDataSource(sys.ClassNetDir, netDevStatus.CollectorName)),
		// iwconfig
		wireless.NewCollector(
			s.cfg.Collectors.Wireless,
			wireless.NewIwconfigDataSource(run.NewCmdRunner(wireless.CollectorName)),
			sys.NewClassNetDataSource(sys.ClassNetDir, wireless.CollectorName),
		),
		// /proc/net/arp
		netARP.NewCollector(s.cfg.Collectors.NetARP,
			netARP.NewDataSource(netARP.NetArpPath, netARP.CollectorName)),
		// /proc/net/netstat
		netStat.NewCollector(s.cfg.Collectors.NetStat,
			netStat.NewDataSource(netStat.NetStatPath, netStat.CollectorName)),
		// /proc/net/snmp
		netSNMP.NewCollector(s.cfg.Collectors.NetSNMP,
			netStat.NewDataSource(netSNMP.NetSNMPPath, netSNMP.CollectorName)),
		// /proc/net/snmp6
		netSNMP6.NewCollector(s.cfg.Collectors.NetSNMP6,
			netSNMP6.NewDataSource(netSNMP6.NetSNMP6Path, netSNMP6.CollectorName)),
		// mdStat
		mdStat.NewCollector(s.cfg.Collectors.MDStat,
			mdStat.NewDataSource(mdStat.MDStatPath, mdStat.CollectorName)),
		// CMD collector
		cmd.NewCollector(s.cfg.Collectors.CMD,
			run.NewCmdRunner(cmd.CollectorName)),
		// nginx
		nginx.NewStubStatusCollector(s.cfg.Collectors.NginxStubStatus,
			nginx.NewStubStatusClient(
				s.cfg.Collectors.NginxStubStatus,
				&http.Client{
					Timeout: time.Second * time.Duration(s.cfg.Collectors.NginxStubStatus.Timeout),
				},
				nginx.CollectorName,
			),
		),
		// entropy
		entropy.NewCollector(s.cfg.Collectors.Entropy,
			entropy.NewDataSource(entropy.EntropyPath, entropy.CollectorName),
		),
	}
}
