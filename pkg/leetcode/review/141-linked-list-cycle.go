package review

// 141. linked-list-cycle

func hasCycle(head *ListNode) bool {
	var slow, fast *ListNode

	slow = head
	fast = head

	for {
		if slow == nil || fast == nil {
			return false
		}

		slow = slow.Next

		if fast.Next == nil {
			return false
		}

		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}
}
