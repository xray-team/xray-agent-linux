package reader

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/xray-team/xray-agent-linux/logger"
)

func ReadFile(filePath, logPrefix string) ([]byte, error) {
	filePath = filepath.Clean(filePath)
	// logger
	logger.LogReadFile(logPrefix, filePath)

	// read file to memory
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		// logger
		logger.LogReadFileError(logPrefix, filePath, err)

		return nil, err
	}

	return data, nil
}

// ReadMultilineFile
func ReadMultilineFile(filePath, logPrefix string) ([]string, error) {
	data, err := ReadFile(filePath, logPrefix)
	if err != nil {
		return nil, err
	}

	// split by newlines
	lines := bytes.Split(data, []byte("\n"))

	// Delete empty lines and Convert []byte to string
	var linesStr []string

	for _, v := range lines {
		if len(v) != 0 {
			linesStr = append(linesStr, string(v))
		}
	}

	if len(linesStr) == 0 {
		err = fmt.Errorf("no lines are parsed")
		// logger
		logger.LogReadFileError(logPrefix, filePath, err)

		return nil, err
	}

	return linesStr, nil
}

// ReadStringFile - reads file as single string.
func ReadStringFile(filePath, logPrefix string) (string, error) {
	out, err := ReadFile(filePath, logPrefix)
	if err != nil {
		return "", err
	}

	return strings.TrimSuffix(string(out), "\n"), nil
}

// ReadInt64File - reads file as single int64 value.
func ReadInt64File(filePath, logPrefix string) (int64, error) {
	data, err := ReadFile(filePath, logPrefix)
	if err != nil {
		return 0, err
	}

	// convert ot int64
	i, err := strconv.ParseInt(string(bytes.TrimSpace(data)), 10, 64)
	if err != nil {
		return 0, fmt.Errorf("'%s': %w", filePath, err)
	}

	return i, nil
}

func ReadVarFile(filePath, logPrefix string) (int64, string, error) {
	// read file to memory
	data, err := ReadFile(filePath, logPrefix)
	if err != nil {
		return 0, "", err
	}

	data = bytes.TrimSuffix(data, []byte("\n"))

	i, err := strconv.ParseInt(string(bytes.TrimSpace(data)), 10, 64)
	if err != nil {
		return 0, string(data), nil
	}

	return i, string(data), nil
}
