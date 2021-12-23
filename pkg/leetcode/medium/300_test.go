package medium

import (
	"reflect"
	"testing"
)

func Test_lengthOfLIS(t *testing.T) {
	type args struct {
		nums []int
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
				name: "",
				args: args{
					nums: []int{10, 9, 2, 5, 3, 7, 101, 18},
				},
				want: 4,
			},
			{
				name: "",
				args: args{
					nums: []int{0, 1, 0, 3, 2, 3},
				},
				want: 4,
			},
			{
				name: "",
				args: args{
					nums: []int{7, 7, 7, 7, 7, 7, 7},
				},
				want: 1,
			},
		}
	}

	toTest := []func([]int) int{lengthOfLIS, lengthOfLIS1}

	for _, f := range toTest {
		for _, tt := range tests() {
			t.Run(tt.name, func(t *testing.T) {
				if got := f(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("%v(%+v) = %+v, want %+v", funcName(f), tt.args, got, tt.want)
				}
			})
		}
	}
}
