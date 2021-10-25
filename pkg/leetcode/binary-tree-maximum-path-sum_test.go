package leetcode

import "testing"

func Test_maxPathSum(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				root: MakeTree2(1, 2, 3),
			},
			want: 6,
		},
		{
			name: "2",
			args: args{
				root: MakeTree2(-10, 9, 20, null, null, 15, 7),
			},
			want: 42,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxPathSum(tt.args.root); got != tt.want {
				t.Errorf("maxPathSum(%+v) = %+v, want %v", tt.args.root, got, tt.want)
			}
		})
	}
}
