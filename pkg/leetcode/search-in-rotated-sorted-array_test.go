package leetcode

import "testing"

func Test_search(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"1",
			args{[]int{4, 5, 6, 7, 0, 1, 2}, 0},
			4},
		{
			"2",
			args{[]int{4, 5, 6, 7, 0, 1, 2}, 3},
			-1},
		{
			"3",
			args{[]int{1}, 0},
			-1},
		{
			"4",
			args{[]int{1}, 1},
			0},
		{
			"5",
			args{[]int{1, 2}, 1},
			0},
		{
			"6",
			args{[]int{2, 1}, 1},
			1},
		{
			"7",
			args{[]int{4, 5, 6, 7, 8, 1, 2, 3}, 8},
			4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
