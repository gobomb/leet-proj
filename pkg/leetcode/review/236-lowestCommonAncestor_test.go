package review

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
				root: MakeTree(3, 5, 1, 6, 2, 0, 8, null, null, 7, 4),
				p:    MakeTree(5),
				q:    MakeTree(1),
			},
			want: MakeTree(3, 5, 1, 6, 2, 0, 8, null, null, 7, 4),
		},
		{
			name: "",
			args: args{
				root: MakeTree(3, 5, 1, 6, 2, 0, 8, null, null, 7, 4),
				p:    MakeTree(5),
				q:    MakeTree(4),
			},
			want: MakeTree(5, 6, 2, null, null, 7, 4),
		},
		{
			name: "",
			args: args{
				root: MakeTree(2, 1),
				p:    MakeTree(2),
				q:    MakeTree(1),
			},
			want: MakeTree(2, 1),
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lowestCommonAncestor(tt.args.root, tt.args.p, tt.args.q); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lowestCommonAncestor(%+v) = %+v, want %+v",tt.args.root, got, tt.want)
			}
		})
	}
}
