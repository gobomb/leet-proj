package review

import (
	"testing"
)

func makeCycle(h *ListNode, pos int) *ListNode {
	// if pos == -1 {
	// 	return h
	// }

	// var p *ListNode
	// s := h
	// c := 0

	// for h.Next != nil {
	// 	h = h.Next
	// 	if c == pos {
	// 		p = h
	// 	}
	// 	c++
	// }

	// h.Next = p
	s, _ := makeCycle2(h, pos)
	return s
}

func Test_hasCycle(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{
				head: makeCycle(MakeListNode(3, 2, 0, -4), 1),
			},
			want: true,
		},
		{
			name: "",
			args: args{
				head: makeCycle(MakeListNode(1, 2), 0),
			},
			want: true,
		},
		{
			name: "",
			args: args{
				head: makeCycle(MakeListNode(1), -1),
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasCycle(tt.args.head); got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
