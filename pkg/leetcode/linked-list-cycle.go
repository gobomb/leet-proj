package leetcode

func hasCycle(head *ListNode) bool {
	fast := head
	slow := head
	if head == nil {
		return false
	}
	if fast.Next != nil {
		fast = fast.Next.Next
	} else {
		return false
	}
	for fast != nil {
		if slow == fast {
			return true
		}
		slow = slow.Next

		if fast.Next != nil {
			fast = fast.Next.Next
			continue
		} else {
			break
		}
	}
	return false
}
