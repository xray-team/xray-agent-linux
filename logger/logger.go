package logger

import (
	"io"
	"io/ioutil"
	"log"
	"log/syslog"
	"os"
)

// Log are global logger variable.
// nolint: gochecknoglobals
var (
	Log *LogSeverity
)

type LogSeverity struct {
	Info  *log.Logger
	Error *log.Logger
	Debug *log.Logger
}

// Init logger.
func Init() {
	Log.SetDefault(os.Stdout, StdoutFlags)
}

func (logSeverity *LogSeverity) SetDefault(out io.Writer, flags int) {
	Log = &LogSeverity{
		Info:  log.New(out, "", flags),
		Error: log.New(out, "", flags),
		Debug: log.New(ioutil.Discard, "", flags),
	}
}

func (logSeverity *LogSeverity) SetDebug(out io.Writer, flags int) {
	Log = &LogSeverity{
		Info:  log.New(out, "", flags),
		Error: log.New(out, "", flags),
		Debug: log.New(out, "", flags),
	}
}

func SetLogger(out, level string) error {
	switch out {
	case "stdout":
		setStdoutLogger(level)
	case "syslog":
		err := setSyslogLogger(level)
		if err != nil {
			return err
		}
	}

	return nil
}

func setStdoutLogger(level string) {
	switch level {
	case "debug":
		Log.SetDebug(os.Stdout, StdoutFlags)
	default:
		Log.SetDefault(os.Stdout, StdoutFlags)
	}
}

func setSyslogLogger(level string) error {
	// connect to local instance of syslog
	sysLog, err := syslog.Dial("", "", syslog.LOG_DEBUG|syslog.LOG_DAEMON, "xray-agent")
	if err != nil {
		return err
	}

	switch level {
	case "debug":
		Log.SetDebug(sysLog, SyslogFlags)
	default:
		Log.SetDefault(sysLog, SyslogFlags)
	}

	return nil
}
