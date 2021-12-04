package easy

/*
	234. Palindrome Linked List
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return false
	}

	// find middle node
	mid := head
	n := head.Next

	for mid != nil && n != nil {
		mid = mid.Next

		if n.Next == nil {
			break
		}

		n = n.Next.Next
	}

	// reverse the list after middle node
	reverse := reverseList(mid)

	// compare the left list and the reverse right list
	h := head
	for h != nil && reverse != nil {
		if h.Val != reverse.Val {
			return false
		}
		h = h.Next
		reverse = reverse.Next
	}

	return true
}

func reverseList(head *ListNode) *ListNode {
	indirect := head

	var dummy *ListNode

	for indirect != nil {
		tmp := dummy
		dummy = indirect
		indirect = indirect.Next
		dummy.Next = tmp
	}

	return dummy
}
