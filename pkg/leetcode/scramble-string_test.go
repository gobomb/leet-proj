package leetcode

import "testing"

func Test_isScramble(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				s1: "great",
				s2: "rgeat",
			},
			want: true,
		},
		{
			name: "2",
			args: args{
				s1: "abcde",
				s2: "caebd",
			},
			want: false,
		},
		{
			name: "3",
			args: args{
				s1: "a",
				s2: "a",
			},
			want: true,
		},
		{
			name: "4",
			args: args{
				s1: "eebaacbcbcadaaedceaaacadccd",
				s2: "eadcaacabaddaceacbceaabeccd",
			},
			want: false,
		},
		//"eebaacbcbcadaaedceaaacadccd"
		// "eadcaacabaddaceacbceaabeccd"
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isScramble(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("isScramble() = %v, want %v", got, tt.want)
			}
		})
	}
}
