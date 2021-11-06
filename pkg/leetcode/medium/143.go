package medium

/*
	143. Reorder List
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	p1 := head
	p2 := head

	for p2.Next != nil && p2.Next.Next != nil {
		p1 = p1.Next
		p2 = p2.Next.Next
	}

	preMid := p1
	mid := p1.Next

	r := reverseList(mid)
	preMid.Next = r

	p1 = head
	p2 = preMid

	for p1 != p2 {
		tp1next := p1.Next
		p1.Next = p2.Next

		tmidnext := p2.Next.Next
		p2.Next.Next = tp1next

		p2.Next = tmidnext
		p1 = tp1next
	}
}

func reverseList(head *ListNode) *ListNode {
	indirect := head

	var dummy *ListNode

	for (indirect) != nil {
		tmp := dummy
		dummy = indirect
		indirect = indirect.Next
		dummy.Next = tmp
	}

	return dummy
}
