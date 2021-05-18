package leetcode

import "testing"

func Test_climbStairs(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"1",
			args{2},
			2,
		},
		{
			"2",
			args{3},
			3,
		},
		{
			"3",
			args{4},
			5,
		},
		{
			"4",
			args{44},
			1134903170,
		},
		{
			"5",
			args{5},
			8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbStairs(tt.args.n); got != tt.want {
				t.Errorf("climbStairs() = %v, want %v", got, tt.want)
			}
		})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dpClimbStairs(tt.args.n); got != tt.want {
				t.Errorf("dpClimbStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
