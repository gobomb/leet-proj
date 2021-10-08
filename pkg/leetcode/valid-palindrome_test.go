package leetcode

import (
	"testing"
)

func Test_isPalindrome2(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				s: "A man, a plan, a canal: Panama",
			},
			want: true,
		},
		{
			name: "2",
			args: args{
				s: "race a car",
			},
			want: false,
		},
		{
			name: "3",
			args: args{
				s: "0P",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome2(tt.args.s); got != tt.want {
				t.Errorf("isPalindrome2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPalindrome4(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				s: "A man, a plan, a canal: Panama",
			},
			want: true,
		},
		{
			name: "2",
			args: args{
				s: "race a car",
			},
			want: false,
		},
		{
			name: "3",
			args: args{
				s: "0P",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome4(tt.args.s); got != tt.want {
				t.Errorf("isPalindrome4() = %v, want %v", got, tt.want)
			}
		})
	}
}
