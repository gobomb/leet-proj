package leetcode

import (
	"log"
	"reflect"
	"testing"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func Test_inorderTraversal(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				root: MakeTree([]int{1, Null, 2, Null, Null, 3}),
			},
			want: []int{1, 3, 2},
		},
		{
			name: "2",
			args: args{
				root: MakeTree([]int{}),
			},
			want: []int{},
		},
		{
			name: "3",
			args: args{
				root: MakeTree([]int{1}),
			},
			want: []int{1},
		},
		{
			name: "4",
			args: args{
				root: MakeTree([]int{1, 2}),
			},
			want: []int{2, 1},
		},
		{
			name: "2",
			args: args{
				root: MakeTree([]int{1, Null, 2}),
			},
			want: []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			log.Printf("%+v\n", tt.args.root)
			if got := inorderTraversal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inorderTraversal(%+v) = %v, want %v", tt.args.root, got, tt.want)
			}
		})
	}
}

func Test_postorderTraversal(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				root: MakeTree([]int{1, Null, 2, Null, Null, 3}),
			},
			want: []int{3, 2, 1},
		},
		{
			name: "2",
			args: args{
				root: MakeTree([]int{}),
			},
			want: []int{},
		},
		{
			name: "3",
			args: args{
				root: MakeTree([]int{1}),
			},
			want: []int{1},
		},
		{
			name: "4",
			args: args{
				root: MakeTree([]int{1, 2}),
			},
			want: []int{2, 1},
		},
		{
			name: "5",
			args: args{
				root: MakeTree([]int{1, Null, 2}),
			},
			want: []int{2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			log.Printf("%+v\n", tt.args.root)
			if got := postorderTraversal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postorderTraversal(%+v) = %v, want %v", tt.args.root, got, tt.want)
			}
		})
	}
}
