package leetcode

import (
	"reflect"
	"testing"
)

func Test_findMode(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				root: MakeTree2(1, null, 2, null, null, 2),
			},
			want: []int{2},
		},
		{
			name: "2",
			args: args{
				root: MakeTree2(1),
			},
			want: []int{1},
		},
		{
			name: "3",
			args: args{
				root: MakeTree2(1, 1, 2, 2),
			},
			want: []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMode(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findMode() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
