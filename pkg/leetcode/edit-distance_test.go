package leetcode

import (
	"log"
	"testing"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

// ref: https://rosettacode.org/wiki/Levenshtein_distance
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
			name: "2",
			args: args{
				word1: "intention",
				word2: "execution",
			},
			want: 5,
		},
		{
			name: "3",
			args: args{
				word1: "",
				word2: "",
			},
			want: 0,
		},
		{
			name: "4",
			args: args{
				word1: "kitten",
				word2: "sitting",
			},
			want: 3,
		},
		{
			name: "5",
			args: args{
				word1: "mississippi",
				word2: "swiss miss",
			},
			want: 8,
		},
		{
			name: "6",
			args: args{
				word1: "rosettacode",
				word2: "raisethysword",
			},
			want: 8,
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
