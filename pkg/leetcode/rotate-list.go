package leetcode

func rotateRight(head *ListNode, k int) *ListNode {
	len := LenOfList(head)
	if len == 0 || len == 1 {
		return head
	}
	nk := k % len
	if nk == 0 {
		return head
	}
	c := head
	for i := 1; i <= len-nk-1; i++ {
		c = c.Next
	}
	newh := c.Next
	c.Next = nil

	l := newh
	for {
		if l.Next == nil {
			l.Next = head
			// log.Printf("%+v %v\n", newh, nk)
			return newh
		}
		l = l.Next
	}
}

func LenOfList(ln *ListNode) int {
	cur := ln
	l := 0
	for cur != nil {
		l++
		cur = cur.Next
	}
	return l
}
