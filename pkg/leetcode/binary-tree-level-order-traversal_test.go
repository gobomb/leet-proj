package leetcode

import (
	"log"
	"reflect"
	"testing"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func Test_levelOrder(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "1",
			args: args{
				root: MakeTree([]int{3, 9, 20, Null, Null, 15, 7}),
			},
			want: [][]int{{3}, {9, 20}, {15, 7}},
		},
		{
			name: "2",
			args: args{
				root: MakeTree([]int{3}),
			},
			want: [][]int{{3}},
		},
		{
			name: "3",
			args: args{
				root: nil,
			},
			want: [][]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := levelOrder(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrder(%+v) = %v, want %v", tt.args.root, got, tt.want)
			}
		})
	}
}
