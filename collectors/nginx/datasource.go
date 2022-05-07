package nginx

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"

	"github.com/xray-team/xray-agent-linux/conf"
	"github.com/xray-team/xray-agent-linux/logger"
)

type StubStatusClient struct {
	config    conf.NginxStubStatus
	client    *http.Client
	logPrefix string
}

func NewStubStatusClient(config *conf.NginxStubStatus, client *http.Client, logPrefix string) *StubStatusClient {
	if config == nil || client == nil {
		return nil
	}

	return &StubStatusClient{
		config:    *config,
		client:    client,
		logPrefix: logPrefix,
	}
}

func (ds *StubStatusClient) GetData() (*StubStatus, error) {
	req, err := http.NewRequest("GET", ds.config.Endpoint, nil)
	if err != nil {
		return nil, err
	}

	// logger
	logger.Log.Debug.Printf(logger.MessageHttpRequest, ds.logPrefix, req.RequestURI)

	resp, err := ds.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code: %v", resp.StatusCode)
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return parseStubStatus(respBody)
}

func parseStubStatus(data []byte) (*StubStatus, error) {
	var (
		out StubStatus
		err error
	)

	re := regexp.MustCompile(`(Active connections: )(\d*)( \nserver accepts handled requests\n )(\d*)( )(\d*)( )(\d*)( \nReading: )(\d*)( Writing: )(\d*)( Waiting: )(\d*)( \n)`)

	groups := re.FindStringSubmatch(string(data))
	if len(groups) != 16 {
		return nil, fmt.Errorf("parseStubStatus: not enough re groups")
	}

	out.Active, err = strconv.ParseUint(groups[2], 10, 64)
	if err != nil {
		return nil, err
	}

	out.Accepts, err = strconv.ParseUint(groups[4], 10, 64)
	if err != nil {
		return nil, err
	}

	out.Handled, err = strconv.ParseUint(groups[6], 10, 64)
	if err != nil {
		return nil, err
	}

	out.Requests, err = strconv.ParseUint(groups[8], 10, 64)
	if err != nil {
		return nil, err
	}

	out.Reading, err = strconv.ParseUint(groups[10], 10, 64)
	if err != nil {
		return nil, err
	}

	out.Writing, err = strconv.ParseUint(groups[12], 10, 64)
	if err != nil {
		return nil, err
	}

	out.Waiting, err = strconv.ParseUint(groups[14], 10, 64)
	if err != nil {
		return nil, err
	}

	return &out, nil
}
