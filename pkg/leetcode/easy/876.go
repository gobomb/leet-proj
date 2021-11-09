package easy

/*
	876. Middle of the Linked List
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func middleNode(head *ListNode) *ListNode {
	i, j := head, head

	for j != nil && j.Next != nil {
		i = i.Next
		j = j.Next.Next
	}

	return i
}
