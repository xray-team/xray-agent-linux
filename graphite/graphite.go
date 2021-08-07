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

		// debug
		if dryRun {
			logger.Log.Printf("metrics %+v", gm)

			continue
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
