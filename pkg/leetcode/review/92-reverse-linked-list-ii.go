package review

// 92. reverse-linked-list-ii

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummyHead := &ListNode{
		Next: head,
	}

	pre := dummyHead

	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}

	cur := pre.Next

	for i := 0; i < right-left; i++ {
		next := cur.Next

		cur.Next = next.Next

		next.Next = pre.Next

		pre.Next = next
	}

	return dummyHead.Next
}

func reverseBetween2(head *ListNode, left int, right int) *ListNode {
	if right <= left {
		return head
	}

	count := 0
	var preHead, end *ListNode

	preHead = &ListNode{
		Next: head,
	}

	h := head
	for {
		count++
		if count == left-1 {
			preHead = h
		}

		if count == right {
			end = h
			break
		}
		h = h.Next
	}

	endNext := end.Next

	reverseBetween4(preHead.Next, end)

	if preHead.Next == head {
		head = end
	}

	preHead.Next.Next = endNext
	preHead.Next = end

	return head
}

func reverseBetween4(head, tail *ListNode) {
	if head == tail {
		return
	}

	reverseBetween4(head.Next, tail)

	head.Next.Next = head
	head.Next = nil
}
