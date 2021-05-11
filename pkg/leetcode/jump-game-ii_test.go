package leetcode

import "testing"

func Test_jump(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want2 bool
	}{
		{
			"1",
			args{
				[]int{2, 3, 1, 1, 4},
			},
			2,
			true,
		},
		{
			"2",
			args{
				[]int{2, 3, 0, 1, 4},
			},
			2,
			true,
		},
		{
			"3",
			args{
				[]int{2, 1, 1, 1, 4},
			},
			3,
			true,
		},
		{
			"4",
			args{
				[]int{0},
			},
			0,
			true,
		},
		{
			"5",
			args{
				[]int{5, 9, 3, 2, 1, 0, 2, 3, 3, 1, 0, 0},
			},
			3,
			true,
		},
		{
			"6",
			args{[]int{3, 2, 1, 0, 4}},
			-1,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := jump(tt.args.nums); got != tt.want {
				t.Errorf("jump() = %v, want %v", got, tt.want)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canJump(tt.args.nums); got != tt.want2 {
				t.Errorf("jump() = %v, want %v", got, tt.want2)
			}
		})
	}
}
