package leetcode

import (
	"reflect"
	"testing"
)

func Test_combinationSum(t *testing.T) {
	type args struct {
		candidates []int
		target     int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"1",
			args{
				[]int{2, 3, 6, 7},
				7,
			},
			[][]int{{2, 2, 3}, {7}},
		},
		{
			"2",
			args{
				[]int{2, 3, 5},
				8,
			},
			[][]int{{2, 3, 3}, {3, 5}, {2, 2, 2, 2}},
		},
		{
			"3",
			args{
				[]int{2},
				1,
			},
			[][]int{},
		},
		{
			"4",
			args{
				[]int{1},
				1,
			},
			[][]int{{1}},
		},
		{
			"5",
			args{
				[]int{1},
				2,
			},
			[][]int{{1, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combinationSum(tt.args.candidates, tt.args.target); !compSliceSlice(got, tt.want) {
				t.Errorf("combinationSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func compSliceSlice(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	cmap := make(map[int]bool)
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			if reflect.DeepEqual(a[i], b[j]) {
				_, ok := cmap[j]
				if ok {
					return false
				}
				cmap[j] = true
				break
			}
		}
	}
	if len(cmap) == len(a) {
		return true
	}
	return false
}
