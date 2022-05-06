package mdStat

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/xray-team/xray-agent-linux/dto"
	"github.com/xray-team/xray-agent-linux/logger"
	"github.com/xray-team/xray-agent-linux/reader"
)

func (ds *mdStatDataSource) parseDevs(path string) (map[string]dto.DevStats, error) {
	var out = make(map[string]dto.DevStats)

	dirs, err := reader.ReadDir(path, ds.logPrefix)
	if err != nil {
		return out, fmt.Errorf("cannot read dir %s. err: %s", path, err)
	}

	re := regexp.MustCompile(`^dev-[a-zA-Z]{2,3}[0-9]{0,3}$`)

	for _, d := range dirs {
		if re.MatchString(d.Name()) {
			devStat, err := ds.parseDev(filepath.Join(path, d.Name()))
			if err != nil {
				logger.Log.Debug.Printf(logger.MessageError, dto.CollectorNameMDStat, fmt.Sprintf("cannot parse dev %s: %s", d.Name(), err.Error()))

				continue
			}

			out[strings.TrimPrefix(d.Name(), "dev-")] = devStat
		}
	}

	return out, nil
}

func (ds *mdStatDataSource) parseDev(path string) (dto.DevStats, error) {
	var (
		out dto.DevStats
		err error
	)

	// Slot
	// "none" if device is spare ...
	out.Slot, err = reader.ReadStringFile(filepath.Join(path, SlotFile), ds.logPrefix)
	if err != nil {
		return dto.DevStats{}, fmt.Errorf("cannot read mdstat file %s. %s", SlotFile, err)
	}

	// Errors
	out.Errors, err = reader.ReadInt64File(filepath.Join(path, ErrorsFile), ds.logPrefix)
	if err != nil {
		return dto.DevStats{}, fmt.Errorf("cannot read mdstat file %s. %s", ErrorsFile, err)
	}

	// State
	out.State, err = reader.ReadStringFile(filepath.Join(path, StateFile), ds.logPrefix)
	if err != nil {
		return dto.DevStats{}, fmt.Errorf("cannot read mdstat file %s. %s", StateFile, err)
	}

	return out, nil
}
