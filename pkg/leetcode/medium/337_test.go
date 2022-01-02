package medium

import (
	"testing"
)

func Test_robIII(t *testing.T) {
	type args struct {
		root *TreeNode
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				root: MakeTree(3, 2, 3, null, 3, null, 1),
			},
			want: 7,
		},
		{
			name: "",
			args: args{
				root: MakeTree(3, 4, 5, 1, 3, null, 1),
			},
			want: 9,
		},
		{
			name: "",
			args: args{
				root: MakeTree(4, 1, null, 2, null, null, null, 3),
			},
			want: 7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("robIII(%+v), want %v", tt.args.root, tt.want)

			if got := robIII(tt.args.root); got != tt.want {
				t.Errorf("robIII(%+v) = %v, want %v", tt.args.root, got, tt.want)
			}
		})
	}
}
