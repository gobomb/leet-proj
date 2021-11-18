package medium

import (
	"testing"
)

func TestConstructorBSTIterator(t *testing.T) {
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
				root: MakeTree(7, 3, 15, null, null, 9, 20),
			},
			want: []int{3, 7, 9, 15, 20},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bSTIterator := ConstructorBSTIterator(tt.args.root)

			for i := 0; i < len(tt.want); i++ {
				if !bSTIterator.HasNext() {
					break
				}

				if bSTIterator.Next() != tt.want[i] {
					t.Fail()
				}
			}
		})
	}
}
