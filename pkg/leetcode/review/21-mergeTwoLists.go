package review

/*
	21. Merge Two Sorted Lists

	4min
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
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

func MergeTwoLists(l1 *ListNode, l2 *ListNode) (l3 *ListNode) {
	return mergeTwoLists(l1, l2)
}
