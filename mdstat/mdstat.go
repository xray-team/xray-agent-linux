package mdstat

import (
	"fmt"
	"path/filepath"
	"regexp"

	"xray-agent-linux/dto"
	"xray-agent-linux/logger"
	"xray-agent-linux/reader"
)

type mdStatDataSource struct {
	path      string
	logPrefix string
}

func NewMDStatDataSource(path, logPrefix string) *mdStatDataSource {
	if path == "" {
		return nil
	}

	return &mdStatDataSource{
		path:      path,
		logPrefix: logPrefix,
	}
}

func (ds *mdStatDataSource) GetData() (*dto.MDStats, error) {
	// ls
	dirs, err := reader.ReadDir(ds.path, ds.logPrefix)
	if err != nil {
		return nil, fmt.Errorf("cannot read mdstat dir %s. err: %s", ds.path, err)
	}

	re := regexp.MustCompile(`^md[0-9]+$`)
	mdStats := make(map[string]dto.MDStat)

	for _, d := range dirs {
		if re.MatchString(d.Name()) {
			mdPath := filepath.Join(ds.path, d.Name())

			mdStat, err := ds.parseMDStat(mdPath)
			if err != nil {
				logger.Log.Printf("cannot parse md %s. err: %s\n", d.Name(), err)

				continue
			}

			mdStats[d.Name()] = mdStat
		}
	}

	if len(mdStats) == 0 {
		return nil, fmt.Errorf("there is no mdstat")
	}

	return &dto.MDStats{Stats: mdStats}, nil
}

func (ds *mdStatDataSource) parseMDStat(path string) (dto.MDStat, error) {
	var (
		out dto.MDStat
		err error
	)

	path = filepath.Join(path, MDSubFolder)

	out.Level, err = reader.ReadStringFile(filepath.Join(path, LevelFile), ds.logPrefix)
	if err != nil {
		return dto.MDStat{}, fmt.Errorf("cannot read mdstat file %s. %s", LevelFile, err)
	}

	out.NumDisks, err = reader.ReadInt64File(filepath.Join(path, RaidDisksFile), ds.logPrefix)
	if err != nil {
		return dto.MDStat{}, fmt.Errorf("cannot read mdstat file %s. %s", RaidDisksFile, err)
	}

	out.ArrayState, err = reader.ReadStringFile(filepath.Join(path, ArrayStateFile), ds.logPrefix)
	if err != nil {
		return dto.MDStat{}, fmt.Errorf("cannot read mdstat file %s. %s", ArrayStateFile, err)
	}

	// /sys/block/mdN/md/array_size
	// The word "default" means the effective size of the array to be whatever size is actually available based on level, chunk_size, and component_siz
	// ToDo: Calculate size
	out.ArraySizeKBytes, _, err = reader.ReadVarFile(filepath.Join(path, ArraySizeFile), ds.logPrefix)
	if err != nil {
		return dto.MDStat{}, fmt.Errorf("cannot read mdstat file %s. %s", ArraySizeFile, err)
	}

	out.ComponentSizeKBytes, err = reader.ReadInt64File(filepath.Join(path, ComponentSizeFile), ds.logPrefix)
	if err != nil {
		return dto.MDStat{}, fmt.Errorf("cannot read mdstat file %s. %s", ComponentSizeFile, err)
	}

	out.DevStats, err = ds.parseDevs(path)
	if err != nil {
		return dto.MDStat{}, fmt.Errorf("cannot read dev dir %s. %s", path, err)
	}

	out.StatRaidWithRedundancy, err = ds.parseStatsRaidWithRedundancy(path)
	if err != nil {
		logger.Log.Printf("cannot read dir %s. err: %s\n", path, err)
	}

	return out, nil
}
