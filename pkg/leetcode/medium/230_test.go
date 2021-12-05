package medium

import (
	"testing"
)

func Test_kthSmallest(t *testing.T) {
	type args struct {
		root *TreeNode
		k    int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				root: MakeTree(3, 1, 4, null, 2),
				k:    1,
			},
			want: 1,
		},
		{
			name: "",
			args: args{
				root: MakeTree(5, 3, 6, 2, 4, null, null, 1),
				k:    4,
			},
			want: 4,
		},
		{
			name: "",
			args: args{
				root: MakeTree(5, 3, 6, 2, 4, null, null, 1),
				k:    3,
			},
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthSmallest(tt.args.root, tt.args.k); got != tt.want {
				t.Errorf("kthSmallest() = %v, want %v", got, tt.want)
			}
		})
	}
}
