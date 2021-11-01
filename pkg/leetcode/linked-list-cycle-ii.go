package leetcode

func detectCycle(head *ListNode) *ListNode {
	fast := head
	slow := head

	c := checkCycle(head)
	if c == nil {
		return c
	}
	step := getCycleStep(c)

	for i := 0; i < step; i++ {
		fast = fast.Next
	}
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return fast
}

func checkCycle(head *ListNode) *ListNode {
	fast := head
	slow := head
	if head == nil {
		return nil
	}
	if fast.Next != nil {
		fast = fast.Next.Next
	} else {
		return nil
	}
	for fast != nil {
		if slow == fast {
			return slow
		}
		slow = slow.Next

		if fast.Next != nil {
			fast = fast.Next.Next
			continue
		} else {
			break
		}
	}
	return nil
}

func getCycleStep(h *ListNode) int {
	n := h
	var i = 1
	for n.Next != h {
		n = n.Next
		i++
	}
	return i
}
