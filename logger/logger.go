package logger

import (
	"github.com/go-playground/validator"
	"log"
	"os"
)

// Log is global logger variable.
// nolint: gochecknoglobals
var Log Logger

// Init logger.
func Init(prefix string) {
	Log = log.New(os.Stdout, prefix, log.LstdFlags|log.Lshortfile)
}

// Logger interface implements one method - Printf.
// You can use the stdlib logger *logger.Logger.
type Logger interface {
	Printf(format string, v ...interface{})
}

func LogReadFile(logPrefix, filePath string) {
	Log.Printf("[%s] Reading file %s", logPrefix, filePath)
}

func LogReadFileError(logPrefix, filePath string, err error) {
	Log.Printf("[%s] Error while trying to read file %s : %s", logPrefix, filePath, err.Error())
}

func LogReadFileFieldError(logPrefix, filePath, field string, err error) {
	Log.Printf("[%s] Error while trying to read file %s, field %s : %s", logPrefix, filePath, field, err.Error())
}

func LogReadDir(logPrefix, path string) {
	Log.Printf("[%s] Reading dir %s", logPrefix, path)
}

func LogReadDirError(logPrefix, path string) {
	Log.Printf("[%s] Error while trying to read dir %s", logPrefix, path)
}

func LogIsExist(logPrefix, path string) {
	Log.Printf("[%s] Checking if file or directory exists %s", logPrefix, path)
}

func LogWarning(logPrefix string, err error) {
	Log.Printf("[%s] %s", logPrefix, err.Error())
}

func LogValidationError(logPrefix string, err error) {
	switch err.(type) {
	case validator.ValidationErrors:
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			Log.Printf("[%s] invalid field: '%s', invalid value: '%v', tag: '%s', param: '%s'", logPrefix, e.Namespace(), e.Value(), e.Tag(), e.Param())
		}
	default:
		Log.Printf("[%s] %s", logPrefix, err.Error())
	}
}

func LogCmdRun(logPrefix, command string) {
	Log.Printf("[%s] Exec command: '%s'", logPrefix, command)
}

func LogHttpRequest(logPrefix, req string) {
	Log.Printf("[%s] Doing request %s", logPrefix, req)
}
