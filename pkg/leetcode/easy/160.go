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

// 剑指offer 37
func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	h1 := headA
	h2 := headB
	l1, l2 := 0, 0

	for h1 != nil {
		h1 = h1.Next
		l1++
	}

	for h2 != nil {
		h2 = h2.Next
		l2++
	}

	h1 = headA
	h2 = headB

	if l1 >= l2 {
		k := l1 - l2

		for i := 0; i < k; i++ {
			h1 = h1.Next
		}
	} else {
		k := l2 - l1

		for i := 0; i < k; i++ {
			h2 = h2.Next
		}
	}

	for h1 != h2 {
		h1 = h1.Next
		h2 = h2.Next
	}

	return h1
}
