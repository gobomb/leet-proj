package medium

import (
	"reflect"
	"testing"
)

func Test_deleteNode(t *testing.T) {
	type args struct {
		root *TreeNode
		key  int
	}

	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "",
			args: args{
				root: MakeTree(5, 3, 6, 2, 4, null, 7),
				key:  3,
			},
			want: MakeTree(5, 4, 6, 2, null, null, 7),
		},
		{
			name: "",
			args: args{
				root: MakeTree(5, 3, 6, 2, 4, null, 7),
				key:  0,
			},
			want: MakeTree(5, 3, 6, 2, 4, null, 7),
		},
		{
			name: "",
			args: args{
				root: MakeTree(0),
				key:  0,
			},
			want: MakeTree(),
		},
		{
			name: "",
			args: args{
				root: MakeTree(50, 30, 70, null, 40, 60, 80),
				key:  50,
			},
			want: MakeTree(60, 30, 70, null, 40, null, 80),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("args %+v", tt.args.root)
			if got := deleteNode(tt.args.root, tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteNode() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
