package ds

import (
	"reflect"
	"testing"
)

func Test_randSliceInt(t *testing.T) {
	type args struct {
		max    int
		length int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"1",
			args{
				50,
				10,
			},
			[]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := randSliceInt(tt.args.max, tt.args.length); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("randSliceInt() = %#v, want %v", got, tt.want)
			}
		})
	}
}
