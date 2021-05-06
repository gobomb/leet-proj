package leetcode

import "testing"

func Test_jump(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"1",
			args{
				[]int{2, 3, 1, 1, 4},
			},
			2,
		},
		{
			"2",
			args{
				[]int{2, 3, 0, 1, 4},
			},
			2,
		},
		{
			"3",
			args{
				[]int{2, 1, 1, 1, 4},
			},
			3,
		},
		{
			"4",
			args{
				[]int{0},
			},
			0,
		},
		{
			"5",
			args{
				[]int{5, 9, 3, 2, 1, 0, 2, 3, 3, 1, 0, 0},
			},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := jump(tt.args.nums); got != tt.want {
				t.Errorf("jump() = %v, want %v", got, tt.want)
			}
		})
	}
}
