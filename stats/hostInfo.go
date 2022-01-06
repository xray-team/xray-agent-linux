package stats

import (
	"fmt"
	"os"
	"time"

	"github.com/xray-team/xray-agent-linux/conf"
	"github.com/xray-team/xray-agent-linux/dto"
)

func getHostInfo(cfg *conf.AgentConf) (*dto.HostInfo, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return nil, fmt.Errorf("error while getting hostname. err %s", err)
	}

	return &dto.HostInfo{
		HostName:   hostname,
		Timestamp:  time.Now().In(time.FixedZone(cfg.TimeZoneName, int(cfg.TimeZoneOffset)*60*60)).Unix(),
		Attributes: cfg.HostAttributes,
	}, nil
}
