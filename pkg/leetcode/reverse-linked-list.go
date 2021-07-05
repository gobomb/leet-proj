package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
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
