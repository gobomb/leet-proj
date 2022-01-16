package easy

import (
	"testing"
)

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				s: "abccccdd",
			},
			want: 7,
		},
		{
			name: "",
			args: args{
				s: "a",
			},
			want: 1,
		},
		{
			name: "",
			args: args{
				s: "bb",
			},
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
