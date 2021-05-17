package leetcode

import "testing"

func Test_mySqrt(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"1",
			args{
				4,
			},
			2,
		},
		{
			"2",
			args{8},
			2,
		},
		{
			"3",
			args{1},
			1,
		},
		{
			"4",
			args{10000},
			100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mySqrt(tt.args.x); got != tt.want {
				t.Errorf("mySqrt() = %v, want %v", got, tt.want)
			}
		})
	}
}
