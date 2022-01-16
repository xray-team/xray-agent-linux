package proc

import (
	"fmt"
	"strings"

	"github.com/xray-team/xray-agent-linux/dto"
	"github.com/xray-team/xray-agent-linux/reader"
)

const (
	netArpNumFields = 6
)

type netARPDataSource struct {
	filePath  string
	logPrefix string
}

// NewNetARPDataSource returns a new DataSource.
func NewNetARPDataSource(filePath, logPrefix string) *netARPDataSource {
	if filePath == "" {
		return nil
	}

	return &netARPDataSource{
		filePath:  filePath,
		logPrefix: logPrefix,
	}
}

func (ds *netARPDataSource) GetData() ([]dto.ARPEntry, error) {
	lines, err := reader.ReadMultilineFile(ds.filePath, ds.logPrefix)
	if err != nil {
		return nil, err
	}

	out := make([]dto.ARPEntry, 0, len(lines)-1) // minus header line

	for i, line := range lines {
		fields := strings.Fields(line)

		if len(fields) < netArpNumFields {
			return nil, fmt.Errorf("cann't parse arp file by path %s. wrong format of line %d: %v", ds.filePath, i+1, fields)
		}

		if strings.ToLower(fields[0]) == "ip" {
			err := checkNetArpHeaders(fields)
			if err != nil {
				return nil, fmt.Errorf("cann't parse file %s. %s", ds.filePath, err)
			}

			continue
		}

		entry := dto.ARPEntry{
			IP:        fields[0],
			HWType:    fields[1],
			Flags:     fields[2],
			HWAddress: fields[3],
			Mask:      fields[4],
			Device:    fields[5],
		}

		out = append(out, entry)
	}

	return out, nil
}

func checkNetArpHeaders(fs []string) error {
	// IP address       HW type     Flags       HW address            Mask     Device
	tmpl := []string{"ip", "address", "hw", "type", "flags", "hw", "address", "mask", "device"}
	if len(fs) < len(tmpl) {
		return fmt.Errorf("cann't parse arp file headers. got %v, want %v", fs, tmpl)
	}

	for i, t := range tmpl {
		if strings.ToLower(fs[i]) != t {
			return fmt.Errorf("cann't parse arp file headers. got %v, want %v", fs, tmpl)
		}
	}

	return nil
}
