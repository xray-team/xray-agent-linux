package run

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"

	"github.com/xray-team/xray-agent-linux/dto"
)

type iwconfigDataSource struct {
	cmdRunner *cmdRunner
}

func NewIwconfigDataSource(runner *cmdRunner) *iwconfigDataSource {
	if runner == nil {
		return nil
	}

	return &iwconfigDataSource{
		cmdRunner: runner,
	}
}

func (ds *iwconfigDataSource) GetInterfaceData(ifName string) (*dto.Iwconfig, error) {
	if err := os.Setenv("PATH", "PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"); err != nil {
		return nil, err
	}

	// Prepare command
	cmd := exec.Command("iwconfig", ifName)

	// Execute command
	stdout, stderr, err := ds.cmdRunner.Run(cmd)
	if err != nil {
		return nil, err
	}

	if stderr != "" {
		return nil, fmt.Errorf("%s", stderr)
	}

	var out dto.Iwconfig

	// SSID
	out.SSID, err = parseIwconfigSSID(stdout)
	if err != nil {
		return nil, fmt.Errorf("iwconfig parsing error")
	}

	// Frequency
	out.Frequency, err = parseIwconfifFrequency(stdout)
	if err != nil {
		return nil, fmt.Errorf("iwconfig parsing error")
	}

	// AccessPoint
	out.AccessPoint, err = parseIwconfigAccessPoint(stdout)
	if err != nil {
		return nil, fmt.Errorf("iwconfig parsing error: %w", err)
	}

	// BitRate
	out.BitRate, err = parseIwconfigBitRate(stdout)
	if err != nil {
		return nil, fmt.Errorf("iwconfig parsing error: %w", err)
	}

	// TxPower
	out.TxPower, err = parseIwconfigTxPower(stdout)
	if err != nil {
		return nil, fmt.Errorf("iwconfig parsing error: %w", err)
	}

	// SignalLevel
	out.SignalLevel, err = parseIwconfigSignalLevel(stdout)
	if err != nil {
		return nil, fmt.Errorf("iwconfig parsing error: %w", err)
	}

	// LinkQuality
	out.LinkQuality, out.LinkQualityLimit, err = parseIwconfigLinkQuality(stdout)
	if err != nil {
		return nil, fmt.Errorf("iwconfig parsing error: %w", err)
	}

	// RxInvalidNwid
	out.RxInvalidNwid, err = parseIwconfigRxInvalidNwid(stdout)
	if err != nil {
		return nil, fmt.Errorf("iwconfig parsing error")
	}

	// RxInvalidCrypt
	out.RxInvalidCrypt, err = parseIwconfigRxInvalidCrypt(stdout)
	if err != nil {
		return nil, fmt.Errorf("iwconfig parsing error: %w", err)
	}

	// RxInvalidFrag
	out.RxInvalidFrag, err = parseIwconfigRxInvalidFrag(stdout)
	if err != nil {
		return nil, fmt.Errorf("iwconfig parsing error: %w", err)
	}

	// TxExcessiveRetries
	out.TxExcessiveRetries, err = parseIwconfigTxExcessiveRetries(stdout)
	if err != nil {
		return nil, fmt.Errorf("iwconfig parsing error: %w", err)
	}

	// InvalidMisc
	out.InvalidMisc, err = parseIwconfigInvalidMisc(stdout)
	if err != nil {
		return nil, fmt.Errorf("iwconfig parsing error: %w", err)
	}

	// MissedBeacon
	out.MissedBeacon, err = parseIwconfigMissedBeacon(stdout)
	if err != nil {
		return nil, fmt.Errorf("iwconfig parsing error: %w", err)
	}

	return &out, nil
}

func parseIwconfigSSID(stdout string) (string, error) {
	re := regexp.MustCompile(`ESSID:"(.{1,32}?)"`)

	groups := re.FindStringSubmatch(stdout)
	if len(groups) != 2 {
		return "", fmt.Errorf("not enough re groups")
	}

	return groups[1], nil
}

func parseIwconfifFrequency(stdout string) (float64, error) {
	re := regexp.MustCompile(`Frequency:(.*)\sGHz`)

	groups := re.FindStringSubmatch(stdout)
	if len(groups) != 2 {
		return 0, fmt.Errorf("not enough re groups")
	}

	return strconv.ParseFloat(groups[1], 64)
}

func parseIwconfigAccessPoint(stdout string) (string, error) {
	re := regexp.MustCompile(`Access\sPoint:\s((((\d|([a-f]|[A-F])){2}:){5}(\d|([a-f]|[A-F])){2})|Not-Associated)`)

	groups := re.FindStringSubmatch(stdout)
	if len(groups) <= 2 {
		return "", fmt.Errorf("not enough re groups")
	}

	return groups[1], nil
}

