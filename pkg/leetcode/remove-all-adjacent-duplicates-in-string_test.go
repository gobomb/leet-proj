package leetcode

import (
	"log"
	"testing"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func Test_removeDuplicatesFromString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				s: "abbaca",
			},
			want: "ca",
		},
		{
			name: "2",
			args: args{
				s: "azxxzy",
			},
			want: "ay",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicatesFromString(tt.args.s); got != tt.want {
				t.Errorf("removeDuplicatesFromString(%v) = %v, want %v", tt.args.s, got, tt.want)
			}
		})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicatesFromString2(tt.args.s); got != tt.want {
				t.Errorf("removeDuplicatesFromString2(%v) = %v, want %v", tt.args.s, got, tt.want)
			}
		})
	}
}
