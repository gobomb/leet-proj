package switchCase

import (
	"testing"
)

func Test_switchItf(t *testing.T) {
	type args struct {
		itf interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "[]int",
			args: args{
				[]int{1, 2, 3},
			},
		},
		{
			name: "[]string",
			args: args{
				[]string{"a", "b"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switchItf(tt.args.itf)
		})
	}
}
