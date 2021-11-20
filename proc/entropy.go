package proc

import (
	"fmt"

	"xray-agent-linux/dto"
	"xray-agent-linux/reader"
)

type entropyDataSource struct {
	filePath  string
	logPrefix string
}

func NewEntropyDataSource(filePath, logPrefix string) *entropyDataSource {
	if filePath == "" {
		return nil
	}

	return &entropyDataSource{
		filePath:  filePath,
		logPrefix: logPrefix,
	}
}

func (ds *entropyDataSource) GetData() (*dto.Entropy, error) {
	var (
		entropy dto.Entropy
		err     error
	)

	// read file to memory
	entropy.Available, err = reader.ReadInt64File(ds.filePath, ds.logPrefix)
	if err != nil {
		return nil, fmt.Errorf("cannot read entropy file %s. %s", ds.filePath, err)
	}

	return &entropy, nil
}
