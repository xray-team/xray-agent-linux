package stats

import (
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
	"github.com/xray-team/xray-agent-linux/collectors/vmStat"
	"github.com/xray-team/xray-agent-linux/collectors/wireless"

	"github.com/xray-team/xray-agent-linux/conf"
	"github.com/xray-team/xray-agent-linux/dto"
	"github.com/xray-team/xray-agent-linux/logger"
)

type Stat struct {
	cfg           *conf.Config
	telemetryChan chan<- *dto.Telemetry
	stopChan      chan bool
	reg           map[string]func([]byte) dto.Collector
	collectors    []dto.Collector
}

func New(cfg *conf.Config, telemetryChan chan<- *dto.Telemetry) *Stat {
	return &Stat{
		cfg:           cfg,
		telemetryChan: telemetryChan,
		stopChan:      make(chan bool),
		reg:           make(map[string]func([]byte) dto.Collector),
	}
}

func (s *Stat) Title() string {
	return "stat"
}

func (s *Stat) RegisterCollector(name string, createFunc func([]byte) dto.Collector) {
	s.reg[name] = createFunc
}

func (s *Stat) RegisterCollectors() {
	s.RegisterCollector(cmd.CollectorName, cmd.CreateCollector)
	s.RegisterCollector(cpuInfo.CollectorName, cpuInfo.CreateCollector)
	s.RegisterCollector(diskSpace.CollectorName, diskSpace.CreateCollector)
	s.RegisterCollector(diskStat.CollectorName, diskStat.CreateCollector)
	s.RegisterCollector(entropy.CollectorName, entropy.CreateCollector)
	s.RegisterCollector(loadAvg.CollectorName, loadAvg.CreateCollector)
	s.RegisterCollector(mdStat.CollectorName, mdStat.CreateCollector)
	s.RegisterCollector(memoryInfo.CollectorName, memoryInfo.CreateCollector)
	s.RegisterCollector(netARP.CollectorName, netARP.CreateCollector)
	s.RegisterCollector(netDev.CollectorName, netDev.CreateCollector)
	s.RegisterCollector(netDevStatus.CollectorName, netDevStatus.CreateCollector)
	s.RegisterCollector(netSNMP.CollectorName, netSNMP.CreateCollector)
	s.RegisterCollector(netSNMP6.CollectorName, netSNMP6.CreateCollector)
	s.RegisterCollector(netStat.CollectorName, netStat.CreateCollector)
	s.RegisterCollector(nginx.CollectorName, nginx.CreateCollector)
	s.RegisterCollector(ps.CollectorName, ps.CreateCollector)
	s.RegisterCollector(psStat.CollectorName, psStat.CreateCollector)
	s.RegisterCollector(stat.CollectorName, stat.CreateCollector)
	s.RegisterCollector(uptime.CollectorName, uptime.CreateCollector)
	s.RegisterCollector(vmStat.CollectorName, vmStat.CreateCollector)
	s.RegisterCollector(wireless.CollectorName, wireless.CreateCollector)
}

func (s *Stat) initCollectors() {
	for name, collectorConfig := range s.cfg.Collectors {
		createFunc, ok := s.reg[name]

		if ok {
			logger.Log.Info.Printf(logger.MessageInitCollector, name)
			s.collectors = append(s.collectors, createFunc(collectorConfig))
		} else {
			logger.Log.Info.Printf(logger.MessageUnknownCollector, logger.TagAgent, name)
		}
	}
}

func (s *Stat) Start() {
	s.RegisterCollectors()
	s.initCollectors()

	ticker := time.NewTicker(time.Duration(s.cfg.Agent.GetStatIntervalSec) * time.Second)
	defer func() {
		ticker.Stop()
		close(s.telemetryChan)
		close(s.stopChan)
	}()

	// start first time before ticker
	s.getStat()

	for {
		select {
		case <-ticker.C:
			s.getStat()
		case <-s.stopChan:
			return
		}
	}
}

func (s *Stat) Stop() {
	s.stopChan <- true
}

func (s *Stat) DryRun() {
	s.RegisterCollectors()
	s.initCollectors()

	defer func() {
		close(s.telemetryChan)
		close(s.stopChan)
	}()

	s.getStat()
}

func (s *Stat) getStat() {
	logger.Log.Info.Printf(logger.Message, logger.TagAgent, "Collect")

	var (
		metrics    []dto.Metric
		numMetrics int
		startTime  = time.Now()
	)

	hostInfo, err := getHostInfo(s.cfg.Agent)
	if err != nil {
		logger.Log.Error.Printf(logger.MessageCollectError, logger.TagAgent, err.Error())

		return
	}

	for _, collector := range s.collectors {
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
	if s.cfg.Agent.EnableSelfMetrics {
		agentSummaryMetrics := s.agentSummaryToMetrics(dto.AgentSummary{
			Duration:      time.Since(startTime),
			MetricsNumber: len(metrics),
		})

		metrics = append(metrics, agentSummaryMetrics...)
	}

	s.telemetryChan <- &dto.Telemetry{
		HostInfo: hostInfo,
		Metrics:  metrics,
	}
}

func (s *Stat) Collect(collector dto.Collector) ([]dto.Metric, error) {
	var (
		metrics            []dto.Metric
		err                error
		collectorStartTime = time.Now()
		summary            = dto.CollectorSummary{
			CollectorName: collector.GetName(),
			Status:        1, // 1 - success
		}
	)

	func() {
		defer func() {
			if s.cfg.Agent.EnableSelfMetrics {
				summary.Duration = time.Since(collectorStartTime)
				summary.MetricsNumber = len(metrics)
				// append self metrics (collector scope)
				metrics = append(metrics, s.collectorSummaryToMetrics(summary)...)
			}
		}()

		metrics, err = collector.Collect()
		if err != nil {
			summary.Status = 2 // 2 - error
		}
	}()

	return metrics, err
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
