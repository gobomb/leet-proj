package easy

func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	rs, dummy := 0, &ListNode{}

	for node := dummy; l1 != nil || l2 != nil || rs > 0; node = node.Next {
		if l1 != nil {
			rs += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			rs += l2.Val
			l2 = l2.Next
		}

		node.Next = &ListNode{Val: rs % 10}
		rs = rs / 10
	}

	return dummy.Next
}
