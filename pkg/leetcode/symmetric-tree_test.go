package leetcode

import (
	"log"
	"testing"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func Test_isSymmetric(t *testing.T) {
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
				root: MakeTree([]int{1, 2, 2, 3, 4, 4, 3}),
			},
			want: true,
		},
		{
			name: "2",
			args: args{
				root: MakeTree([]int{1, 2, 2, Null, 3, Null, 3}),
			},
			want: false,
		},
		{
			name: "3",
			args: args{
				root: MakeTree([]int{1}),
			},
			want: true,
		},
		{
			name: "4",
			args: args{
				root: MakeTree([]int{1, 2, 2, 2, Null, 2}),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSymmetric(tt.args.root); got != tt.want {
				t.Errorf("isSymmetric(%+v) = %+v, want %v", tt.args.root, got, tt.want)
			}
		})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSymmetricWithQueue(tt.args.root); got != tt.want {
				t.Errorf("isSymmetric(%+v) = %+v, want %v", tt.args.root, got, tt.want)
			}
		})
	}
}
