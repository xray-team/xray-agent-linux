package entropy

import (
	"fmt"

	"github.com/xray-team/xray-agent-linux/reader"
)

type entropyDataSource struct {
	filePath  string
	logPrefix string
}

// NewDataSource returns a new DataSource.
func NewDataSource(filePath, logPrefix string) *entropyDataSource {
	if filePath == "" {
		return nil
	}

	return &entropyDataSource{
		filePath:  filePath,
		logPrefix: logPrefix,
	}
}

func (ds *entropyDataSource) GetData() (*Entropy, error) {
	var (
		entropy Entropy
		err     error
	)

	// read file to memory
	entropy.Available, err = reader.ReadInt64File(ds.filePath, ds.logPrefix)
	if err != nil {
		return nil, fmt.Errorf("cannot read entropy file %s. %s", ds.filePath, err)
	}

	return &entropy, nil
}
