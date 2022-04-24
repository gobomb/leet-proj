package review

import (
	"reflect"
	"testing"
)

func makeCycle2(h *ListNode, pos int) (*ListNode, *ListNode) {
	if pos == -1 {
		return h, nil
	}

	var p *ListNode
	s := h
	c := 0

	for h.Next != nil {
		if c == pos {
			p = h
		}

		h = h.Next
		c++
	}

	h.Next = p
	return s, p
}

func Test_detectCycle(t *testing.T) {
	type args struct {
		head *ListNode
	}
	var heads, wants []*ListNode

	ff := func(a []int, b int) {
		h, w := makeCycle2(MakeListNode(a...), b)

		heads = append(heads, h)
		wants = append(wants, w)
	}

	ff([]int{3, 2, 0, -4}, 1)

	ff([]int{1, 2}, 0)

	ff([]int{-4}, -1)

	ff([]int{-21, 10, 17, 8, 4,
		26, 15, 35, 33, -7,
		-16, 27, -12, 6, 29,
		-12, 15, 9, 20, 14,
		14, 2, 13, -24, 21,
		23, -21, 5}, 24)

	tests := []struct {
		name string
		args args
		want *ListNode
	}{}

	for i := range wants {
		tests = append(tests, struct {
			name string
			args args
			want *ListNode
		}{
			args: args{
				head: heads[i],
			},
			want: wants[i],
		})
	}
	// tail := tests[0].args.head.FindNode(21)
	// tests[0].args.head.FindNode(5).Next = tail
	// tests[0].want = tail
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := detectCycle2(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("detectCycle() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
