package service

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/xray-team/xray-agent-linux/conf"
	"github.com/xray-team/xray-agent-linux/dto"
	"github.com/xray-team/xray-agent-linux/graphite"
	"github.com/xray-team/xray-agent-linux/logger"
	"github.com/xray-team/xray-agent-linux/stats"
)

type Module interface {
	Start()
	Stop()
	Title() string
}

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

// RunModules runs each of the modules in a separate goroutine.
func RunModules(modules ...Module) {
	defer func() {
		for _, m := range modules {
			logger.Log.Info.Printf(logger.Message, logger.TagAgent, fmt.Sprintf("Stopping module %s", m.Title()))

			m.Stop()
		}
		logger.Log.Info.Printf(logger.Message, logger.TagAgent, "Stopped all modules")
	}()

	for _, m := range modules {
		logger.Log.Info.Printf(logger.Message, logger.TagAgent, fmt.Sprintf("Starting module %s", m.Title()))
		go m.Start()
	}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGTERM, syscall.SIGINT)

	for {
		select {
		case sig := <-signalChan:
			switch sig {
			case syscall.SIGTERM, syscall.SIGINT:
				logger.Log.Info.Printf(logger.Message, logger.TagAgent, "Got SIGTERM/SIGINT, exiting.")

				return
			}
		}
	}
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
