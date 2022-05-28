package template

import (
	"fmt"
)

type templateDataSource struct {
	logPrefix string
}

// NewDataSource returns a new DataSource.
func NewDataSource(logPrefix string) *templateDataSource {
	return &templateDataSource{
		logPrefix: logPrefix,
	}
}

func (ds *templateDataSource) GetData() (*Template, error) {
	var (
		data Template
		err  error
	)

	// Implement GetData logic here
	data.TemplateValue, err = 1, nil
	if err != nil {
		return nil, fmt.Errorf("cannot get data: %s", err.Error())
	}

	return &data, nil
}
