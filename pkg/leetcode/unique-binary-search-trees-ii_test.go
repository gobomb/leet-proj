package leetcode

import (
	"log"
	"reflect"
	"testing"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func Test_generateTrees(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []*TreeNode
	}{
		{
			name: "1",
			args: args{
				n: 3,
			},
			// want: []*TreeNode{MakeTree([]int{1, Null, 2, Null, 3}), MakeTree2([]int{1, Null, 3, 2}),
			// 	MakeTree2([]int{2, 1, 3}), MakeTree2([]int{3, 1, Null, Null, 2}), MakeTree2([]int{3, 2, Null, 1})},
		},
	}
	for _, tt := range tests {
		// for i:=range tt.want{
		// log.Printf("%+v\n",	tt.want[i])
		// }
		// continue
		t.Run(tt.name, func(t *testing.T) {
			if got := generateTrees(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateTrees(%v) = %v, want %+v", tt.args.n, got, tt.want)
			}
		})
	}
}
