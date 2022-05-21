package main

import (
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
	flags := conf.ParseFlags()

	// Read and parse config
	config, err := conf.GetConfiguration(flags)
	if err != nil {
		logger.LogValidationError(err)

		return
	}

	// update logger params
	if *flags.DryRun {
		err = logger.SetLogger("stdout", "debug")
	} else {
		err = logger.SetLogger(config.Agent.LogOut, config.Agent.LogLevel)
	}

	if err != nil {
		logger.Log.Error.Printf(logger.MessageSetLogParamsError, logger.TagAgent, err.Error())

		return
	}

	// Start service
	if *flags.DryRun {
		telemetryChan := make(chan *dto.Telemetry, 1)
		agent := service.New(stats.New(config, telemetryChan), graphite.New(config.TSDB.Graphite, telemetryChan))
		agent.DryRun()
	} else {
		telemetryChan := make(chan *dto.Telemetry)
		agent := service.New(stats.New(config, telemetryChan), graphite.New(config.TSDB.Graphite, telemetryChan))
		agent.Start()
	}
}
