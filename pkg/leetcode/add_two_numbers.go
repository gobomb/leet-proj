package leetcode

import (
	"log"
)

func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	if l1 == nil || l2 == nil {
		return nil
	}

	head := &ListNode{Val: 0, Next: nil}
	current := head
	carry := 0

	for l1 != nil || l2 != nil {
		var x, y int
		if l1 == nil {
			x = 0
		} else {
			x = l1.Val
		}
		if l2 == nil {
			y = 0
		} else {
			y = l2.Val
		}

		current.Next = &ListNode{Val: (x + y + carry) % 10, Next: nil}
		current = current.Next
		carry = (x + y + carry) / 10

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}

	if carry > 0 {
		current.Next = &ListNode{Val: carry % 10, Next: nil}
	}

	return head.Next
}

func TestAddTwoNumbers() {
	func() {
		v1 := &ListNode{Val: 2, Next: nil}
		v1.Next = &ListNode{Val: 4, Next: nil}
		v1.Next.Next = &ListNode{Val: 3, Next: nil}

		v2 := &ListNode{Val: 5, Next: nil}
		v2.Next = &ListNode{Val: 6, Next: nil}
		v2.Next.Next = &ListNode{Val: 4, Next: nil}
		v3 := addTwoNumbers(v1, v2)

		l := v3
		for l != nil {
			log.Printf("value: %d \n", l.Val)
			l = l.Next
		}
	}()

	func() {
		v1 := &ListNode{Val: 9, Next: nil}
		v1.Next = &ListNode{Val: 9, Next: nil}
		v1.Next.Next = &ListNode{Val: 9, Next: nil}

		v2 := &ListNode{Val: 1, Next: nil}
		v2.Next = &ListNode{Val: 1, Next: nil}
		v2.Next.Next = &ListNode{Val: 1, Next: nil}
		v3 := addTwoNumbers(v1, v2)

		l := v3
		for l != nil {
			log.Printf("value: %d \n", l.Val)
			l = l.Next
		}
	}()
}
