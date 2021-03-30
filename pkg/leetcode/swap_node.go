package leetcode

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func testSwapPairs() {
	tests := []*ListNode{
		&ListNode{
			1,
			&ListNode{
				2,
				&ListNode{
					3,
					&ListNode{
						4,
						nil,
					},
				},
			},
		},
		nil,
		{
			1,
			nil,
		},
		{
			1,
			&ListNode{
				2,
				&ListNode{
					3,
					nil,
				},
			},
		},
	}
	for i := range tests {
		println("before:")
		printList(tests[i])
		r := swapPairs(tests[i])
		println("after:")
		printList(r)
	}
}

func printList(h *ListNode) {
	if h == nil {
		return
	}
	printListNode(h)
	println()
}

func printListNode(h *ListNode) {
	fmt.Printf("print: %+v\n", h.Val)
	if h.Next != nil {
		printListNode(h.Next)
	} else {
		return
	}
}

func swapPairs(head *ListNode) *ListNode {
	h := head
	last := &ListNode{}
	for {
		if h == nil || h.Next == nil {
			return head
		}

		temp := h.Next
		h.Next = temp.Next
		temp.Next = h
		h = temp
		last.Next = h
		if h.Next == head {
			head = h
		}

		if h.Next.Next != nil {
			last = h.Next
			h = h.Next.Next
		} else {
			return head
		}

	}
}
