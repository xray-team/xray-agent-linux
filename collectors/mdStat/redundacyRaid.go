package mdStat

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/xray-team/xray-agent-linux/dto"
	"github.com/xray-team/xray-agent-linux/reader"
)

func (ds *mdStatDataSource) parseStatsRaidWithRedundancy(path string) (*dto.StatRaidWithRedundancy, error) {
	var (
		out dto.StatRaidWithRedundancy
		err error
	)

	out.SyncAction, err = reader.ReadStringFile(filepath.Join(path, SyncActionFile), ds.logPrefix)
	if err != nil {
		return nil, fmt.Errorf("cannot parse mdstat file %s. %s", SyncActionFile, err)
	}

	out.NumDegraded, err = reader.ReadInt64File(filepath.Join(path, DegradedFile), ds.logPrefix)
	if err != nil {
		return nil, fmt.Errorf("cannot parse mdstat file %s. %s", DegradedFile, err)
	}

	out.MismatchCnt, err = reader.ReadInt64File(filepath.Join(path, MismatchCntFile), ds.logPrefix)
	if err != nil {
		return nil, fmt.Errorf("cannot parse mdstat file %s. %s", MismatchCntFile, err)
	}

	// May content word "none"
	out.SyncSpeed, _, err = reader.ReadVarFile(filepath.Join(path, SyncSpeedFile), ds.logPrefix)
	if err != nil {
		return nil, fmt.Errorf("cannot parse mdstat file %s. %s", SyncSpeedFile, err)
	}

	// ToDo: add "delayed" handling
	// May content words "none", "delayed"
	syncSectors, err := ds.parseSyncSectors(path)
	if err != nil {
		return nil, err
	}

	out.SyncCompletedSectors = syncSectors.CompletedSectors
	out.NumSectors = syncSectors.Sectors

	return &out, nil
}

type SyncCompleted struct {
	CompletedSectors int64
	Sectors          int64
	isNone           bool
	isDelayed        bool
}

func (ds *mdStatDataSource) parseSyncSectors(path string) (SyncCompleted, error) {
	var (
		out  SyncCompleted
		data string
		err  error
	)

	data, err = reader.ReadStringFile(filepath.Join(path, SyncCompletedFile), ds.logPrefix)
	if err != nil {
		return SyncCompleted{}, fmt.Errorf("cannot read mdstat file %s. %s", SyncCompletedFile, err)
	}

	if data == "none" {
		return SyncCompleted{isNone: true}, nil
	}

	if data == "delayed" {
		return SyncCompleted{isDelayed: true}, nil
	}

	fields := strings.Split(data, "/")
	if len(fields) != 2 {
		return SyncCompleted{}, fmt.Errorf("wrong mdstat file %s, value: %s. %s", SyncCompletedFile, data, err)
	}

	out.CompletedSectors, err = strconv.ParseInt(strings.TrimSpace(fields[0]), 10, 64)
	if err != nil {
		return SyncCompleted{}, fmt.Errorf("cannot parse syncCompletedSectors field in file %s. %s", SyncCompletedFile, err)
	}

	out.Sectors, err = strconv.ParseInt(strings.TrimSpace(fields[1]), 10, 64)
	if err != nil {
		return SyncCompleted{}, fmt.Errorf("cannot parse Sectors field in file %s. %s", SyncCompletedFile, err)
	}

	return out, nil
}
