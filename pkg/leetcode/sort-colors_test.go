package leetcode

import (
	"log"
	"reflect"
	"testing"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func Test_sortColors(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				nums: []int{2, 0, 2, 1, 1, 0},
			},
			want: []int{0, 0, 1, 1, 2, 2},
		},
		{
			name: "2",
			args: args{
				nums: []int{2, 0, 1},
			},
			want: []int{0, 1, 2},
		},
		{
			name: "3",
			args: args{
				nums: []int{0},
			},
			want: []int{0},
		},
		{
			name: "4",
			args: args{
				nums: []int{1},
			},
			want: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortColors(tt.args.nums)
			if got := tt.args.nums; !reflect.DeepEqual(tt.args.nums, got) {
				t.Errorf("simplifyPath2(%v) = %v, want %v", tt.args.nums, got, tt.want)
			}
		})
	}
}
