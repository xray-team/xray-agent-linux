package service

import (
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"xray-agent-linux/conf"
	"xray-agent-linux/dto"
	"xray-agent-linux/graphite"
	"xray-agent-linux/logger"
	"xray-agent-linux/stats"
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

func NewAgent(f conf.Flags) (*Agent, error) {
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
			logger.Log.Printf("Stopping module %s", m.Title())
			m.Stop()
		}
		logger.Log.Printf("Stopped all modules")
	}()

	for _, m := range modules {
		logger.Log.Printf("Starting module %s", m.Title())
		go m.Start()
	}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP)

	for {
		select {
		case sig := <-signalChan:
			switch sig {
			case syscall.SIGTERM, syscall.SIGINT:
				log.Printf("Got SIGTERM/SIGINT, exiting.")
				return
			}
		}
	}
}

func (a *Agent) init(f conf.Flags) error {
	a.Lock()
	defer a.Unlock()

	cfg, err := conf.GetConfiguration(&f)
	if err != nil {
		logger.LogValidationError(logger.ConfigValidationPrefix, err)

		return err
	}

	a.cfg = cfg.Agent
	telemetryChan := make(chan *dto.Telemetry)

	statsParser := stats.New(cfg, telemetryChan)

	// TODO TBD
	cfg.TSDB.Graphite.DryRun = *cfg.Agent.Flags.DryRun

	gr, err := graphite.New(cfg.TSDB.Graphite, telemetryChan)
	if err != nil {
		logger.Log.Printf("cannot init graphite. err %s", err)

		return err
	}

	a.statGetter = statsParser
	a.statSender = gr

	return nil
}
