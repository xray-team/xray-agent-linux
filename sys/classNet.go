package sys

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"xray-agent-linux/dto"
	"xray-agent-linux/reader"
)

const (
	ClassNetDir = "/sys/class/net/"

	NetDevUeventFile     = "uevent"
	NetDevTypeFile       = "type"
	NetDevSpeedFile      = "speed"
	NetDevOperStateFile  = "operstate"
	NetDevMACAddressFile = "address"
	NetDevBondMasterDir  = "bonding"
	NetDevBondSlaveDir   = "bonding_slave"
	NetDevDeviceDir      = "device"
	NetDevVirtualDir     = "/sys/devices/virtual/net"
)

type classNetDataSource struct {
	path      string
	logPrefix string
}

func NewClassNetDataSource(path, logPrefix string) *classNetDataSource {
	if path == "" {
		return nil
	}

	return &classNetDataSource{
		path:      path,
		logPrefix: logPrefix,
	}
}

// udevadm info /sys/class/net/lo
func (ds *classNetDataSource) GetData() (map[string]dto.ClassNet, error) {
	dirs, err := reader.ReadDir(ds.path, ds.logPrefix)
	if err != nil {
		return nil, fmt.Errorf("cannot read dir %s. err: %s", ds.path, err)
	}

	var out = make(map[string]dto.ClassNet)

	for _, dir := range dirs {
		var ifInventory dto.ClassNet

		// Uevent
		ifInventory.Uevent, err = ds.parseNetDevUevent(filepath.Join(ds.path, dir.Name(), NetDevUeventFile))
		if err != nil {
			continue
		}

		// ProtocolType
		ifInventory.ProtocolType, err = reader.ReadInt64File(filepath.Join(ds.path, dir.Name(), NetDevTypeFile), ds.logPrefix)
		if err != nil {
			continue
		}

		// Speed
		// Speed is not applied for some interfaces (lo, wireless, ...)
		ifInventory.Speed, _ = reader.ReadInt64File(filepath.Join(ds.path, dir.Name(), NetDevSpeedFile), ds.logPrefix)

		// OperState
		ifInventory.OperState, err = reader.ReadStringFile(filepath.Join(ds.path, dir.Name(), NetDevOperStateFile), ds.logPrefix)
		if err != nil {
			continue
		}

		// MACAddress
		ifInventory.MACAddress, err = reader.ReadStringFile(filepath.Join(ds.path, dir.Name(), NetDevMACAddressFile), ds.logPrefix)
		if err != nil {
			continue
		}

		ifInventory.Device = reader.IsExist(filepath.Join(ds.path, dir.Name(), NetDevDeviceDir), ds.logPrefix)
		ifInventory.Virtual = reader.IsExist(filepath.Join(NetDevVirtualDir, dir.Name()), ds.logPrefix)
		ifInventory.BondMaster = reader.IsExist(filepath.Join(ds.path, dir.Name(), NetDevBondMasterDir), ds.logPrefix)
		ifInventory.BondSlave = reader.IsExist(filepath.Join(ds.path, dir.Name(), NetDevBondSlaveDir), ds.logPrefix)

		ifInventory.Lower, ifInventory.Upper, err = ds.parseLowerUpper(filepath.Join(ds.path, dir.Name()))

		out[dir.Name()] = ifInventory
	}

	return out, nil
}

func (ds *classNetDataSource) parseNetDevUevent(filePath string) (*dto.NetDevUevent, error) {
	lines, err := reader.ReadMultilineFile(filePath, ds.logPrefix)
	if err != nil {
		return nil, err
	}

	var out dto.NetDevUevent

	for _, line := range lines {
		fields := strings.Split(line, "=")
		// skip incorrect lines
		if len(fields) < 2 {
			continue
		}

		if fields[0] == "DEVTYPE" {
			out.DevType = fields[1]

			continue
		}

		if fields[0] == "INTERFACE" {
			out.Interface = fields[1]

			continue
		}

		if fields[0] == "IFINDEX" {
			out.IfIndex, err = strconv.ParseInt(fields[1], 10, 64)
			if err != nil {
				return nil, err
			}
		}
	}

	// if dev type not set
	if out.DevType == "" {
		out.DevType = dto.NetDevTypeGeneric
	}

	return &out, nil
}

func (ds *classNetDataSource) parseLowerUpper(path string) ([]string, []string, error) {
	dirs, err := reader.ReadDir(path, ds.logPrefix)
	if err != nil {
		return nil, nil, fmt.Errorf("cannot read dir %s. err: %s", path, err)
	}

	reLower := regexp.MustCompile(`^lower_.+$`)
	reUpper := regexp.MustCompile(`^upper_.+$`)

	var lower, upper []string

	for _, dir := range dirs {
		if reLower.MatchString(dir.Name()) {
			lower = append(lower, strings.TrimPrefix(dir.Name(), "lower_"))
		}

		if reUpper.MatchString(dir.Name()) {
			upper = append(upper, strings.TrimPrefix(dir.Name(), "upper_"))
		}
	}

	return lower, upper, nil
}
