package graphite

import (
	"fmt"
	"strings"
	"time"

	"github.com/xray-team/xray-agent-linux/conf"
	"github.com/xray-team/xray-agent-linux/dto"
	"github.com/xray-team/xray-agent-linux/logger"

	"github.com/crazygreenpenguin/graphite"
)

type Graphite struct {
	cfg           *conf.GraphiteConf
	telemetryChan <-chan *dto.Telemetry
	stopChan      chan bool
}

func New(cfg *conf.GraphiteConf, telemetryChan <-chan *dto.Telemetry) *Graphite {
	return &Graphite{
		cfg:           cfg,
		telemetryChan: telemetryChan,
		stopChan:      make(chan bool),
	}
}

func (g *Graphite) Start() {
	for t := range g.telemetryChan {
		err := g.sendMetrics(t)
		if err != nil {
			logger.Log.Error.Printf(logger.Message, logger.TagAgent, fmt.Sprintf("send metrics error: %s", err.Error()))
		}
	}

	g.stopChan <- true
}

func (g *Graphite) Stop() {
	<-g.stopChan
	close(g.stopChan)
}

func (g *Graphite) DryRun() {
	for t := range g.telemetryChan {
		g.printMetrics(t)
	}
}

func (g *Graphite) Title() string {
	return "graphite"
}

func (g *Graphite) sendMetrics(telemetry *dto.Telemetry) error {
	var (
		gm     []graphite.Metric
		errors = make([]string, 0)
	)

	for _, serverConf := range g.cfg.Servers {
		switch serverConf.Mode {
		case dto.GraphiteModeTree:
			gm = genGraphiteTreeMetrics(*telemetry)
		case dto.GraphiteModeTags:
			gm = genGraphiteTagsMetrics(*telemetry)
		default:
			return fmt.Errorf("bad graphite mode")
		}

		// Initializing graphite client
		graphiteClient, err := graphite.NewGraphiteTCP(convertGraphiteConf(telemetry.HostInfo, &serverConf))
		if err != nil {
			errors = append(errors, fmt.Sprintf("error while sending metric to server %s: %s", serverConf.Address, err.Error()))

			continue
		}

		// Sending metrics to server
		err = graphiteClient.SendMetrics(&gm)
		if err != nil {
			errors = append(errors, fmt.Sprintf("error while sending metric to server %s: %s", serverConf.Address, err.Error()))
		}
	}

	if len(errors) != 0 {
		return fmt.Errorf(strings.Join(errors, "; "))
	}

	return nil
}

func (g *Graphite) printMetrics(telemetry *dto.Telemetry) {
	var gm []graphite.Metric

	for _, serverConf := range g.cfg.Servers {
		switch serverConf.Mode {
		case dto.GraphiteModeTree:
			gm = genGraphiteTreeMetrics(*telemetry)
		case dto.GraphiteModeTags:
			gm = genGraphiteTagsMetrics(*telemetry)
		}

		logger.Log.Info.Printf("metrics for server:%s %+v", serverConf.Address, gm)
	}
}

func convertGraphiteConf(hostInfo *dto.HostInfo, cfg *conf.GraphiteServerConf) *graphite.Config {
	out := graphite.Config{
		Address: cfg.Address,
		Timeout: time.Duration(cfg.Timeout) * time.Second,
	}

	// Protocol
	switch cfg.Protocol {
	case "tcp":
		out.Protocol = 0
	case "udp":
		out.Protocol = 1
	case "stdout":
		out.Protocol = 2
	default:
		out.Protocol = 0
	}

	// Prefix
	if cfg.Mode == dto.GraphiteModeTree {
		for _, attr := range hostInfo.Attributes {
			out.Prefix = fmt.Sprintf("%s.%s", out.Prefix, attr.Value)
		}
	}

	return &out
}
