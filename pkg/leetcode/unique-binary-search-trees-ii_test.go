package leetcode

import (
	"log"
	"reflect"
	"testing"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func Test_generateTrees(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []*TreeNode
	}{
		{
			name: "1",
			args: args{
				n: 3,
			},
			want: []*TreeNode{MakeTree([]int{1, Null, 2, Null, Null, Null, 3}),
				MakeTree([]int{1, Null, 3, Null, Null, 2}),
				MakeTree([]int{2, 1, 3}),
				MakeTree([]int{3, 1, Null, Null, 2}),
				MakeTree([]int{3, 2, Null, 1}),
			},
		},
	}
	for _, tt := range tests {
		// for i:=range tt.want{
		// log.Printf("%+v\n",	tt.want[i])
		// }
		// continue
		t.Run(tt.name, func(t *testing.T) {
			if got := generateTrees(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateTrees(%v) = %+v, want %+v", tt.args.n, got, tt.want)
			}
		})
	}
}

func Test_numTrees(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				n: 3,
			},
			want: 5,
		},
		{
			name: "2",
			args: args{
				n: 1,
			},
			want: 1,
		},
		{
			name: "3",
			args: args{
				n: 2,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numTrees(tt.args.n); got != tt.want {
				t.Errorf("numTrees(%v) = %v, want %v", tt.args.n, got, tt.want)
			}
		})
	}
}
