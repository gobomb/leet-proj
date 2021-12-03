package easy

import (
	"reflect"
	"testing"
)

func Test_moveZeroes(t *testing.T) {
	type args struct {
		nums []int
	}

	tests := func() []struct {
		name string
		args args
		want []int
	} {
		return []struct {
			name string
			args args
			want []int
		}{
			{
				name: "",
				args: args{
					nums: []int{0, 1, 0, 3, 12},
				},
				want: []int{1, 3, 12, 0, 0},
			},
			{
				name: "",
				args: args{
					nums: []int{0},
				},
				want: []int{0},
			},
			{
				name: "",
				args: args{
					nums: []int{0, 1, 0, 3, 12, 0, 0, 0, 0, 1},
				},
				want: []int{1, 3, 12, 1, 0, 0, 0, 0, 0, 0},
			},
		}
	}

	toTest := []func([]int){moveZeroes, moveZeroes1}

	for _, f := range toTest {
		for _, tt := range tests() {
			t.Run(tt.name, func(t *testing.T) {
				if f(tt.args.nums); !reflect.DeepEqual(tt.args.nums, tt.want) {
					t.Errorf("%v(%+v) = %+v, want %+v", funcName(f), tt.args, tt.args.nums, tt.want)
				}
			})
		}
	}
}
