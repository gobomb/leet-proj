package medium

/*
	148. Sort List
*/

// 3340 ms	7.3 MB
func sortList(head *ListNode) *ListNode {
	return insertionSortList(head)
}

// https://books.halfrost.com/leetcode/ChapterFour/0100~0199/0148.Sort-List/
func sortList1(head *ListNode) *ListNode {
	length := 0
	cur := head

	for cur != nil {
		length++

		cur = cur.Next
	}

	if length <= 1 {
		return head
	}

	middleNode := middleNode(head)
	// log.Printf("%+v %+v", head, middleNode)

	// 利用中间节点的前一个节点，来切割左右半部分的连接
	cur = middleNode.Next
	middleNode.Next = nil
	middleNode = cur

	// log.Printf("%+v %+v", head, middleNode)

	left := sortList(head)
	right := sortList(middleNode)

	return mergeTwoLists(left, right)
}

// 返回的是中间节点的前一个节点
func middleNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	p1 := head
	p2 := head

	for p2.Next != nil && p2.Next.Next != nil {
		p1 = p1.Next
		p2 = p2.Next.Next
	}

	return p1
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
