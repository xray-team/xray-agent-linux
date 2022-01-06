package main

import (
	"flag"

	"github.com/xray-team/xray-agent-linux/conf"
	"github.com/xray-team/xray-agent-linux/logger"
	"github.com/xray-team/xray-agent-linux/service"
)

func main() {
	logger.Init("")

	// Parse flags
	var f conf.Flags
	f.ConfigFilePath = flag.String("config", "./config.json", "path to config file")
	f.DryRun = flag.Bool("dryrun", false, "test run")
	flag.Parse()

	agent, err := service.NewAgent(f)
	if err != nil {
		return
	}

	agent.Start()
}
