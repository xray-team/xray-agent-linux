package reader

import (
	"io/ioutil"
	"os"

	"xray-agent-linux/logger"
)

func ReadDir(path, logPrefix string) ([]os.FileInfo, error) {
	// logger
	logger.LogReadDir(logPrefix, path)

	f, err := ioutil.ReadDir(path)
	if err != nil {
		// logger
		logger.LogReadDirError(logPrefix, path)

		return nil, err
	}

	return f, nil
}

func IsExist(path, logPrefix string) bool {
	// logger
	logger.LogIsExist(logPrefix, path)

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}

	return true
}
