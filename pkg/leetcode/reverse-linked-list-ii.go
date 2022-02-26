package leetcode

func reverseBetween2(head *ListNode, left int, right int) *ListNode {
	i := head
	c := 0

	// lastSave 存left前一个节点
	// h, t 子链的头和尾
	// last i的前一个节点
	var last, t, h, lastSave *ListNode

	for i != nil {
		c++

		if c == left {
			t = i
			h = i

			// 保存 last
			lastSave = last

			i = t.Next
			continue
		}

		if c > left && c <= right {
			// t 总是不变
			t.Next = i.Next

			// h 每轮都要更新
			oldh := h
			h = i
			h.Next = oldh

			i = t.Next

			continue
		}

		last = i
		i = i.Next
	}

	if left == 1 {
		return h
	}

	lastSave.Next = h

	return head
}
