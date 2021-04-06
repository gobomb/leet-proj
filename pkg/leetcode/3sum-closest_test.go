package leetcode

import "testing"

func Test_threeSumClosest(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{nums: []int{-1, 2, 1, -4}, target: 1}, 2},
		{"2", args{nums: []int{-1, 2, 1, 4}, target: 1}, 2},
		{"3", args{nums: []int{2, 2, 2, 2}, target: 1}, 6},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSumClosest(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("threeSumClosest() = %v, want %v", got, tt.want)
			}
		})
	}
}
