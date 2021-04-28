package leetcode

import (
	"reflect"
	"testing"
)

func Test_combinationSum2(t *testing.T) {
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
			[][]int{{7}},
		},
		{
			"2",
			args{
				[]int{2, 3, 5},
				8,
			},
			[][]int{{3, 5}},
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
			[][]int{},
		},
		{
			"6",
			args{
				[]int{10, 1, 2, 7, 6, 1, 5},
				8,
			},
			[][]int{{1, 1, 6},
				{1, 2, 5},
				{1, 7},
				{2, 6}},
		},
		{
			"7",
			args{
				[]int{2, 5, 2, 1, 2},
				5,
			},
			[][]int{{1, 2, 2}, {5}},
		},
	}
	for _, tt := range tests {
		// if tt.name != "7" {
		// 	continue
		// }
		t.Run(tt.name, func(t *testing.T) {

			if got := combinationSum2(tt.args.candidates, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combinationSum2() = %#v, want %#v", got, tt.want)
			}
		})
	}
}
