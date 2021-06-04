package leetcode

import (
	"log"
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
	}{
		{
			name: "1",
			args: args{
				root: MakeTree([]int{3, Null, 1, Null, Null, 2}),
			},
		},
		{
			name: "2",
			args: args{
				root: MakeTree([]int{10, 40, 60, 19, 45, 55, 65, 50, 35}),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			log.Printf("%+v\n", tt.args.root)

			recoverTree(tt.args.root)
			log.Printf("%+v\n", tt.args.root)

		})
	}
}
