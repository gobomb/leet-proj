package leetcode

import (
	"reflect"
	"testing"
)

func Test_nextPermutation(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{

		{
			"1",
			args{
				[]int{1, 2, 3},
			},
			[]int{1, 3, 2},
		},
		{
			"2",
			args{
				[]int{3, 2, 1},
			},
			[]int{1, 2, 3},
		},
		{
			"3",
			args{
				[]int{1, 1, 5},
			},
			[]int{1, 5, 1},
		},
		{
			"4",
			args{
				[]int{1},
			},
			[]int{1},
		},
		{
			"5",
			args{
				[]int{8, 9, 6, 10, 7, 2},
			},
			[]int{8, 9, 7, 2, 6, 10},
		},
		{
			"6",
			args{
				[]int{1, 2, 7, 4, 3, 1},
			},
			[]int{1, 3, 1, 2, 4, 7},
		},
		{
			"4",
			args{
				[]int{2, 3, 1},
			},
			[]int{3, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// arg := make([]int, len(tt.args.nums))
			// copy(arg, tt.args.nums)
			if nextPermutation(tt.args.nums); !reflect.DeepEqual(tt.args.nums, tt.want) {
				t.Errorf("isMatch() = %v, want %v", tt.args.nums, tt.want)
			}
		})
	}
}
