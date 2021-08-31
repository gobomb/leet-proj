package leetcode

import (
	"log"
	"testing"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func Test_minWindow(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				s: "ADOBECODEBANC",
				t: "ABC",
			},
			want: "BANC",
		},
		{
			name: "2",
			args: args{
				s: "a",
				t: "a",
			},
			want: "a",
		},
		{
			name: "3",
			args: args{
				s: "a",
				t: "aa",
			},
			want: "",
		},
		{
			name: "4",
			args: args{
				s: "ADOBECODEBANC",
				t: "ECODE",
			},
			want: "ECODE",
		},
		{
			name: "5",
			args: args{
				s: "bba",
				t: "ab",
			},
			want: "ba",
		},
		{
			name: "6",
			args: args{
				// s: "b ccbaca aaba baabcb abbbbabbcca",
				s: "bccbacaaababaabcbabbbbabbcca",

				t: "cacca",
			},
			want: "ccbaca",
		},
		{
			name: "7",
			args: args{
				s: "aaaaaaaaaaaabbbbbcdd",
				t: "abcdd",
			},
			want: "abbbbbcdd",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minWindow(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("minWindow(%v, %v) = %v, want %v", tt.args.s, tt.args.t, got, tt.want)
			}
		})
	}
}

// 123

// 19244443142 5 67
// 19244443
//

// 314432
// 1
//
/*

ADOBECODEBANC

ABC

a
ad
ado
adob
adobe
adobec
dobec
obec
bec
beco
becod
becode

ecodeb
codeb
codeba

odeba
deba
eba
ba
ban
banc

*/
