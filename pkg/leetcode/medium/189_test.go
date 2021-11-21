package medium

import (
	"reflect"
	"testing"
)

func Test_rotate(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 6, 7},
				k:    3,
			},
			want: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			name: "",
			args: args{
				nums: []int{-1, -100, 3, 99},
				k:    2,
			},
			want: []int{3, 99, -1, -100},
		},
		{
			name: "",
			args: args{
				nums: []int{1, 2, 3, 4},
				k:    3,
			},
			want: []int{2, 3, 4, 1},
		},
		{
			name: "",
			args: args{
				nums: []int{1, 2},
				k:    3,
			},
			want: []int{2, 1},
		},
	}

	toTest := []func([]int, int){rotate, rotate1}
	for _, f := range toTest {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				nums := make([]int, len(tt.args.nums))
				copy(nums, tt.args.nums)
				rotate(nums, tt.args.k)

				if !reflect.DeepEqual(nums, tt.want) {
					t.Errorf("%v() = %v, want %v", funcName(f), nums, tt.want)
				}
			})
		}
	}
}
