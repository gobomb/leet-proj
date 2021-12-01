package medium

import (
	"testing"
)

func Test_countNodes(t *testing.T) {
	type args struct {
		root *TreeNode
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				root: MakeTree(1, 2, 3, 4, 5, 6),
			},
			want: 6,
		},
		{
			name: "",
			args: args{
				root: MakeTree(),
			},
			want: 0,
		},
		{
			name: "",
			args: args{
				root: MakeTree(1),
			},
			want: 1,
		},
	}

	toTest := []func(*TreeNode) int{countNodes, countNodes2}

	for _, f := range toTest {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := f(tt.args.root); got != tt.want {
					t.Errorf("%v(%+v) = %+v, want %+v", tt.args, funcName(f), got, tt.want)
				}
			})
		}
	}
}
