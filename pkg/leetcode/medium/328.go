package medium

/*
	328. Odd Even Linked List
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/*
 	Runtime: 13 ms, faster than 5.23% of Go online submissions for Odd Even Linked List.
	Memory Usage: 3.6 MB, less than 10.09% of Go online submissions for Odd Even Linked List.
*/
func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	h := head
	n := head.Next
	odd := &ListNode{}
	temp := odd

	for {
		if n == nil {
			break
		}

		odd.Next = n
		odd = odd.Next

		h.Next = n.Next
		if h.Next == nil {
			break
		}

		h = h.Next
		n = h.Next
	}

	odd.Next = nil
	h.Next = temp.Next

	return head
}
