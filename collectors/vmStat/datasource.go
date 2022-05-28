package vmStat

import (
	"strconv"
	"strings"

	"github.com/xray-team/xray-agent-linux/logger"
	"github.com/xray-team/xray-agent-linux/reader"
)

type vmStatDataSource struct {
	filePath  string
	logPrefix string
}

// NewDataSource returns a new DataSource.
func NewDataSource(filePath, logPrefix string) *vmStatDataSource {
	if filePath == "" {
		return nil
	}

	return &vmStatDataSource{
		filePath:  filePath,
		logPrefix: logPrefix,
	}
}

func (ds *vmStatDataSource) GetData() (*VMStat, error) {
	var (
		data VMStat
		err  error
	)

	// Implement GetData logic here
	lines, err := reader.ReadMultilineFile(ds.filePath, ds.logPrefix)
	if err != nil {
		return nil, err
	}

	for _, v := range lines {
		fields := strings.Fields(v)
		// skip incorrect lines
		if len(fields) < 2 {
			continue
		}

		switch fields[0] {
		case "pgpgin":
			data.PgPgIn, err = strconv.ParseUint(fields[1], 10, 64)
			if err != nil {
				logger.Log.Debug.Printf(logger.MessageReadFileFieldError, ds.logPrefix, ds.filePath, "pgpgin", err.Error())
				// mandatory field
				return nil, err
			}
		case "pgpgout":
			data.PgPgOut, err = strconv.ParseUint(fields[1], 10, 64)
			if err != nil {
				logger.Log.Debug.Printf(logger.MessageReadFileFieldError, ds.logPrefix, ds.filePath, "pgpgout", err.Error())
				// mandatory field
				return nil, err
			}
		case "pswpin":
			data.PSwpIn, err = strconv.ParseUint(fields[1], 10, 64)
			if err != nil {
				logger.Log.Debug.Printf(logger.MessageReadFileFieldError, ds.logPrefix, ds.filePath, "pswpin", err.Error())
				// mandatory field
				return nil, err
			}
		case "pswpout":
			data.PSwpOut, err = strconv.ParseUint(fields[1], 10, 64)
			if err != nil {
				logger.Log.Debug.Printf(logger.MessageReadFileFieldError, ds.logPrefix, ds.filePath, "pgpgout", err.Error())
				// mandatory field
				return nil, err
			}
		case "pgfault":
			data.PgFault, err = strconv.ParseUint(fields[1], 10, 64)
			if err != nil {
				logger.Log.Debug.Printf(logger.MessageReadFileFieldError, ds.logPrefix, ds.filePath, "pgfault", err.Error())
				// mandatory field
				return nil, err
			}
		case "pgmajfault":
			data.PgMajFault, err = strconv.ParseUint(fields[1], 10, 64)
			if err != nil {
				logger.Log.Debug.Printf(logger.MessageReadFileFieldError, ds.logPrefix, ds.filePath, "pgmajfault", err.Error())
				// mandatory field
				return nil, err
			}
		case "pgfree":
			data.PgFree, err = strconv.ParseUint(fields[1], 10, 64)
			if err != nil {
				logger.Log.Debug.Printf(logger.MessageReadFileFieldError, ds.logPrefix, ds.filePath, "pgfree", err.Error())
				// mandatory field
				return nil, err
			}
		case "pgactivate":
			data.PgActivate, err = strconv.ParseUint(fields[1], 10, 64)
			if err != nil {
				logger.Log.Debug.Printf(logger.MessageReadFileFieldError, ds.logPrefix, ds.filePath, "pgactivate", err.Error())
				// mandatory field
				return nil, err
			}
		case "pgdeactivate":
			data.PgDeactivate, err = strconv.ParseUint(fields[1], 10, 64)
			if err != nil {
				logger.Log.Debug.Printf(logger.MessageReadFileFieldError, ds.logPrefix, ds.filePath, "pgdeactivate", err.Error())
				// mandatory field
				return nil, err
			}
		case "pglazyfree":
			data.PgLazyFree, err = strconv.ParseUint(fields[1], 10, 64)
			if err != nil {
				logger.Log.Debug.Printf(logger.MessageReadFileFieldError, ds.logPrefix, ds.filePath, "pglazyfree", err.Error())
				// mandatory field
				return nil, err
			}
		case "pglazyfreed":
			data.PgLazyFreed, err = strconv.ParseUint(fields[1], 10, 64)
			if err != nil {
				logger.Log.Debug.Printf(logger.MessageReadFileFieldError, ds.logPrefix, ds.filePath, "pglazyfreed", err.Error())
				// mandatory field
				return nil, err
			}
		case "pgrefill":
			data.PgRefill, err = strconv.ParseUint(fields[1], 10, 64)
			if err != nil {
				logger.Log.Debug.Printf(logger.MessageReadFileFieldError, ds.logPrefix, ds.filePath, "pgrefill", err.Error())
				// mandatory field
				return nil, err
			}
		case "numa_hit":
			data.NumaHit, err = strconv.ParseUint(fields[1], 10, 64)
			if err != nil {
				logger.Log.Debug.Printf(logger.MessageReadFileFieldError, ds.logPrefix, ds.filePath, "numa_hit", err.Error())
				// mandatory field
				return nil, err
			}
		case "numa_miss":
			data.NumaMiss, err = strconv.ParseUint(fields[1], 10, 64)
			if err != nil {
				logger.Log.Debug.Printf(logger.MessageReadFileFieldError, ds.logPrefix, ds.filePath, "numa_miss", err.Error())
				// mandatory field
				return nil, err
			}
		case "numa_foreign":
			data.NumaForeign, err = strconv.ParseUint(fields[1], 10, 64)
			if err != nil {
				logger.Log.Debug.Printf(logger.MessageReadFileFieldError, ds.logPrefix, ds.filePath, "numa_foreign", err.Error())
				// mandatory field
				return nil, err
			}
		case "numa_interleave":
			data.NumaInterleave, err = strconv.ParseUint(fields[1], 10, 64)
			if err != nil {
				logger.Log.Debug.Printf(logger.MessageReadFileFieldError, ds.logPrefix, ds.filePath, "numa_interleave", err.Error())
				// mandatory field
				return nil, err
			}
		case "numa_local":
			data.NumaLocal, err = strconv.ParseUint(fields[1], 10, 64)
			if err != nil {
				logger.Log.Debug.Printf(logger.MessageReadFileFieldError, ds.logPrefix, ds.filePath, "numa_local", err.Error())
				// mandatory field
				return nil, err
			}
		case "numa_other":
			data.NumaOther, err = strconv.ParseUint(fields[1], 10, 64)
			if err != nil {
				logger.Log.Debug.Printf(logger.MessageReadFileFieldError, ds.logPrefix, ds.filePath, "numa_other", err.Error())
				// mandatory field
				return nil, err
			}
		case "oom_kill":
			data.OOMKill, err = strconv.ParseUint(fields[1], 10, 64)
			if err != nil {
				logger.Log.Debug.Printf(logger.MessageReadFileFieldError, ds.logPrefix, ds.filePath, "oom_kill", err.Error())
				// mandatory field
				return nil, err
			}
		}
	}

	return &data, nil
}
