package tree

import (
	"log"
	"reflect"
	"testing"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

var trees = [][]int{{1, Null, 2, Null, Null, Null, 3},
	{2, 1, 3},
	{3, 1, Null, Null, 2},
	{3, 2, Null, 1},
	{1, 2, Null, 3, 4, Null, Null, Null, 5, 6},
}

func TestMakeTree(t *testing.T) {
	type args struct {
		vals []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "1",
			args: args{
				vals: []int{1, 2, 3},
			},
			want: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
		},
		{
			name: "2",
			args: args{
				vals: []int{1, Null, 2, Null, Null, 3},
			},
			want: &TreeNode{
				Val:  1,
				Left: nil,
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MakeTree(tt.args.vals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeTree(%v) = %v, want %v", tt.args.vals, got, tt.want)
			} else {
				log.Printf("%+v\n", got)
			}
		})
	}
}

func Test_getTreeHeight(t *testing.T) {
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
				root: MakeTree(trees[0]),
			},
			want: 3,
		},
		{
			name: "2",
			args: args{
				root: MakeTree(trees[1]),
			},
			want: 2,
		},
		{
			name: "3",
			args: args{
				root: MakeTree(trees[2]),
			},
			want: 3,
		},
		{
			name: "4",
			args: args{
				root: MakeTree(trees[3]),
			},
			want: 3,
		},
		{
			name: "5",
			args: args{
				root: MakeTree(trees[4]),
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getTreeHeight(tt.args.root); got != tt.want {
				t.Errorf("getTreeHeight(%v) = %v, want %v", tt.args.root, got, tt.want)
			} else {
				log.Printf("%+v\n", tt.args.root)
			}
		})
	}
}
