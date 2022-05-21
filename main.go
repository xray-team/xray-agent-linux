package main

import (
	"flag"
	"github.com/xray-team/xray-agent-linux/conf"
	"github.com/xray-team/xray-agent-linux/dto"
	"github.com/xray-team/xray-agent-linux/graphite"
	"github.com/xray-team/xray-agent-linux/logger"
	"github.com/xray-team/xray-agent-linux/service"
	"github.com/xray-team/xray-agent-linux/stats"
)

func main() {
	// init default logger
	logger.Init()

	// Parse flags
	var f conf.Flags
	f.ConfigFilePath = flag.String("config", "./config.json", "path to config file")
	f.DryRun = flag.Bool("dryrun", false, "test run")
	flag.Parse()

	// Read and parse config
	cfg, err := conf.GetConfiguration(&f)
	if err != nil {
		logger.LogValidationError(err)

		return
	}

	// update logger params
	if *f.DryRun {
		err = logger.SetLogger("stdout", "debug")
	} else {
		err = logger.SetLogger(cfg.Agent.LogOut, cfg.Agent.LogLevel)
	}

	if err != nil {
		logger.Log.Error.Printf(logger.MessageSetLogParamsError, logger.TagAgent, err.Error())

		return
	}

	// Start service
	telemetryChan := make(chan *dto.Telemetry)
	cfg.TSDB.Graphite.DryRun = *cfg.Agent.Flags.DryRun

	agent := service.New(stats.New(cfg, telemetryChan), graphite.New(cfg.TSDB.Graphite, telemetryChan))
	agent.Start()
}
