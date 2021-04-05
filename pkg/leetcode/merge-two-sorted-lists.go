package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	i := l1
	j := l2
	if i == nil {
		return j
	}
	if j == nil {
		return i
	}
	// for {
	// 	if i.Val <= j.Val && j.Val <= i.Next.Val {
	// 		t := i.Next
	// 		i.Next = j
	// 		j = i.Next.Next
	// 		i.Next.Next = t

	// 		i = i.Next
	// 		continue
	// 	}
	// 	if i.Val > j.Val {

	// 	}
	// }
	return l1
}
