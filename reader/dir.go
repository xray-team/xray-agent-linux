package reader

import (
	"os"

	"github.com/xray-team/xray-agent-linux/logger"
)

func ReadDir(path, logPrefix string) ([]os.DirEntry, error) {
	// logger
	logger.Log.Debug.Printf(logger.MessageReadDir, logPrefix, path)

	f, err := os.ReadDir(path)
	if err != nil {
		// logger
		logger.Log.Debug.Printf(logger.MessageReadDirError, logPrefix, path)

		return nil, err
	}

	return f, nil
}

func IsExist(path, logPrefix string) bool {
	// logger
	logger.Log.Debug.Printf(logger.MessageIsExist, logPrefix, path)

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}

	return true
}
