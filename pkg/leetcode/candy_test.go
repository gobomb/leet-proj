package leetcode

import (
	"log"
	"testing"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func Test_candy(t *testing.T) {
	type args struct {
		ratings []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				ratings: []int{1, 0, 2},
			},
			want: 5,
		},
		{
			name: "2",
			args: args{
				ratings: []int{1, 2, 2},
			},
			want: 4,
		},
		{
			name: "3",
			args: args{
				ratings: []int{5, 4, 3, 2, 1},
			},
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := candy(tt.args.ratings); got != tt.want {
				t.Errorf("candy(%v) = %v, want %v", tt.args.ratings, got, tt.want)
			}
		})
	}
}
