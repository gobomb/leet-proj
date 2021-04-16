package leetcode

func reverseKGroup(head *ListNode, k int) *ListNode {
	next := head
	for i := 0; i < k; i++ {
		if next == nil {
			return head
		}

		next = next.Next
	}

	// head 会被改变，成为第 k-1 个 node
	rv := reverse(head, next)
	// fmt.Printf("reverse rv %+v, node %+v head %+v\n", rv, node, head)

	// 递归下一组 k 长度的链表
	head.Next = reverseKGroup(next, k)

	return rv
}

func reverse(first, last *ListNode) *ListNode {
	// prev 是新 head，原始 first 链将逆序并接到 prev 后
	prev := &ListNode{
		Next: last,
	}
	for first != last {
		tmp := prev.Next
		prev.Next = first
		next := first.Next
		first.Next = tmp
		first = next
	}
	return prev.Next
}
