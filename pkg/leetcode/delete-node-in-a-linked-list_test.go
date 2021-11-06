package leetcode

import (
	"reflect"
	"testing"
)

func Test_deleteNode(t *testing.T) {
	nodes := []*ListNode{}
	nodes = append(nodes, MakeListNode(4, 5, 1, 9))
	nodes = append(nodes, MakeListNode(4, 5, 1, 9))
	nodes = append(nodes, MakeListNode(1, 2, 3, 4))
	nodes = append(nodes, MakeListNode(0, 1))
	nodes = append(nodes, MakeListNode(-3, 5, -99))

	wants := []*ListNode{}
	wants = append(wants, MakeListNode(4, 1, 9))
	wants = append(wants, MakeListNode(4, 5, 9))
	wants = append(wants, MakeListNode(1, 2, 4))
	wants = append(wants, MakeListNode(1))
	wants = append(wants, MakeListNode(5, -99))

	finds := []int{5, 1, 3, 0, -3}

	type args struct {
		node *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{}

	for i, node := range nodes {
		tests = append(tests, struct {
			name string
			args args
			want *ListNode
		}{
			args: args{node: node.FindNode(finds[i])},
			want: wants[i],
		})
	}

	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deleteNode(tt.args.node)
			if got := nodes[i]; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteNode() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
