package leetcode

import (
	"log"
	"testing"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func Test_flatten(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "1",
			args: args{
				root: MakeTree([]int{1, 2, 5, 3, 4, null, 6}),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flatten(tt.args.root)
			if checkHasLeft(tt.args.root) {
				t.Errorf("%+v\n", tt.args.root)
			}
		})
	}
}

func Test_preorderTraversalFlatten2(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "1",
			args: args{
				root: MakeTree([]int{1, 2, 5, 3, 4, null, 6}),
			},
		},
		{
			name: "2",
			args: args{
				root: MakeTree([]int{}),
			},
		},
		{
			name: "3",
			args: args{
				root: MakeTree([]int{1, 2}),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			preorderTraversalFlatten2(tt.args.root)
			if checkHasLeft(tt.args.root) {
				t.Errorf("%+v\n", tt.args.root)
			}
		})
	}
}

func checkHasLeft(root *TreeNode) bool {
	for root != nil {
		if root.Left != nil {
			return true
		}
		root = root.Right
	}
	return false
}

func Test_flatten2(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "1",
			args: args{
				root: MakeTree([]int{1, 2, 5, 3, 4, null, 6}),
			},
		},
		{
			name: "2",
			args: args{
				root: MakeTree([]int{}),
			},
		},
		{
			name: "3",
			args: args{
				root: MakeTree([]int{1, 2}),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flatten2(tt.args.root)
			if checkHasLeft(tt.args.root) {
				t.Errorf("%+v\n", tt.args.root)
			}
		})
	}
}
