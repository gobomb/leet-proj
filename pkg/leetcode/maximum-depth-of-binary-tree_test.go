package leetcode

import (
	"log"
	"testing"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

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

func Test_minDepth(t *testing.T) {
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
			want: 2,
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
		{
			name: "4",
			args: args{
				root: MakeTree([]int{2, Null, 3, Null, Null, Null, 4, Null,
					Null, Null, Null, Null, Null, Null, 5}),
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDepth(tt.args.root); got != tt.want {
				t.Errorf("minDepth(%+v) = %v, want %v", tt.args.root, got, tt.want)
			}
		})
	}
}