func parseIwconfigBitRate(stdout string) (float64, error) {
	re := regexp.MustCompile(`Bit\sRate=(.*)\s(Mb/s|Gb/s|kb/s)`)

	groups := re.FindStringSubmatch(stdout)
	if len(groups) != 3 {
		return 0, fmt.Errorf("not enough re groups")
	}

	bitRate, err := strconv.ParseFloat(groups[1], 64)
	if err != nil {
		return 0, err
	}

	switch groups[2] {
	case "Mb/s":
	case "Gb/s":
		bitRate *= 1000
	case "kb/s":
		bitRate /= 1000
	default:
		return 0, fmt.Errorf("unknown BitRate suffix")
	}

	return bitRate, nil
}

func parseIwconfigTxPower(stdout string) (int64, error) {
	re := regexp.MustCompile(`Tx-Power=(.*?)\sdBm`)

	groups := re.FindStringSubmatch(stdout)

	if len(groups) != 2 {
		return 0, fmt.Errorf("not enough re groups")
	}

	return strconv.ParseInt(groups[1], 10, 64)
}

func parseIwconfigSignalLevel(stdout string) (int64, error) {
	re := regexp.MustCompile(`Signal\slevel=(.*)\sdBm`)

	groups := re.FindStringSubmatch(stdout)
	if len(groups) != 2 {
		return 0, fmt.Errorf("not enough re groups")
	}

	return strconv.ParseInt(groups[1], 10, 64)
}

func parseIwconfigLinkQuality(stdout string) (int64, int64, error) {
	re := regexp.MustCompile(`Link\sQuality=(\d*)/(\d*)`)
	groups := re.FindStringSubmatch(stdout)

	if len(groups) != 3 {
		return 0, 0, fmt.Errorf("not enough re groups")
	}

	linkQuality, err := strconv.ParseInt(groups[1], 10, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("linkQuality parsing error: %w", err)
	}

	linkQualityLimit, err := strconv.ParseInt(groups[2], 10, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("linkQualityLimit parsing error: %w", err)
	}

	return linkQuality, linkQualityLimit, nil
}

func parseIwconfigRxInvalidNwid(stdout string) (uint64, error) {
	re := regexp.MustCompile(`Rx\sinvalid\snwid:(\d*)`)
	groups := re.FindStringSubmatch(stdout)

	if len(groups) != 2 {
		return 0, fmt.Errorf("not enough re groups")
	}

	return strconv.ParseUint(groups[1], 10, 64)
}

func parseIwconfigRxInvalidCrypt(stdout string) (uint64, error) {
	re := regexp.MustCompile(`Rx\sinvalid\scrypt:(\d*)`)
	groups := re.FindStringSubmatch(stdout)

	if len(groups) != 2 {
		return 0, fmt.Errorf("not enough re groups")
	}

	return strconv.ParseUint(groups[1], 10, 64)
}

func parseIwconfigRxInvalidFrag(stdout string) (uint64, error) {
	re := regexp.MustCompile(`Rx\sinvalid\sfrag:(\d*)`)
	groups := re.FindStringSubmatch(stdout)

	if len(groups) != 2 {
		return 0, fmt.Errorf("not enough re groups")
	}

	return strconv.ParseUint(groups[1], 10, 64)
}

func parseIwconfigTxExcessiveRetries(stdout string) (uint64, error) {
	re := regexp.MustCompile(`Tx\sexcessive\sretries:(\d*)`)
	groups := re.FindStringSubmatch(stdout)

	if len(groups) != 2 {
		return 0, fmt.Errorf("not enough re groups")
	}

	return strconv.ParseUint(groups[1], 10, 64)
}

func parseIwconfigInvalidMisc(stdout string) (uint64, error) {
	re := regexp.MustCompile(`Invalid\smisc:(\d*)`)
	groups := re.FindStringSubmatch(stdout)

	if len(groups) != 2 {
		return 0, fmt.Errorf("not enough re groups")
	}

	return strconv.ParseUint(groups[1], 10, 64)
}

func parseIwconfigMissedBeacon(stdout string) (uint64, error) {
	re := regexp.MustCompile(`Missed\sbeacon:(\d*)`)
	groups := re.FindStringSubmatch(stdout)

	if len(groups) != 2 {
		return 0, fmt.Errorf("not enough re groups")
	}

	return strconv.ParseUint(groups[1], 10, 64)
}
