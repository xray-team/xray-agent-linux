package netDev

/*
/proc/net/dev
/proc/$PID/net/dev - for containers
*/

import (
	"strconv"
	"strings"

	"github.com/xray-team/xray-agent-linux/logger"
	"github.com/xray-team/xray-agent-linux/reader"
)

type netDevDataSource struct {
	filePath  string
	logPrefix string
}

// NewNetDevDataSource returns a new DataSource.
func NewNetDevDataSource(filePath, logPrefix string) *netDevDataSource {
	if filePath == "" {
		return nil
	}

	return &netDevDataSource{
		filePath:  filePath,
		logPrefix: logPrefix,
	}
}

func (ds *netDevDataSource) GetData() (map[string]NetDevStatistics, error) {
	out := make(map[string]NetDevStatistics)

	lines, err := reader.ReadMultilineFile(ds.filePath, ds.logPrefix)
	if err != nil {
		return nil, err
	}

	for _, v := range lines {
		fields := strings.Fields(v)

		// skip table header and other unknown lines
		if len(fields) != 17 {
			continue
		}

		// parse interface name and trim ":"
		ifaceName := strings.TrimSuffix(fields[0], ":")

		var netDevStatistics NetDevStatistics

		// Rx.Bytes
		netDevStatistics.RxBytes, err = strconv.ParseUint(fields[1], 10, 64)
		if err != nil {
			logger.Log.Debug.Printf(logger.MessageReadFileFieldError, ds.logPrefix, ds.filePath, "Rx.Bytes", err.Error())

			continue
		}
		// Rx.Packets
		netDevStatistics.RxPackets, err = strconv.ParseUint(fields[2], 10, 64)
		if err != nil {
			logger.Log.Debug.Printf(logger.MessageReadFileFieldError, ds.logPrefix, ds.filePath, "Rx.Packets", err.Error())

			continue
		}
		// Rx.Errs
		netDevStatistics.RxErrs, err = strconv.ParseUint(fields[3], 10, 64)
		if err != nil {
			logger.Log.Debug.Printf(logger.MessageReadFileFieldError, ds.logPrefix, ds.filePath, "Rx.Errs", err.Error())

			continue
		}
		// Rx.Drop
		netDevStatistics.RxDrop, err = strconv.ParseUint(fields[4], 10, 64)
		if err != nil {
			logger.Log.Debug.Printf(logger.MessageReadFileFieldError, ds.logPrefix, ds.filePath, "Rx.Drop", err.Error())

			continue
		}
		// Rx.Fifo
		netDevStatistics.RxFifoErrs, err = strconv.ParseUint(fields[5], 10, 64)
		if err != nil {
			logger.Log.Debug.Printf(logger.MessageReadFileFieldError, ds.logPrefix, ds.filePath, "Rx.Fifo", err.Error())

			continue
		}
		// Rx.Frame
		netDevStatistics.RxFrameErrs, err = strconv.ParseUint(fields[6], 10, 64)
		if err != nil {
			logger.Log.Debug.Printf(logger.MessageReadFileFieldError, ds.logPrefix, ds.filePath, "Rx.Frame", err.Error())

			continue
		}
		// Rx.Compressed
		netDevStatistics.RxCompressed, err = strconv.ParseUint(fields[7], 10, 64)
		if err != nil {
			logger.Log.Debug.Printf(logger.MessageReadFileFieldError, ds.logPrefix, ds.filePath, "Rx.Compressed", err.Error())

			continue
		}
		// Rx.Multicast
		netDevStatistics.Multicast, err = strconv.ParseUint(fields[8], 10, 64)
		if err != nil {
			logger.Log.Debug.Printf(logger.MessageReadFileFieldError, ds.logPrefix, ds.filePath, "Rx.Multicast", err.Error())

			continue
		}
		// Tx.Bytes
		netDevStatistics.TxBytes, err = strconv.ParseUint(fields[9], 10, 64)
		if err != nil {
			logger.Log.Debug.Printf(logger.MessageReadFileFieldError, ds.logPrefix, ds.filePath, "Tx.Bytes", err.Error())

			continue
		}
		// Tx.Packets
		netDevStatistics.TxPackets, err = strconv.ParseUint(fields[10], 10, 64)
		if err != nil {
			logger.Log.Debug.Printf(logger.MessageReadFileFieldError, ds.logPrefix, ds.filePath, "Tx.Packets", err.Error())

			continue
		}
		// Tx.Errs
		netDevStatistics.TxErrs, err = strconv.ParseUint(fields[11], 10, 64)
		if err != nil {
			logger.Log.Debug.Printf(logger.MessageReadFileFieldError, ds.logPrefix, ds.filePath, "Tx.Errs", err.Error())

			continue
		}
		// Tx.Drop
		netDevStatistics.TxDrop, err = strconv.ParseUint(fields[12], 10, 64)
		if err != nil {
			logger.Log.Debug.Printf(logger.MessageReadFileFieldError, ds.logPrefix, ds.filePath, "Tx.Drop", err.Error())

			continue
		}
		// Tx.Fifo
		netDevStatistics.TxFifoErrs, err = strconv.ParseUint(fields[13], 10, 64)
		if err != nil {
			logger.Log.Debug.Printf(logger.MessageReadFileFieldError, ds.logPrefix, ds.filePath, "Tx.Fifo", err.Error())

			continue
		}
		// Tx.Colls
		netDevStatistics.Collisions, err = strconv.ParseUint(fields[14], 10, 64)
		if err != nil {
			logger.Log.Debug.Printf(logger.MessageReadFileFieldError, ds.logPrefix, ds.filePath, "Tx.Colls", err.Error())

			continue
		}
		// Tx.Carrier
		netDevStatistics.TxCarrierErrs, err = strconv.ParseUint(fields[15], 10, 64)
		if err != nil {
			logger.Log.Debug.Printf(logger.MessageReadFileFieldError, ds.logPrefix, ds.filePath, "Tx.Carrier", err.Error())

			continue
		}
		// Tx.Compressed
		netDevStatistics.TxCompressed, err = strconv.ParseUint(fields[16], 10, 64)
		if err != nil {
			logger.Log.Debug.Printf(logger.MessageReadFileFieldError, ds.logPrefix, ds.filePath, "Tx.Compressed", err.Error())

			continue
		}

		out[ifaceName] = netDevStatistics
	}

	return out, nil
}
