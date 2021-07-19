package reader

import (
	"testing"
	"xray-agent-linux/logger"
)

func TestReadStringFile(t *testing.T) {
	logger.Init("")

	tests := []struct {
		caseDescription string
		filePath        string
		want            string
		wantErr         bool
	}{
		{
			caseDescription: "no file",
			filePath:        "../proc/testfiles/nofile",
			want:            "",
			wantErr:         true,
		},
		{
			caseDescription: "version",
			filePath:        "../proc/testfiles/version/version_signature-Mint19.2",
			want:            "Ubuntu 5.0.0-32.34~18.04.2-generic 5.0.21",
			wantErr:         false,
		},
	}

	for _, testCase := range tests {
		tt := testCase

		t.Run(tt.caseDescription, func(t *testing.T) {
			got, err := ReadStringFile(tt.filePath, "logPrefix")

			if (err != nil) != tt.wantErr {
				t.Errorf("readStringFile() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			if got != tt.want {
				t.Errorf("readStringFile() got = %v, want %v", got, tt.want)
			}
		})
	}
}
