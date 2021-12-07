package easy

/*
	1290. Convert Binary Number in a Linked List to Integer
*/

func getDecimalValue(h *ListNode) int {
	b := 0

	for h != nil {
		b += h.Val
		h = h.Next
		b <<= 1
	}

	return b >> 1
}
