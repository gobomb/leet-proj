package leetcode

import (
	"log"
	"testing"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func Test_searchRotated(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				nums:   []int{2, 5, 6, 0, 0, 1, 2},
				target: 0,
			},
			want: true,
		},
		{
			name: "2",
			args: args{
				nums:   []int{2, 5, 6, 0, 0, 1, 2},
				target: 3,
			},
			want: false,
		},
		{
			name: "1",
			args: args{
				nums:   []int{0, 1, 2, 4, 4, 4, 5, 6, 6, 7},
				target: 5,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// if got := searchRotated(tt.args.nums, tt.args.target); got != tt.want {
			// 	t.Errorf("searchRotated(%v, %v) = %v, want %v", tt.args.nums, tt.args.target, got, tt.want)
			// }
		})
	}
}
