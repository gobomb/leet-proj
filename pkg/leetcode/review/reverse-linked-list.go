package review

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 206. Reverse Linked List

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

func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	last := reverseList2(head.Next)

	head.Next.Next = head
	head.Next = nil
	return last
}
