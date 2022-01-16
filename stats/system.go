package stats

import (
	"net/http"
	"path/filepath"
	"time"

	"github.com/xray-team/xray-agent-linux/collectors"
	"github.com/xray-team/xray-agent-linux/conf"
	"github.com/xray-team/xray-agent-linux/dto"
	"github.com/xray-team/xray-agent-linux/logger"
	"github.com/xray-team/xray-agent-linux/mdstat"
	"github.com/xray-team/xray-agent-linux/nginx"
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
		logger.LogWarning("getStat error", err)
	}
	s.telemetryChan <- stats

	for {
		select {
		case <-ticker.C:
			stats, err := s.getStat()
			if err != nil {
				logger.LogWarning("getStat error", err)

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
			logger.LogWarning("collect error", err)
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
		collectors.NewUptimeCollector(s.cfg.Collectors,
			proc.NewUptimeDataSource(filepath.Join(s.cfg.Collectors.RootPath, proc.ProcPath, proc.UptimePath), dto.CollectorNameUptime)),
		// /proc/loadavg
		collectors.NewLoadAvgCollector(s.cfg.Collectors,
			proc.NewLoadAvgDataSource(filepath.Join(s.cfg.Collectors.RootPath, proc.ProcPath, proc.LoadAvgPath), dto.CollectorNameLoadAvg)),
		// PS
		collectors.NewPSCollector(s.cfg.Collectors,
			proc.NewPSDataSource(filepath.Join(s.cfg.Collectors.RootPath, proc.ProcPath), dto.CollectorNamePS)),
		// PS stat
		collectors.NewPSStatCollector(s.cfg.Collectors,
			proc.NewPSStatDataSource(filepath.Join(s.cfg.Collectors.RootPath, proc.ProcPath), dto.CollectorNamePSStat)),
		// /proc/stat
		collectors.NewStatCollector(s.cfg.Collectors,
			proc.NewStatDataSource(filepath.Join(s.cfg.Collectors.RootPath, proc.ProcPath, proc.StatPath), dto.CollectorNameStat)),
		// /proc/cpuinfo
		collectors.NewCpuInfoCollector(s.cfg.Collectors,
			proc.NewCPUInfoDataSource(filepath.Join(s.cfg.Collectors.RootPath, proc.ProcPath, proc.CPUInfoPath), dto.CollectorNameCPUInfo)),
		// /proc/meminfo
		collectors.NewMemoryInfoCollector(s.cfg.Collectors,
			proc.NewMemoryDataSource(filepath.Join(s.cfg.Collectors.RootPath, proc.ProcPath, proc.MemInfoPath), dto.CollectorNameMemoryInfo)),
		// /proc/diskstat
		collectors.NewDiskStatCollector(
			s.cfg.Collectors,
			proc.NewBlockDevDataSource(filepath.Join(s.cfg.Collectors.RootPath, proc.ProcPath, proc.DiskStatsPath), dto.CollectorNameDiskStat),
			sys.NewClassBlockDataSource(filepath.Join(s.cfg.Collectors.RootPath, sys.ClassBlockDir), dto.CollectorNameDiskStat),
		),
		// disk space
		collectors.NewDiskSpaceCollector(s.cfg.Collectors, proc.NewMountsDataSource(filepath.Join(s.cfg.Collectors.RootPath, proc.ProcPath, proc.MountsPath), dto.CollectorNameDiskSpace)),
		// /proc/net/dev
		collectors.NewNetDevCollector(
			s.cfg.Collectors,
			proc.NewNetDevDataSource(filepath.Join(s.cfg.Collectors.RootPath, proc.ProcPath, proc.NetDevPath), dto.CollectorNameNetDev),
			sys.NewClassNetDataSource(filepath.Join(s.cfg.Collectors.RootPath, sys.ClassNetDir), dto.CollectorNameNetDev),
		),
		// /sys/class/net
		collectors.NewNetDevStatusCollector(s.cfg.Collectors,
			sys.NewClassNetDataSource(filepath.Join(s.cfg.Collectors.RootPath, sys.ClassNetDir), dto.CollectorNameNetDevStatus)),
		// iwconfig
		collectors.NewWirelessCollector(
			s.cfg.Collectors,
			run.NewIwconfigDataSource(run.NewCmdRunner(dto.CollectorNameWireless)),
			sys.NewClassNetDataSource(filepath.Join(s.cfg.Collectors.RootPath, sys.ClassNetDir), dto.CollectorNameWireless),
		),
		// /proc/net/arp
		collectors.NewNetARPCollector(s.cfg.Collectors,
			proc.NewNetARPDataSource(filepath.Join(s.cfg.Collectors.RootPath, proc.ProcPath, proc.NetArpPath), dto.CollectorNameNetARP)),
		// /proc/net/netstat
		collectors.NewNetStatCollector(s.cfg.Collectors,
			proc.NewNetStatDataSource(filepath.Join(s.cfg.Collectors.RootPath, proc.ProcPath, proc.NetStatPath), dto.CollectorNameNetStat)),
		// /proc/net/snmp
		collectors.NewNetSNMPCollector(s.cfg.Collectors,
			proc.NewNetStatDataSource(filepath.Join(s.cfg.Collectors.RootPath, proc.ProcPath, proc.NetSNMPPath), dto.CollectorNameNetSNMP)),
		// /proc/net/snmp6
		collectors.NewNetSNMP6Collector(s.cfg.Collectors,
			proc.NewNetSNMP6DataSource(filepath.Join(s.cfg.Collectors.RootPath, proc.ProcPath, proc.NetSNMP6Path), dto.CollectorNameNetSNMP6)),
		// mdStat
		collectors.NewMDStatCollector(s.cfg.Collectors,
			mdstat.NewMDStatDataSource(filepath.Join(s.cfg.Collectors.RootPath, mdstat.MDStatPath), dto.CollectorNameMDStat)),
		// CMD collector
		collectors.NewCmdCollector(s.cfg.Collectors,
			run.NewCmdRunner(dto.CollectorNameCMD)),
		// nginx
		collectors.NewNginxStubStatusCollector(s.cfg.Collectors,
			nginx.NewStubStatusClient(
				s.cfg.Collectors.NginxStubStatus,
				&http.Client{
					Timeout: time.Second * time.Duration(s.cfg.Collectors.NginxStubStatus.Timeout),
				},
				dto.CollectorNameNginx,
			),
		),
		// entropy
		collectors.NewEntropyCollector(s.cfg.Collectors,
			proc.NewEntropyDataSource(filepath.Join(s.cfg.Collectors.RootPath, proc.ProcPath, proc.EntropyPath), dto.CollectorNameEntropy),
		),
	}
}
