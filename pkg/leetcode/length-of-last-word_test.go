package leetcode

import "testing"

func Test_lengthOfLastWord(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"1",
			args{"Hello World"},
			5,
		},
		{
			"2",
			args{" "},
			0,
		},
		{
			"3",
			args{
				"aa ",
			},
			2,
		},
		{
			"4",
			args{
				"aa",
			},
			2,
		},
		{
			"5",
			args{
				"bb aa  ",
			},
			2,
		},
		{
			"6",
			args{
				"  ",
			},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLastWord(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLastWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
