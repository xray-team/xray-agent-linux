package graphite

import (
	"fmt"
	"strings"
	"time"

	"xray-agent-linux/conf"
	"xray-agent-linux/dto"
	"xray-agent-linux/logger"

	"github.com/crazygreenpenguin/graphite"
)

type Graphite struct {
	cfg           *conf.GraphiteConf
	telemetryChan <-chan *dto.Telemetry
	stopChan      chan bool
}

func New(cfg *conf.GraphiteConf, telemetryChan <-chan *dto.Telemetry) (*Graphite, error) {
	return &Graphite{cfg: cfg, telemetryChan: telemetryChan, stopChan: make(chan bool)}, nil
}

func (g *Graphite) Start() {
	for t := range g.telemetryChan {
		err := g.sendMetrics(t, g.cfg.DryRun)
		if err != nil {
			logger.Log.Printf("send metrics error %s", err)
		}
	}

	g.stopChan <- true
}

func (g *Graphite) Stop() {
	<-g.stopChan
	close(g.stopChan)
}

func (g *Graphite) Title() string {
	return "graphite"
}

func (g *Graphite) sendMetrics(telemetry *dto.Telemetry, dryRun bool) error {
	var gm []graphite.Metric

	switch g.cfg.Mode {
	case dto.GraphiteModeTree:
		gm = genGraphiteTreeMetrics(*telemetry)
	case dto.GraphiteModeTags:
		gm = genGraphiteTagsMetrics(*telemetry)
	default:
		return fmt.Errorf("bad graphite mode")
	}

	// debug
	if dryRun {
		logger.Log.Printf("metrics %+v", gm)

		return nil
	}

	return g.sendMetricsToGraphite(telemetry.HostInfo, gm)
}

func (g *Graphite) sendMetricsToGraphite(hostInfo *dto.HostInfo, metrics []graphite.Metric) error {
	if metrics == nil {
		logger.Log.Printf("nothing to send to graphite\n")

		return nil
	}

	errors := make([]string, 0)

	for _, serverConf := range g.cfg.Servers {
		// Initializing graphite client
		graphiteClient, err := graphite.NewGraphiteTCP(convertGraphiteConf(hostInfo, &serverConf, g.cfg.Mode))
		if err != nil {
			errors = append(errors, fmt.Sprintf("error while sending metric to server %s: %s", serverConf.Address, err.Error()))

			continue
		}

		// Sending metrics to server
		err = graphiteClient.SendMetrics(&metrics)
		if err != nil {
			errors = append(errors, fmt.Sprintf("error while sending metric to server %s: %s", serverConf.Address, err.Error()))
		}
	}

	if len(errors) != 0 {
		return fmt.Errorf(strings.Join(errors, "; "))
	}

	return nil
}

func convertGraphiteConf(hostInfo *dto.HostInfo, cfg *conf.GraphiteServerConf, mode string) *graphite.Config {
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
	if mode == dto.GraphiteModeTree {
		for _, attr := range hostInfo.Attributes {
			out.Prefix = fmt.Sprintf("%s.%s", out.Prefix, attr.Value)
		}
	}

	return &out
}
