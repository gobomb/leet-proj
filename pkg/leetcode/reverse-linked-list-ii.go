package leetcode

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	indirect := head
	var count int

	var dummy, ll, rr, llsave *ListNode
	for (indirect) != nil {
		count++
		if count >= left && count <= right {
			tmp := dummy
			dummy = indirect
			indirect = indirect.Next
			dummy.Next = tmp

			if count == left {
				llsave = ll
				rr = dummy
			}
			rr.Next = indirect
		} else {
			ll = indirect
			indirect = indirect.Next
		}
	}

	// log.Printf("%+v\n", llsave)
	// log.Printf("%+v\n", dummy)

	if llsave != nil {
		llsave.Next = dummy
	} else {
		return dummy
	}
	return head
}
