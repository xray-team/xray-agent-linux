package proc

import (
	"regexp"
	"strconv"
	"strings"

	"xray-agent-linux/dto"
	"xray-agent-linux/logger"
	"xray-agent-linux/reader"
)

type mountsDataSource struct {
	filePath  string
	logPrefix string
}

func NewMountsDataSource(filePath, logPrefix string) *mountsDataSource {
	if filePath == "" {
		return nil
	}

	return &mountsDataSource{
		filePath:  filePath,
		logPrefix: logPrefix,
	}
}

func (ds *mountsDataSource) GetData() ([]dto.Mounts, error) {
	out := make([]dto.Mounts, 0)

	lines, err := reader.ReadMultilineFile(ds.filePath, ds.logPrefix)
	if err != nil {
		return out, err
	}

	for _, v := range lines {
		fields := strings.Fields(v)

		// skip incorrect lines
		if len(fields) != 6 {
			continue
		}

		// skip comments
		re := regexp.MustCompile("^#")
		if re.Match([]byte(fields[0])) {
			continue
		}

		var mount dto.Mounts

		mount.Dev = fields[0]
		mount.MountPoint = fields[1]
		mount.FileSystemType = fields[2]
		mount.MountOptions = fields[3]
		mount.Dump, err = strconv.ParseInt(fields[4], 10, 64)

		if err != nil {
			logger.LogReadFileFieldError(ds.logPrefix, ds.filePath, "Dump", err)

			continue
		}

		mount.Pass, err = strconv.ParseInt(fields[5], 10, 64)

		if err != nil {
			logger.LogReadFileFieldError(ds.logPrefix, ds.filePath, "Pass", err)

			continue
		}

		out = append(out, mount)
	}

	return out, err
}
