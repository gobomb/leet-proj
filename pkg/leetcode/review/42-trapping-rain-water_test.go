package review

import (
	"reflect"
	"testing"
)

func Test_trap(t *testing.T) {
	type args struct {
		height []int
	}

	tests := func() []struct {
		name string
		args args
		want int
	} {
		return []struct {
			name string
			args args
			want int
		}{
			{
				"1",
				args{
					[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
				},
				6,
			},
			{
				"2",
				args{
					[]int{4, 2, 0, 3, 2, 5},
				},
				9,
			},
		}
	}
	toTest := []func([]int) int{trap, trap2}

	for _, f := range toTest {
		for _, tt := range tests() {
			t.Run(tt.name, func(t *testing.T) {
				if got := f(tt.args.height); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("%v(%+v) = %+v, want %+v", funcName(f), tt.args, got, tt.want)
				}
			})
		}
	}
}
