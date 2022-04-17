package review


// 25. reverse K group
func reverseKGroup(head *ListNode, k int) *ListNode {
	newHead := head

	for i := 1; i < k && newHead != nil; i++ {
		newHead = newHead.Next
	}

	if newHead == nil {
		return head
	}

	next := newHead.Next

	reverseLink3(head, newHead)

	// 递归下一组 k 长度的链表
	head.Next = reverseKGroup(next, k)

	return newHead
}

// 迭代，翻转链表
func reverseLink(h *ListNode, e *ListNode) {
	pseudo := &ListNode{}

	end := e.Next

	for h != end {
		tmp := pseudo
		pseudo = h
		h = h.Next
		pseudo.Next = tmp
	}
}

// 递归，翻转链表
func reverseLink3(h *ListNode, e *ListNode) {
	if h == e {
		return
	}

	reverseLink3(h.Next, e)

	h.Next.Next = h

	h.Next = nil
}
