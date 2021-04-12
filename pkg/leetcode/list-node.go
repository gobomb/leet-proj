package leetcode

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	if l == nil {
		return fmt.Sprintf("(nil)")
	}
	return fmt.Sprintf("%v->%v", l.Val, l.Next)
}

func makeListNode(in []int) *ListNode {
	var last *ListNode
	for i := len(in) - 1; i >= 0; i-- {
		last = &ListNode{
			Val:  in[i],
			Next: last,
		}
	}
	return last
}
