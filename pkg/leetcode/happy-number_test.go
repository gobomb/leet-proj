package leetcode

import (
	"log"
	"testing"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func Test_isHappy(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				n: 19,
			},
			want: true,
		},
		{
			name: "2",
			args: args{
				n: 2,
			},
			want: false,
		},
		{
			name: "3",
			args: args{
				n: 3,
			},
			want: false,
		},
		{
			name: "4",
			args: args{
				n: 123,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isHappy(tt.args.n); got != tt.want {
				t.Errorf("isHappy(%v) = %v, want %v", tt.args.n, got, tt.want)
			}
		})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isHappy2(tt.args.n); got != tt.want {
				t.Errorf("isHappy2(%v) = %v, want %v", tt.args.n, got, tt.want)
			}
		})
	}
}
