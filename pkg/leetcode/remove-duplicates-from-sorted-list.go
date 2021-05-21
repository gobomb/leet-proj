package leetcode

func deleteDuplicates(head *ListNode) *ListNode {
	h := head
	l := &ListNode{Val: -101}
	for h != nil {
		if l.Val == h.Val {
			l.Next = h.Next
			h = h.Next

			continue
		}
		l = h
		h = h.Next
	}
	return head
}
