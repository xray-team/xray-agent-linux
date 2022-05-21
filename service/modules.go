package service

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/xray-team/xray-agent-linux/logger"
)

type Module interface {
	Start()
	Stop()
	DryRun()
	Title() string
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

func DryRunModules(modules ...Module) {
	for _, m := range modules {
		logger.Log.Info.Printf(logger.Message, logger.TagAgent, fmt.Sprintf("Starting module %s (DryRun mode)", m.Title()))
		m.DryRun()
	}
}
