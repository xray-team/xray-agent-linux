package main

import (
	"flag"

	"github.com/xray-team/xray-agent-linux/conf"
	"github.com/xray-team/xray-agent-linux/logger"
	"github.com/xray-team/xray-agent-linux/service"
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

	// Set logger params
	err = logger.SetLogger(cfg.Agent.LogOut, cfg.Agent.LogLevel)
	if err != nil {
		logger.Log.Error.Printf("can't set logger: %s", err)

		return
	}

	agent, err := service.NewAgent(*cfg)
	if err != nil {
		return
	}

	agent.Start()
}
