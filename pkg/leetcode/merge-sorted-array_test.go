package leetcode

import (
	"reflect"
	"testing"
)

func Test_merge88(t *testing.T) {
	type args struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				nums1: []int{1, 2, 3, 0, 0, 0},
				m:     3,
				nums2: []int{2, 5, 6},
				n:     3,
			},
			want: []int{1, 2, 2, 3, 5, 6},
		},
		{
			name: "2",
			args: args{
				nums1: []int{1},
				m:     1,
				nums2: []int{},
				n:     0,
			},
			want: []int{1},
		},
		{
			name: "3",
			args: args{
				nums1: []int{},
				m:     0,
				nums2: []int{},
				n:     0,
			},
			want: []int{},
		},
		{
			name: "4",
			args: args{
				nums1: []int{1, 2, 3, 4, 0, 0, 0, 0, 0},
				m:     4,
				nums2: []int{1, 2, 3, 4, 5},
				n:     5,
			},
			want: []int{1, 1, 2, 2, 3, 3, 4, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merge88(tt.args.nums1, tt.args.m, tt.args.nums2, tt.args.n)
			if !reflect.DeepEqual(tt.want, tt.args.nums1) {
				t.Fatalf("failed want %v but got %v", tt.want, tt.args.nums1)
			}
		})
	}
}
