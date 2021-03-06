package diskSpace

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/xray-team/xray-agent-linux/logger"
	"github.com/xray-team/xray-agent-linux/reader"
)

type mountsDataSource struct {
	filePath  string
	logPrefix string
}

// NewMountsDataSource returns a new DataSource.
func NewMountsDataSource(filePath, logPrefix string) *mountsDataSource {
	if filePath == "" {
		return nil
	}

	return &mountsDataSource{
		filePath:  filePath,
		logPrefix: logPrefix,
	}
}

func (ds *mountsDataSource) GetData() ([]Mounts, error) {
	out := make([]Mounts, 0)

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

		var mount Mounts

		mount.Dev = fields[0]
		mount.MountPoint = fields[1]
		mount.FileSystemType = fields[2]
		mount.MountOptions = fields[3]
		mount.Dump, err = strconv.ParseInt(fields[4], 10, 64)

		if err != nil {
			logger.Log.Debug.Printf(logger.MessageReadFileFieldError, ds.logPrefix, ds.filePath, "Dump", err.Error())

			continue
		}

		mount.Pass, err = strconv.ParseInt(fields[5], 10, 64)

		if err != nil {
			logger.Log.Debug.Printf(logger.MessageReadFileFieldError, ds.logPrefix, ds.filePath, "Pass", err.Error())

			continue
		}

		out = append(out, mount)
	}

	return out, err
}
