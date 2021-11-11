package easy

/*
	160. Intersection of Two Linked Lists

*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	h1 := headA
	h2 := headB

	for h1 != h2 {
		if h1 == nil {
			h1 = headB
		} else {
			h1 = h1.Next
		}

		if h2 == nil {
			h2 = headA
		} else {
			h2 = h2.Next
		}
	}
	return h1
}
