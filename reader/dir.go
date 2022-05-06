package reader

import (
	"io/ioutil"
	"os"

	"github.com/xray-team/xray-agent-linux/logger"
)

func ReadDir(path, logPrefix string) ([]os.FileInfo, error) {
	// logger
	logger.Log.Debug.Printf(logger.MessageReadDir, logPrefix, path)

	f, err := ioutil.ReadDir(path)
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
