package algorithm

import (
	"justest/pkg/utils"
	"testing"
)

func TestBin(t *testing.T) {
	type args struct {
		arr []int
		k   int
	}

	// sorted := genSortedSlices(3)
	// log.Printf("%#v", sorted)
	// [][]int{[]int{9, 12, 12, 17, 26, 64, 66, 67, 76, 85}, []int{0, 13, 29, 50, 53, 61, 70, 93, 95, 99}, []int{0, 14, 23, 23, 28, 40, 50, 66, 86, 92}}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"1",
			args{
				[]int{9, 12, 12, 17, 26, 64, 66, 67, 76, 85},
				76,
			},
			8,
		},
		{
			"2",
			args{
				[]int{0, 13, 29, 50, 53, 61, 70, 93, 95, 99},
				100,
			},
			-1,
		},
		{
			"3",
			args{
				[]int{0, 14, 23, 23, 28, 40, 50, 66, 86, 92},
				15,
			},
			-1,
		},
		{
			"4",
			args{
				[]int{0, 14, 23, 23, 28, 40, 50, 66, 86, 92},
				92,
			},
			9,
		},
	}

	testFunc := []func([]int, int) int{
		BinSearch,
		BinSearch1,
	}
	for _, f := range testFunc {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := f(tt.args.arr, tt.args.k); got != tt.want {
					t.Errorf("%v() = %v, want %v", utils.FuncName(f), got, tt.want)
				}
			})
		}
	}
}