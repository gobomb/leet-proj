package leetcode

import "testing"

func Test_isMatchwildcard(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"1",
			args{
				"aa", "a",
			},
			false,
		},
		{
			"2",
			args{
				"aa", "*",
			},
			true,
		},
		{
			"3",
			args{
				"cb", "?a",
			},
			false,
		},
		{
			"4",
			args{
				"adceb", "*a*b",
			},
			true,
		},
		{
			"5",
			args{
				"acdcb",
				"a*c?b",
			},
			false,
		},
		{
			"6",
			args{
				"dab", "*ab",
			},
			true,
		},
		{
			"7",
			args{
				"",
				"******",
			},
			true,
		},
		{
			"8",
			args{
				"aaaa",
				"***a",
			},
			true,
		},
		{
			"9",
			args{
				"a",
				"",
			},
			false,
		},
		{
			"10",
			args{
				"aa",
				"a",
			},
			false,
		},
		{
			"11",
			args{
				"c",
				"*?*",
			},
			true,
		},
		{
			"12",
			args{
				"bbabab",
				"**?a*",
			},
			true,
		},
		{
			"13",
			args{
				"abbabaaabbabbaababbabbbbbabbbabbbabaaaaababababbbabababaabbababaabbbbbbaaaabababbbaabbbbaabbbbababababbaabbaababaabbbababababbbbaaabbbbbabaaaabbababbbbaababaabbababbbbbababbbabaaaaaaaabbbbbaabaaababaaaabb",
				"**aa*****ba*a*bb**aa*ab****a*aaaaaa***a*aaaa**bbabb*b*b**aaaaaaaaa*a********ba*bbb***a*ba*bb*bb**a*b*bb",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatchwildcard(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("isMatchwildcard() = %v, want %v", got, tt.want)
			}
		})
	}
}
