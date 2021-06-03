package leetcode

import (
	"log"
	"testing"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func Test_isValidBST(t *testing.T) {
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
				root: MakeTree([]int{2, 1, 3}),
			},
			want: true,
		},
		{
			name: "2",
			args: args{
				root: MakeTree([]int{5, 1, 4, Null, Null, 3, 6}),
			},
			want: false,
		},
		{
			name: "3",
			args: args{
				root: MakeTree([]int{1, Null, 2, Null, Null, Null, 3}),
			},
			want: true,
		},
		{
			name: "4",
			args: args{
				root: MakeTree([]int{3, 1, Null, Null, 2}),
			},
			want: true,
		},
		{
			name: "5",
			args: args{
				root: MakeTree([]int{21, 20, 29, 10, 30}),
			},
			want: false,
		},
		{
			name: "6",
			args: args{
				root: MakeTree([]int{0}),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidBST(tt.args.root); got != tt.want {
				t.Errorf("isValidBST(%+v) = %v, want %v", tt.args.root, got, tt.want)
			}

			if got := isValidBST3(tt.args.root); got != tt.want {
				t.Errorf("isValidBST3(%+v) = %v, want %v", tt.args.root, got, tt.want)
			}
		})
	}
}
