package run

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/xray-team/xray-agent-linux/logger"
)

type CmdRunner struct {
	logPrefix string
}

func NewCmdRunner(logPrefix string) *CmdRunner {
	return &CmdRunner{
		logPrefix: logPrefix,
	}
}

func (r *CmdRunner) Run(cmd *exec.Cmd) (string, string, error) {
	var (
		stdout bytes.Buffer
		stderr bytes.Buffer
	)

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	// logger
	logger.Log.Debug.Printf(logger.MessageCmdRun, r.logPrefix, cmd.String())

	err := cmd.Run()
	if err != nil {
		return stdout.String(), stderr.String(), fmt.Errorf("error while trying to execute command: %s: %w: %s", cmd.Args, err, cmd.Stderr)
	}

	return stdout.String(), stderr.String(), nil
}

func (r *CmdRunner) RunPipeLine(pipeLine []*exec.Cmd) (string, string, error) {
	var (
		stdout bytes.Buffer
		stderr bytes.Buffer
	)

	for i := range pipeLine[:len(pipeLine)-1] {
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
	logger.Log.Debug.Printf(logger.MessageCmdRun, r.logPrefix, "pipeline")

	// Last command
	pipeLine[len(pipeLine)-1].Stdout = &stdout
	pipeLine[len(pipeLine)-1].Stderr = &stderr

	for i := range pipeLine {
		if err := pipeLine[i].Start(); err != nil {
			return stdout.String(), stderr.String(), err
		}
	}

	for i := range pipeLine {
		if err := pipeLine[i].Wait(); err != nil {
			return stdout.String(), stderr.String(), err
		}
	}

	return stdout.String(), stderr.String(), nil
}
