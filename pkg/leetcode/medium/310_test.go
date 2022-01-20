package medium

import (
	"reflect"
	"testing"
)

func Test_findMinHeightTrees(t *testing.T) {
	type args struct {
		n     int
		edges [][]int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "",
			args: args{
				n:     4,
				edges: [][]int{{1, 0}, {1, 2}, {1, 3}},
			},
			want: []int{1},
		},
		{
			name: "",
			args: args{
				n:     1,
				edges: [][]int{},
			},
			want: []int{0},
		},
		{
			name: "",
			args: args{
				n:     2,
				edges: [][]int{{1, 0}},
			},
			want: []int{1, 0},
		},
		{
			name: "",
			args: args{
				n:     3,
				edges: [][]int{{1, 0}, {1, 2}},
			},
			want: []int{1},
		},
		{
			name: "",
			args: args{
				n:     6,
				edges: [][]int{{3, 0}, {3, 1}, {3, 2}, {3, 4}, {5, 4}},
			},
			want: []int{4, 3},
		},
		{
			name: "",
			args: args{
				n:     10,
				edges: [][]int{{1, 0}, {0, 5}, {0, 6}, {6, 7}, {1, 2}, {2, 4}, {2, 3}, {3, 9}, {4, 8}},
			},
			want: []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMinHeightTrees(tt.args.n, tt.args.edges); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findMinHeightTrees() = %v, want %v", got, tt.want)
			}
		})
	}
}
