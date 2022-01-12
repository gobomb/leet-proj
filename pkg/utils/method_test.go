package utils

import (
	"fmt"
	"testing"
)

func TestDumpMethodSet(t *testing.T) {
	type Itest interface {
		test()
	}

	type args struct {
		i interface{}
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "",
			args: args{
				i: (*fmt.Formatter)(nil),
			},
		},
		{
			name: "",
			args: args{
				i: (*Itest)(nil),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DumpMethodSet(tt.args.i)
		})
	}
}
