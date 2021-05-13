package leetcode

import (
	"testing"
)

func Test_uniquePaths(t *testing.T) {
	type args struct {
		m int
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"1",
			args{
				3, 7,
			},
			28,
		},
		{
			"2",
			args{
				3, 2,
			},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniquePaths(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("uniquePaths() = %v, want %v", got, tt.want)
			}
		})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniquePathsLoop(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("uniquePathsLoop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_uniquePathsWithObstacles(t *testing.T) {
	type args struct {
		obstacleGrid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"1",
			args{
				[][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}},
			},
			2,
		},
		{
			"2",
			args{
				[][]int{{0, 1}, {0, 0}},
			},
			1,
		},
		{
			"3",
			args{
				[][]int{{1}},
			},
			0,
		},
		{
			"4",
			args{
				[][]int{{0, 1}, {1, 0}},
			},
			0,
		},
		{
			"5",
			args{
				[][]int{{0, 0, 1, 0, 0}},
			},
			0,
		},
		{
			"6",
			args{
				[][]int{{1, 0, 0, 0, 0}},
			},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniquePathsWithObstacles(tt.args.obstacleGrid); got != tt.want {
				t.Errorf("uniquePathsWithObstacles() = %v, want %v", got, tt.want)
			}
		})
	}
}
