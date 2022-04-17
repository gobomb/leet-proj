package review

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

	tests := func() []struct {
		name string
		args args
		want [][]int
	} {
		return []struct {
			name string
			args args
			want [][]int
		}{
			{
				name: "1",
				args: args{
					root: MakeTree(3, 9, 20, Null, Null, 15, 7),
				},
				want: [][]int{{3}, {9, 20}, {15, 7}},
			},
			{
				name: "2",
				args: args{
					root: MakeTree(3),
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
	}
	toTest := []func(*TreeNode) [][]int{levelOrder, levelOrderIter}

	for _, f := range toTest {
		for _, tt := range tests() {
			t.Run(tt.name, func(t *testing.T) {
				if got := f(tt.args.root); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("%v(%+v) = %+v, want %#v", funcName(f), tt.args.root, got, tt.want)
				}
			})
		}
	}
}
