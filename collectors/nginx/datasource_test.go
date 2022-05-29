package nginx

import (
	"reflect"
	"testing"
)

func Test_parseStubStatus(t *testing.T) {
	tests := []struct {
		name    string
		data    []byte
		want    *StubStatus
		wantErr bool
	}{
		{
			data: []byte("Active connections: 1 \nserver accepts handled requests\n 2 3 4 \nReading: 5 Writing: 6 Waiting: 7 \n"),
			want: &StubStatus{
				Active:   1,
				Accepts:  2,
				Handled:  3,
				Requests: 4,
				Reading:  5,
				Writing:  6,
				Waiting:  7,
			},
			wantErr: false,
		},
	}
	for _, testCase := range tests {
		tt := testCase

		t.Run(tt.name, func(t *testing.T) {
			got, err := parseStubStatus(tt.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseStubStatus() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseStubStatus() got = %v, want %v", got, tt.want)
			}
		})
	}
}
