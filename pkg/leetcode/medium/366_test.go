package medium

import (
	"reflect"
	"testing"
)

func Test_find1(t *testing.T) {
	type args struct {
		n *TreeNode
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
					n: MakeTree(1, 2, 3, 4, 5),
				},
				want: [][]int{{4, 5, 3}, {2}, {1}},
			},
		}
	}

	toTest := []func(*TreeNode) [][]int{find1, find2}

	for _, f := range toTest {
		for _, tt := range tests() {
			t.Run(tt.name, func(t *testing.T) {
				if got := f(tt.args.n); !reflect.DeepEqual(got, tt.want) && tt.args.n != nil {
					t.Errorf("%v(%+v) = %+v, want %+v", funcName(f), tt.args, got, tt.want)
				}
			})
		}
	}
}
