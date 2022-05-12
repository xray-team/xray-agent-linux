package service

import (
	"fmt"
	"sync"

	"github.com/xray-team/xray-agent-linux/conf"
	"github.com/xray-team/xray-agent-linux/dto"
	"github.com/xray-team/xray-agent-linux/graphite"
	"github.com/xray-team/xray-agent-linux/logger"
	"github.com/xray-team/xray-agent-linux/stats"
)

type Agent struct {
	sync.Mutex
	cfg        *conf.AgentConf
	statGetter Module
	statSender Module
}

func NewAgent(f conf.Config) (*Agent, error) {
	agent := &Agent{}
	err := agent.init(f)
	if err != nil {
		return nil, err
	}

	return agent, nil
}

func (a *Agent) Start() {
	RunModules(a.statGetter, a.statSender)
}

func (a *Agent) init(cfg conf.Config) error {
	a.Lock()
	defer a.Unlock()

	a.cfg = cfg.Agent
	telemetryChan := make(chan *dto.Telemetry)

	statsParser := stats.New(&cfg, telemetryChan)

	// TODO TBD
	cfg.TSDB.Graphite.DryRun = *cfg.Agent.Flags.DryRun

	gr, err := graphite.New(cfg.TSDB.Graphite, telemetryChan)
	if err != nil {
		logger.Log.Error.Printf(logger.MessageError, logger.TagAgent, fmt.Sprintf("cannot init graphite: %s", err.Error()))

		return err
	}

	a.statGetter = statsParser
	a.statSender = gr

	return nil
}
