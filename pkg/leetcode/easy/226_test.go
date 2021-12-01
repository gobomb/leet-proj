package easy

import (
	"reflect"
	"testing"
)

func Test_invertTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}

	getTests := func() []struct {
		name string
		args args
		want *TreeNode
	} {
		return []struct {
			name string
			args args
			want *TreeNode
		}{
			{
				name: "",
				args: args{
					root: MakeTree(4, 2, 7, 1, 3, 6, 9),
				},
				want: MakeTree(4, 7, 2, 9, 6, 3, 1),
			},
			{
				name: "",
				args: args{
					root: MakeTree(2, 1, 3),
				},
				want: MakeTree(2, 3, 1),
			},
			{
				name: "",
				args: args{
					root: MakeTree(2, null, 3),
				},
				want: MakeTree(2, 3, null),
			},
			{
				name: "",
				args: args{
					root: MakeTree(),
				},
				want: MakeTree(),
			},
		}
	}

	toTest := []func(*TreeNode) *TreeNode{invertTree, invertTree2}

	for _, f := range toTest {
		for _, tt := range getTests() {
			t.Run(tt.name, func(t *testing.T) {
				if got := f(tt.args.root); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("%v() = %+v, want %+v", funcName(f), got, tt.want)
				}
			})
		}
	}
}
