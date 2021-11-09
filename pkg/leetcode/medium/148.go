package medium

/*
	148. Sort List
*/

// 3340 ms	7.3 MB
func sortList(head *ListNode) *ListNode {
	return insertionSortList(head)
}

// 36 ms	9 MB
func sortListMerge(head *ListNode) *ListNode {
	return partion(head, nil)
}

func partion(head, end *ListNode) *ListNode {
	// log.Printf("%+v %+v", head, end)
	if head == end || head.Next == end {
		if head != nil {
			head.Next = nil
		}

		return head
	}

	if head.Next.Next == end {
		if head.Val > head.Next.Val {
			head, head.Next.Next, head.Next = head.Next, head, nil
		} else {
			head.Next.Next = nil
		}

		return head
	}

	i, j := head, head

	for j != end && j.Next != end {
		i = i.Next
		j = j.Next.Next
	}

	p1 := partion(head, i)
	p2 := partion(i, end)

	return mergeTwoLists(p1, p2)
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) (l3 *ListNode) {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	}

	l2.Next = mergeTwoLists(l1, l2.Next)

	return l2
}
