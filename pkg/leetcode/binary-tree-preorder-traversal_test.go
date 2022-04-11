package leetcode

import (
	"reflect"
	"testing"
)

func Test_preorderTraversal(t *testing.T) {
	type args struct {
		root *TreeNode
	}

	tests := func() []struct {
		name string
		args args
		want []int
	} {
		return []struct {
			name string
			args args
			want []int
		}{
			{
				name: "",
				args: args{
					root: MakeTree2(2, 3, 4, 5),
				},
				want: []int{2, 3, 5, 4},
			},
			{
				name: "",
				args: args{
					root: MakeTree2(2, 3, Null, 4, 5),
				},
				want: []int{2, 3, 4, 5},
			},
		}
	}

	toTest := []func(*TreeNode) []int{preorderTraversal, preorderTraversal3}

	for _, f := range toTest {
		for _, tt := range tests() {
			t.Run(tt.name, func(t *testing.T) {
				if got := f(tt.args.root); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("%v(%+v) = %+v, want %+v", funcName(f), tt.args, got, tt.want)
				}
			})
		}
	}
}
