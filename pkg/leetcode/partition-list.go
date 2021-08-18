package leetcode

func partition(head *ListNode, x int) *ListNode {
	var pDummy = &ListNode{}
	dummy := pDummy

	dHead := &ListNode{}
	dHead.Next = head

	p := head
	last := dHead

	for p != nil {
		if p.Val >= x {
			last.Next = p.Next
			p.Next = nil
			pDummy.Next = p
			pDummy = pDummy.Next
			p = last.Next
		} else {
			last = p
			p = p.Next
		}
	}

	last.Next = dummy.Next
	return dHead.Next
}
