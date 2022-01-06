package sys

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/xray-team/xray-agent-linux/dto"
	"github.com/xray-team/xray-agent-linux/logger"
	"github.com/xray-team/xray-agent-linux/reader"
)

const (
	ClassBlockDir = "/sys/class/block/"

	BlockUeventFile = "uevent"
)

type classBlockDataSource struct {
	path      string
	logPrefix string
}

func NewClassBlockDataSource(path, logPrefix string) *classBlockDataSource {
	if path == "" {
		return nil
	}

	return &classBlockDataSource{
		path:      path,
		logPrefix: logPrefix,
	}
}

func (ds *classBlockDataSource) GetData() (map[string]dto.ClassBlock, error) {
	dirs, err := reader.ReadDir(ds.path, ds.logPrefix)
	if err != nil {
		return nil, fmt.Errorf("cannot read dir %s. err: %s", ds.path, err)
	}

	var out = make(map[string]dto.ClassBlock)

	for _, dir := range dirs {
		var blockDev dto.ClassBlock

		// Uevent
		blockDev.Uevent, err = ds.parseBlockUevent(filepath.Join(ds.path, dir.Name(), BlockUeventFile))
		if err != nil {
			return nil, fmt.Errorf("cannot parse uevent file: %s. err: %s", filepath.Join(ds.path, dir.Name(), BlockUeventFile), err)
		}

		out[dir.Name()] = blockDev
	}

	return out, nil
}

func (ds *classBlockDataSource) parseBlockUevent(filePath string) (*dto.BlockDevUevent, error) {
	lines, err := reader.ReadMultilineFile(filePath, ds.logPrefix)
	if err != nil {
		return nil, err
	}

	var out dto.BlockDevUevent

	for _, line := range lines {
		fields := strings.Split(line, "=")
		// skip incorrect lines
		if len(fields) < 2 {
			continue
		}

		switch fields[0] {
		case "MAJOR":
			out.Major, err = strconv.ParseInt(fields[1], 10, 64)
			if err != nil {
				logger.LogReadFileFieldError(ds.logPrefix, filePath, "MAJOR", err)
			}
		case "MINOR":
			out.Minor, err = strconv.ParseInt(fields[1], 10, 64)
			if err != nil {
				logger.LogReadFileFieldError(ds.logPrefix, filePath, "MAJOR", err)
			}
		case "DEVNAME":
			out.DevName = fields[1]
		case "DEVTYPE":
			out.DevType = fields[1]
		case "PARTN":
			out.PartNumber = fields[1]
		case "PARTNAME":
			out.PartName = fields[1]
		}
	}

	return &out, nil
}
