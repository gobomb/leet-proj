package leetcode

import (
	"log"
	"reflect"
	"testing"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func Test_recoverTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "1",
			args: args{
				root: MakeTree([]int{3, Null, 1, Null, Null, 2}),
			},
			want: MakeTree([]int{1, Null, 3, Null, Null, 2}),
		},
		{
			name: "2",
			args: args{
				root: MakeTree([]int{10, 40, 60, 19, 45, 55, 65, 50, 35}),
			},
			want: MakeTree([]int{50, 40, 60, 19, 45, 55, 65, 10, 35}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// log.Printf("%+v\n", tt.args.root)
			recoverTree(tt.args.root)
			if !reflect.DeepEqual(tt.args.root, tt.want) {
				t.Errorf("failed at recoverTree got: %+v, but want %+v\n", tt.args.root, tt.want)
			}
		})
	}
}

func Test_recoverTree2(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "1",
			args: args{
				root: MakeTree([]int{3, Null, 1, Null, Null, 2}),
			},
			want: MakeTree([]int{1, Null, 3, Null, Null, 2}),
		},
		{
			name: "2",
			args: args{
				root: MakeTree([]int{10, 40, 60, 19, 45, 55, 65, 50, 35}),
			},
			want: MakeTree([]int{50, 40, 60, 19, 45, 55, 65, 10, 35}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// log.Printf("%+v\n", tt.args.root)

			recoverTree2(tt.args.root)
			if !reflect.DeepEqual(tt.args.root, tt.want) {
				t.Errorf("failed at recoverTree2 got: %+v, but want %+v\n", tt.args.root, tt.want)
			}
			// log.Printf("%+v\n", tt.args.root)
		})
	}
}
