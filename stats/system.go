package stats

import (
	"net/http"
	"path/filepath"
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
	"github.com/xray-team/xray-agent-linux/proc"
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
		uptime.NewUptimeCollector(s.cfg.Collectors,
			uptime.NewUptimeDataSource(filepath.Join(s.cfg.Collectors.RootPath, proc.ProcPath, uptime.UptimePath), uptime.CollectorName)),
		// /proc/loadavg
		loadAvg.NewLoadAvgCollector(s.cfg.Collectors,
			loadAvg.NewLoadAvgDataSource(filepath.Join(s.cfg.Collectors.RootPath, proc.ProcPath, loadAvg.LoadAvgPath), loadAvg.CollectorName)),
		// PS
		ps.NewPSCollector(s.cfg.Collectors,
			ps.NewPSDataSource(filepath.Join(s.cfg.Collectors.RootPath, proc.ProcPath), ps.CollectorName)),
		// PS stat
		psStat.NewPSStatCollector(s.cfg.Collectors,
			psStat.NewPSStatDataSource(filepath.Join(s.cfg.Collectors.RootPath, proc.ProcPath), psStat.CollectorName)),
		// /proc/stat
		stat.NewStatCollector(s.cfg.Collectors,
			stat.NewStatDataSource(filepath.Join(s.cfg.Collectors.RootPath, proc.ProcPath, stat.StatPath), stat.CollectorName)),
		// /proc/cpuinfo
		cpuInfo.NewCpuInfoCollector(s.cfg.Collectors,
			cpuInfo.NewCPUInfoDataSource(filepath.Join(s.cfg.Collectors.RootPath, proc.ProcPath, cpuInfo.CPUInfoPath), cpuInfo.CollectorName)),
		// /proc/meminfo
		memoryInfo.NewMemoryInfoCollector(s.cfg.Collectors,
			memoryInfo.NewMemoryDataSource(filepath.Join(s.cfg.Collectors.RootPath, proc.ProcPath, memoryInfo.MemInfoPath), memoryInfo.CollectorName)),
		// /proc/diskstat
		diskStat.NewDiskStatCollector(
			s.cfg.Collectors,
			diskStat.NewBlockDevDataSource(filepath.Join(s.cfg.Collectors.RootPath, proc.ProcPath, diskStat.DiskStatsPath), diskStat.CollectorName),
			sys.NewClassBlockDataSource(filepath.Join(s.cfg.Collectors.RootPath, sys.ClassBlockDir), diskStat.CollectorName),
		),
		// disk space
		diskSpace.NewDiskSpaceCollector(s.cfg.Collectors, diskSpace.NewMountsDataSource(filepath.Join(s.cfg.Collectors.RootPath, proc.ProcPath, diskSpace.MountsPath), diskSpace.CollectorName)),
		// /proc/net/dev
		netDev.NewNetDevCollector(
			s.cfg.Collectors,
			netDev.NewNetDevDataSource(filepath.Join(s.cfg.Collectors.RootPath, proc.ProcPath, netDev.NetDevPath), netDev.CollectorName),
			sys.NewClassNetDataSource(filepath.Join(s.cfg.Collectors.RootPath, sys.ClassNetDir), netDev.CollectorName),
		),
		// /sys/class/net
		netDevStatus.NewNetDevStatusCollector(s.cfg.Collectors,
			sys.NewClassNetDataSource(filepath.Join(s.cfg.Collectors.RootPath, sys.ClassNetDir), netDevStatus.CollectorName)),
		// iwconfig
		wireless.NewWirelessCollector(
			s.cfg.Collectors,
			wireless.NewIwconfigDataSource(run.NewCmdRunner(wireless.CollectorName)),
			sys.NewClassNetDataSource(filepath.Join(s.cfg.Collectors.RootPath, sys.ClassNetDir), wireless.CollectorName),
		),
		// /proc/net/arp
		netARP.NewNetARPCollector(s.cfg.Collectors,
			netARP.NewNetARPDataSource(filepath.Join(s.cfg.Collectors.RootPath, proc.ProcPath, netARP.NetArpPath), netARP.CollectorName)),
		// /proc/net/netstat
		netStat.NewNetStatCollector(s.cfg.Collectors,
			netStat.NewNetStatDataSource(filepath.Join(s.cfg.Collectors.RootPath, proc.ProcPath, netStat.NetStatPath), netStat.CollectorName)),
		// /proc/net/snmp
		netSNMP.NewNetSNMPCollector(s.cfg.Collectors,
			netStat.NewNetStatDataSource(filepath.Join(s.cfg.Collectors.RootPath, proc.ProcPath, netSNMP.NetSNMPPath), netSNMP.CollectorName)),
		// /proc/net/snmp6
		netSNMP6.NewNetSNMP6Collector(s.cfg.Collectors,
			netSNMP6.NewNetSNMP6DataSource(filepath.Join(s.cfg.Collectors.RootPath, proc.ProcPath, netSNMP6.NetSNMP6Path), netSNMP6.CollectorName)),
		// mdStat
		mdStat.NewMDStatCollector(s.cfg.Collectors,
			mdStat.NewMDStatDataSource(filepath.Join(s.cfg.Collectors.RootPath, mdStat.MDStatPath), mdStat.CollectorName)),
		// CMD collector
		cmd.NewCmdCollector(s.cfg.Collectors,
			run.NewCmdRunner(cmd.CollectorName)),
		// nginx
		nginx.NewNginxStubStatusCollector(s.cfg.Collectors,
			nginx.NewStubStatusClient(
				s.cfg.Collectors.NginxStubStatus,
				&http.Client{
					Timeout: time.Second * time.Duration(s.cfg.Collectors.NginxStubStatus.Timeout),
				},
				nginx.CollectorName,
			),
		),
		// entropy
		entropy.NewEntropyCollector(s.cfg.Collectors,
			entropy.NewEntropyDataSource(filepath.Join(s.cfg.Collectors.RootPath, proc.ProcPath, entropy.EntropyPath), entropy.CollectorName),
		),
	}
}
