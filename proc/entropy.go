package proc

import (
	"fmt"

	"github.com/xray-team/xray-agent-linux/dto"
	"github.com/xray-team/xray-agent-linux/reader"
)

type entropyDataSource struct {
	filePath  string
	logPrefix string
}

// NewEntropyDataSource returns a new DataSource.
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
