package leetcode

import "testing"

func Test_guessNumber(t *testing.T) {
	type args struct {
		n    int
		pick int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"1",
			args{
				10,
				6,
			},
			6,
		},
		{
			"2",
			args{
				1,
				1,
			},
			1,
		},
		{
			"3",
			args{
				2,
				1,
			},
			1,
		},
		{
			"4",
			args{
				2,
				2,
			},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pick = tt.args.pick
			if got := guessNumber(tt.args.n); got != tt.want {
				t.Errorf("guessNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
