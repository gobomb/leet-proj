package leetcode

import (
	"testing"
)

func Test_maxDepth(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				root: MakeTree([]int{3, 9, 20, Null, Null, 15, 7}),
			},
			want: 3,
		},
		{
			name: "2",
			args: args{
				root: MakeTree([]int{1, Null, 2}),
			},
			want: 2,
		},
		{
			name: "3",
			args: args{
				root: MakeTree([]int{}),
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDepth(tt.args.root); got != tt.want {
				t.Errorf("maxDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
