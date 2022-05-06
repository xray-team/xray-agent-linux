package diskSpace

import "testing"

func Test_rewriteMount(t *testing.T) {
	tests := []struct {
		name  string
		mount string // in
		want  string // out
	}{
		{
			mount: "/",
			want:  "root",
		},
		{
			mount: "/var",
			want:  "root_var",
		},
		{
			mount: "/var/lib",
			want:  "root_var_lib",
		},
	}
	for _, testCase := range tests {
		tt := testCase

		t.Run(tt.name, func(t *testing.T) {
			if got := rewriteMount(tt.mount); got != tt.want {
				t.Errorf("rewriteMount() = %v, want %v", got, tt.want)
			}
		})
	}
}
