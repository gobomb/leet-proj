package medium

import (
	"reflect"
	"testing"
)

func Test_rightSideView(t *testing.T) {
	type args struct {
		root *TreeNode
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "",
			args: args{
				root: MakeTree(1, 2, 3, null, 5, null, 4),
			},
			want: []int{1, 3, 4},
		},
		{
			name: "",
			args: args{
				root: MakeTree(1, null, 3),
			},
			want: []int{1, 3},
		},
		{
			name: "",
			args: args{
				root: MakeTree(),
			},
			want: []int{},
		},
		{
			name: "",
			args: args{
				root: MakeTree(1, 2),
			},
			want: []int{1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rightSideView(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rightSideView() = %v, want %v", got, tt.want)
			}
		})
	}
}
