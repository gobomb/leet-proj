package leetcode

import (
	"log"
	"testing"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func Test_isSameTree(t *testing.T) {
	type args struct {
		p *TreeNode
		q *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				p: MakeTree([]int{1, 2, 3}),
				q: MakeTree([]int{1, 2, 3}),
			},
			want: true,
		},
		{
			name: "2",
			args: args{
				p: MakeTree([]int{1, 2}),
				q: MakeTree([]int{1, Null, 2}),
			},
			want: false,
		},
		{
			name: "3",
			args: args{
				p: MakeTree([]int{1, 2, 1}),
				q: MakeTree([]int{1, 1, 2}),
			},
			want: false,
		},
		{
			name: "4",
			args: args{
				p: MakeTree([]int{1, 1}),
				q: MakeTree([]int{1, Null, 1}),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSameTree(tt.args.p, tt.args.q); got != tt.want {
				t.Errorf("isSameTree(%+v, %+v) = %v, want %v", tt.args.p, tt.args.q, got, tt.want)
			}
		})
	}
}
