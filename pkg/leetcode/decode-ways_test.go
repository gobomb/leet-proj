package leetcode

import (
	"testing"
)

func Test_numDecodings(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				s: "12",
			},
			want: 2,
		},
		{
			name: "2",
			args: args{
				s: "226",
			},
			want: 3,
		},
		{
			name: "3",
			args: args{
				s: "0",
			},
			want: 0,
		},
		{
			name: "4",
			args: args{
				s: "2101",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numDecodings(tt.args.s); got != tt.want {
				t.Errorf("numDecodings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_numDecodings2(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				s: "12",
			},
			want: 2,
		},
		{
			name: "2",
			args: args{
				s: "226",
			},
			want: 3,
		},
		{
			name: "3",
			args: args{
				s: "0",
			},
			want: 0,
		},
		{
			name: "4",
			args: args{
				s: "2101",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numDecodings2(tt.args.s); got != tt.want {
				t.Errorf("numDecodings2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkNum(t *testing.T) {
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
				s: "12",
			},
			want: true,
		},
		{
			name: "2",
			args: args{
				s: "1",
			},
			want: true,
		},
		{
			name: "3",
			args: args{
				s: "2",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkNum(tt.args.s); got != tt.want {
				t.Errorf("checkNum() = %v, want %v", got, tt.want)
			}
		})
	}
}
