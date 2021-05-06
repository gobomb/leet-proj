package leetcode

import (
	"reflect"
	"testing"
)

func Test_permute(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"1",
			args{[]int{1, 2, 3}},
			[][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}},
		},
		{
			"2",
			args{[]int{0, 1}},
			[][]int{{0, 1}, {1, 0}},
		},
		{
			"3",
			args{[]int{1}},
			[][]int{{1}},
		},
		{
			"4",
			args{[]int{5, 4, 6, 2}},
			[][]int{
				{5, 4, 6, 2}, {5, 4, 2, 6}, {5, 6, 4, 2}, {5, 6, 2, 4}, {5, 2, 4, 6}, {5, 2, 6, 4}, {4, 5, 6, 2}, {4, 5, 2, 6}, {4, 6, 5, 2}, {4, 6, 2, 5}, {4, 2, 5, 6}, {4, 2, 6, 5}, {6, 5, 4, 2}, {6, 5, 2, 4}, {6, 4, 5, 2}, {6, 4, 2, 5}, {6, 2, 5, 4}, {6, 2, 4, 5}, {2, 5, 4, 6}, {2, 5, 6, 4}, {2, 4, 5, 6}, {2, 4, 6, 5}, {2, 6, 5, 4}, {2, 6, 4, 5},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := permute(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("permute() = %v, want %v", got, tt.want)
			}
		})
	}
	tests = append(tests, []struct {
		name string
		args args
		want [][]int
	}{{
		"5",
		args{[]int{1, 1, 2}},
		[][]int{{1, 1, 2}, {1, 2, 1}, {2, 1, 1}},
	}, {
		"6",
		args{[]int{1, 1, 1}},
		[][]int{{1, 1, 1}},
	}}...,
	)
	// for _, tt := range tests {
	// 	t.Run(tt.name, func(t *testing.T) {
	// 		if got := permuteUnique2(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
	// 			t.Errorf("permute() = %v, want %v", got, tt.want)
	// 		}
	// 	})
	// }
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := permuteUnique(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("permute() = %v, want %v", got, tt.want)
			}
		})
	}
}
