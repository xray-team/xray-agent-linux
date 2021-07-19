package run

import (
	"bytes"
	"fmt"
	"os/exec"

	"xray-agent-linux/logger"
)

type cmdRunner struct {
	logPrefix string
}

func NewCmdRunner(logPrefix string) *cmdRunner {
	return &cmdRunner{
		logPrefix: logPrefix,
	}
}

func (r *cmdRunner) Run(cmd *exec.Cmd) (string, string, error) {
	var (
		stdout bytes.Buffer
		stderr bytes.Buffer
	)

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	// logger
	logger.LogCmdRun(r.logPrefix, cmd.String())

	err := cmd.Run()
	if err != nil {
		return stdout.String(), stderr.String(), fmt.Errorf("error while trying to execute command: %s: %w: %s", cmd.Args, err, cmd.Stderr)
	}

	return stdout.String(), stderr.String(), nil
}

func (r *cmdRunner) RunPipeLine(pipeLine []*exec.Cmd) (string, string, error) {
	var (
		stdout bytes.Buffer
		stderr bytes.Buffer
	)

	for i, _ := range pipeLine[:len(pipeLine)-1] {
		var err error

		//  Connect each command's standard output to the standard input of the next command
		pipeLine[i+1].Stdin, err = pipeLine[i].StdoutPipe()
		if err != nil {
			return "", "", err
		}

		// Connect command's stderr to a buffer
		pipeLine[i].Stderr = &stderr
	}

	// logger
	logger.LogCmdRun(r.logPrefix, "pipeline")

	// Last command
	pipeLine[len(pipeLine)-1].Stdout = &stdout
	pipeLine[len(pipeLine)-1].Stderr = &stderr

	for i, _ := range pipeLine {
		if err := pipeLine[i].Start(); err != nil {
			return stdout.String(), stderr.String(), err
		}
	}

	for i, _ := range pipeLine {
		if err := pipeLine[i].Wait(); err != nil {
			return stdout.String(), stderr.String(), err
		}
	}

	return stdout.String(), stderr.String(), nil
}
