package collectors

import "testing"

func TestConvertCarrierChangesToLinkFlaps(t *testing.T) {
	tests := []struct {
		name string
		in   int64 // in
		want int64 // out
	}{
		{
			in:   1,
			want: 0,
		},
		{
			in:   2,
			want: 0,
		},
		{
			in:   3,
			want: 0,
		},
		{
			in:   4,
			want: 1,
		},
		{
			in:   5,
			want: 1,
		},
		{
			in:   6,
			want: 2,
		},
		{
			in:   12,
			want: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertCarrierChangesToLinkFlaps(tt.in); got != tt.want {
				t.Errorf("ConvertCarrierChangesToLinkFlaps() = %v, want %v", got, tt.want)
			}
		})
	}
}
