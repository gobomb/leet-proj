package medium

import (
	"reflect"
	"testing"
)

func Test_cloneGraph(t *testing.T) {
	type args struct {
		node *GraphNode
	}

	tests := []struct {
		name string
		args args
		want *GraphNode
	}{
		{
			name: "1",
			args: args{
				node: makeGraphNode([][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}}),
			},
			want: makeGraphNode([][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}}),
		},
		{
			name: "2",
			args: args{
				node: makeGraphNode([][]int{{}}),
			},
			want: makeGraphNode([][]int{{}}),
		},
		{
			name: "3",
			args: args{
				node: makeGraphNode([][]int{}),
			},
			want: makeGraphNode([][]int{}),
		},
		{
			name: "4",
			args: args{
				node: makeGraphNode([][]int{{2}, {1}}),
			},
			want: makeGraphNode([][]int{{2}, {1}}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cloneGraph(tt.args.node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("cloneGraph() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_makeGraphNode(t *testing.T) {
	type args struct {
		input [][]int
		val   int
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "",
			args: args{
				input: [][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}},
				val:   0,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := makeGraphNode(tt.args.input)
			if len(got.Neighbors) != 2 {
				t.Errorf("make graph node failed because num of neighors == %v, want %v", len(got.Neighbors), len(tt.args.input[0]))
			}
		})
	}
}
