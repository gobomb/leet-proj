package leetcode

import (
	"reflect"
	"testing"
)

func Test_merge(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"1",
			args{
				[][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			},
			[][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			"2",
			args{
				[][]int{{1, 4}, {4, 5}},
			},
			[][]int{{1, 5}},
		},
		{
			"3",
			args{
				[][]int{{3, 7}, {1, 4}, {4, 5}},
			},
			[][]int{{1, 7}},
		},
		{
			"4",
			args{
				[][]int{{3, 7}, {1, 8}, {4, 5}},
			},
			[][]int{{1, 8}},
		},
		{
			"5",
			args{
				[][]int{{3, 7}, {3, 7}, {4, 5}},
			},
			[][]int{{3, 7}},
		},
		{
			"6",
			args{
				[][]int{{1, 2}, {3, 7}},
			},
			[][]int{{1, 7}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := merge(tt.args.intervals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_insert(t *testing.T) {
	type args struct {
		intervals   [][]int
		newInterval []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{
			"1",
			args{
				[][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
				[]int{1, 3},
			},
			[][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			"2",
			args{
				[][]int{{1, 4}, {4, 5}},
				[]int{5, 9},
			},
			[][]int{{1, 9}},
		},
		{
			"3",
			args{
				[][]int{{3, 7}, {1, 4}, {4, 5}},
				[]int{0, 0},
			},
			[][]int{{0, 7}},
		},
		{
			"4",
			args{
				[][]int{{3, 7}, {1, 8}, {4, 5}},
				[]int{10, 11},
			},
			[][]int{{1, 8}, {10, 11}},
		},
		{
			"5",
			args{
				[][]int{{3, 7}, {3, 7}, {4, 5}},
				[]int{1, 2},
			},
			[][]int{{1, 7}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insert(tt.args.intervals, tt.args.newInterval); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insert() = %v, want %v", got, tt.want)
			}
		})
	}
}
