package easy

import (
	"reflect"
	"testing"
)

func Test_lowestCommonAncestor(t *testing.T) {
	type args struct {
		root *TreeNode
		p    *TreeNode
		q    *TreeNode
	}

	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "",
			args: args{
				root: MakeTree(6, 2, 8, 0, 4, 7, 9, null, null, 3, 5),
				p:    MakeTree(2),
				q:    MakeTree(8),
			},
			want: MakeTree(6, 2, 8, 0, 4, 7, 9, null, null, 3, 5),
		},
		{
			name: "",
			args: args{
				root: MakeTree(6, 2, 8, 0, 4, 7, 9, null, null, 3, 5),
				p:    MakeTree(2),
				q:    MakeTree(4),
			},
			want: MakeTree(2, 0, 4, null, null, 3, 5),
		},
		{
			name: "",
			args: args{
				root: MakeTree(2, 1),
				p:    MakeTree(2),
				q:    MakeTree(1),
			},
			want: MakeTree(2, 1),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lowestCommonAncestor(tt.args.root, tt.args.p, tt.args.q); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lowestCommonAncestor() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
