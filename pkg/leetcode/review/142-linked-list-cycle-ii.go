package review

import "log"

// 142. linked-list-cycle-ii

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

// 3 2 0 -4 2 0 -4 2 0 -4

// 3 2 0 -4 2 0 -4 2 0 -4
// 3 0 2 -4 0 2 -4
// 4 3

// a 

// (2n-a)%b = (n-a)%b
// 3, 2, 0, -4

func detectCycle2(head *ListNode) *ListNode {
	a := head
	b := head

	c1 := 0
	c2 := 0
	// return nil

	for {
		if b.Next == nil {
			return nil
		}

		a = a.Next
		b = b.Next.Next

		c1++

		if a == b {
			break
		}
	}

	log.Println(c1, "11111111111111111")

	for {
		if b.Next == nil {
			return nil
		}

		c2++

		a = a.Next
		b = b.Next.Next

		if a == b {
			break
		}
	}

	log.Println(c2)

	a = head

	c3 := 0

	for {
		if c3 == c1-c2 {
			log.Println(c3)

			return a
		}
		a = a.Next
	}

}



// 1 2 3 4 5 6 7 8 9 10 3 4 5 6 7 8 9 10

// x=3 要求的值
// a=8 第一次重合
// c= 16 第二次重合一共的次数

// 1 2 3  4 5 6 7 8 9 10 3 4 5 6 7 8 9
// 3+5+8
// 3+5+4+4

// 1 3 5  7 9 3 5 7 9 3  5 7 9 3 5 7 9



// a+b
// a+b+b

// 1 2 3         2
// 3 4 5 6 7     4

// 1 3           1
// 3 5 7 3 5 7   5



// 7 8 3 4 5 6 7	// 6
// 7 3 5 7 3 5 7





// x+4=x/2+5=6