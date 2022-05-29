package logger

import "log"

const (
	StdoutFlags = log.LstdFlags
	SyslogFlags = 0 // disable time in log for syslog
)

// Error tags
const (
	TagAgent  = "Agent"
	TagConfig = "Config"
)

// Error messages
const (
	Message                   = "[%s] %s"
	MessageError              = "[%s] Error: %s"
	MessageInitCollector      = "[%s] Init collector"
	MessageInitCollectorError = "[%s] Collector init params error"
	MessageUnknownCollector   = "[%s] Unknown collector '%s'"
	MessageCollectError       = "[%s] Collect error: %s"
	MessageSetLogParamsError  = "[%s] Can't set logger params: %s"

	MessageReadFile           = "[%s] Reading file %s"
	MessageReadFileError      = "[%s] Error while trying to read file %s : %s"
	MessageReadFileFieldError = "[%s] Error while trying to read file %s, field %s : %s"
	MessageReadDir            = "[%s] Reading dir %s"
	MessageReadDirError       = "[%s] Error while trying to read dir %s"
	MessageIsExist            = "[%s] Checking if file or directory exists %s"
	MessageCmdRun             = "[%s] Exec command: '%s'"
	MessageCmdRunError        = "[%s] Exec error: %s"
	MessageHttpRequest        = "[%s] Doing request %s"
)
