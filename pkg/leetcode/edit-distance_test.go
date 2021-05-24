package leetcode

import (
	"log"
	"testing"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func Test_minDistance(t *testing.T) {
	type args struct {
		word1 string
		word2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				word1: "horse",
				word2: "ros",
			},
			want: 3,
		},
		{
			name: "5",
			args: args{
				word1: "intention",
				word2: "execution",
			},
			want: 5,
		},
		{
			name: "",
			args: args{
				word1: "",
				word2: "",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDistance(tt.args.word1, tt.args.word2); got != tt.want {
				t.Errorf("minDistance(%v, %v) = %v, want %v", tt.args.word1, tt.args.word2, got, tt.want)
			}
		})
	}
}
