package cmd

import (
	"context"
	"fmt"
	"os/exec"
	"strings"
	"time"

	"github.com/xray-team/xray-agent-linux/dto"
	"github.com/xray-team/xray-agent-linux/logger"
)

type DataSource interface {
	RunPipeLine(pipeLine []*exec.Cmd) (string, string, error)
}

type Collector struct {
	Config     *Config
	DataSource DataSource
}

// NewCollector returns a new collector object.
func NewCollector(config *Config, dataSource DataSource) dto.Collector {
	if config == nil || dataSource == nil {
		logger.Log.Error.Printf(logger.MessageInitCollectorError, CollectorName)

		return nil
	}

	// exit if collector disabled
	if !config.Enabled {
		return nil
	}

	return &Collector{
		Config:     config,
		DataSource: dataSource,
	}
}

// GetName returns the collector's name.
func (c *Collector) GetName() string {
	return CollectorName
}

// Collect collects and returns metrics.
func (c *Collector) Collect() ([]dto.Metric, error) {
	metrics := make([]dto.Metric, 0, len(c.Config.Metrics))

	for i := range c.Config.Metrics {
		err := c.processPipeLine(&c.Config.Metrics[i], c.Config.Timeout, &metrics)
		if err != nil {
			logger.Log.Error.Printf(logger.MessageCmdRunError, CollectorName, err.Error())
		}
	}

	return metrics, nil
}

func (c *Collector) processPipeLine(cfg *MetricConfig, timeout int, out *[]dto.Metric) error {
	// Create a new context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()

	pipeLine := make([]*exec.Cmd, 0, len(cfg.PipeLine))

	for _, cmd := range cfg.PipeLine {
		if len(cmd) == 1 {
			pipeLine = append(pipeLine, exec.CommandContext(ctx, cmd[0]))
		}

		if len(cmd) >= 2 {
			pipeLine = append(pipeLine, exec.CommandContext(ctx, cmd[0], cmd[1:]...))
		}
	}

	stdout, stderr, err := c.DataSource.RunPipeLine(pipeLine)
	if err != nil {
		return err
	}

	// Timeout
	if ctx.Err() == context.DeadlineExceeded {
		return fmt.Errorf("command timed out")
	}

	// Check stderr
	if stderr != "" {
		return fmt.Errorf("stderr is not empty: '%s'", stderr)
	}

	values := strings.Split(strings.TrimSpace(stdout), cfg.Delimiter)

	if len(values) != len(cfg.Names) {
		return fmt.Errorf("metric count mismatch: config=%v, output=%v", len(cfg.Names), len(values))
	}

	for i, name := range cfg.Names {
		// skip ignored values
		if name == "-" {
			continue
		}

		*out = append(*out, dto.Metric{
			Name: cfg.Names[i],
			Attributes: append([]dto.MetricAttribute{
				{
					Name:  dto.ResourceAttr,
					Value: ResourceName,
				},
			}, cfg.Attributes...),
			Value: values[i],
		})
	}

	return nil
}
