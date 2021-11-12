package easy

/*
	203. Remove Linked List Elements

*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	fake := &ListNode{
		Next: head,
	}

	pre := fake
	h := fake.Next

	for h != nil {
		if h.Val == val {
			pre.Next = h.Next
			h = pre.Next
		} else {
			pre = h
			h = h.Next
		}
	}

	return fake.Next
}
