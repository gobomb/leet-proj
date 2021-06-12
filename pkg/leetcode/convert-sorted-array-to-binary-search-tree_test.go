package leetcode

import (
	"reflect"
	"testing"
)

func Test_sortedArrayToBST(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "1",
			args: args{
				nums: []int{-10, -3, 0, 5, 9},
			},
			want: MakeTree([]int{0, -3, 9, -10, Null, 5}),
		},
		{
			name: "2",
			args: args{
				nums: []int{1, 3},
			},
			want: MakeTree([]int{3, 1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortedArrayToBST(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortedArrayToBST() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
