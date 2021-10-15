package leetcode

import "testing"

func Test_sumOfLeftLeaves(t *testing.T) {
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
				root: MakeTree2(3, 9, 20, null, null, 15, 7),
			},
			want: 24,
		},
		{
			name: "2",
			args: args{
				root: MakeTree2(1),
			},
			want: 0,
		},
		{
			name: "3",
			args: args{
				root: MakeTree2(3, 9),
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumOfLeftLeaves(tt.args.root); got != tt.want {
				t.Errorf("sumOfLeftLeaves() = %v, want %v", got, tt.want)
			}
		})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumOfLeftLeaves3(tt.args.root); got != tt.want {
				t.Errorf("sumOfLeftLeaves3() = %v, want %v", got, tt.want)
			}
		})
	}
}
