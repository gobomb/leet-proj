package leetcode

import (
	"log"
	"reflect"
	"testing"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func init() {
	log.SetFlags(log.Lshortfile)
}

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

func Test_sortedListToBST(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "1",
			args: args{
				head: MakeListNode(-10, -3, 0, 5, 9),
			},
			want: MakeTree([]int{0, -3, 9, -10, Null, 5}),
		},
		{
			name: "2",
			args: args{
				head: MakeListNode(1, 3),
			},
			want: MakeTree([]int{3, 1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortedListToBST(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortedListToBST(%v) = %v, want %v", tt.args.head, got, tt.want)
			}
		})
	}
}

func Test_isBalanced(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				root: MakeTree2(3, 9, 20, null, null, 15, 7),
			},
			want: true,
		},
		{
			name: "2",
			args: args{
				root: MakeTree2(1, 2, 2, 3, 3, null, null, 4, 4),
			},
			want: false,
		},
		{
			name: "3",
			args: args{
				root: MakeTree2(),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isBalanced(tt.args.root); got != tt.want {
				t.Errorf("isBalanced(%v) = %v, want %v", tt.args.root, got, tt.want)
			}
		})
	}
}
