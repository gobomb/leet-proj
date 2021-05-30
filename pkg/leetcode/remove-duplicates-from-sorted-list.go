package leetcode

import "fmt"

func deleteDuplicates(head *ListNode) *ListNode {
	h := head
	l := &ListNode{Val: -101}
	for h != nil {
		if l.Val == h.Val {
			l.Next = h.Next
			h = h.Next

			continue
		}
		l = h
		h = h.Next
	}
	return head
}

func deleteDuplicates2(head *ListNode) *ListNode {
	h := head
	for h.Next != nil {
		if h.Val == h.Next.Val {
			h.Next = h.Next.Next
			continue
		}
		h = h.Next
	}
	return head
}

func deleteDuplicatesII(head *ListNode) *ListNode {
	h := head
	l := &ListNode{Val: -101}
	saved := make(map[int]struct{})
	for h != nil {
		if l.Val == h.Val {
			l.Next = h.Next
			h = h.Next
			saved[l.Val] = struct{}{}
			continue
		}
		l = h
		h = h.Next
	}
	h = head
	for h != nil {
		if _, ok := saved[h.Val]; ok {
			if h == head {
				head = h.Next
				l = h
				h = h.Next
				continue
			}
			l.Next = h.Next
			h = h.Next
		} else {
			l = h
			h = h.Next
		}
	}
	return head
}

// https://books.halfrost.com/leetcode/ChapterFour/0001~0099/0082.Remove-Duplicates-from-Sorted-List-II/
func deleteDuplicatesII2(head *ListNode) *ListNode {
	newHead := &ListNode{Next: head, Val: -999999}
	cur := newHead
	last := newHead
	front := head
	for front.Next != nil {
		if front.Val == cur.Val {
			front = front.Next
			continue
		} else {
			if cur.Next != front {
				last.Next = front
				if front.Next != nil && front.Next.Val != front.Val {
					last = front
				}
				cur = front
				front = front.Next
				if front == nil {
					fmt.Println("ss")
				}
			} else {
				front = front.Next
				last = cur
				cur = cur.Next
			}
		}
	}

	if front.Val == cur.Val {
		last.Next = nil
	} else {
		if cur.Next != front {
			last.Next = front
		}
	}

	return newHead.Next
}
