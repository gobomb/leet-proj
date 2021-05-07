package leetcode

import "testing"

func Test_myPow(t *testing.T) {
	type args struct {
		x float64
		n int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			"1",
			args{
				2.00000,
				10,
			},
			1024.00000,
		},
		{
			"2",
			args{
				2.10000,
				3,
			},
			9.26100,
		},
		{
			"3",
			args{
				2.00000,
				-2,
			},
			0.25000,
		},
		{
			"4",
			args{34.00515,
				-3},
			0.00003,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myPow(tt.args.x, tt.args.n); got != tt.want {
				t.Errorf("myPow() = %v, want %v", got, tt.want)
			}
		})
	}
}
