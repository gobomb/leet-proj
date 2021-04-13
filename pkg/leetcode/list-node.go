package leetcode

import (
	"fmt"
	"io"
)

type Looper interface {
	Loop(o, n *ListNode) bool
}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// Ref: github.com/pkg/errors
func (l *ListNode) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			fmt.Fprintf(s, "%+v", ListStringer(l))
			return
		}

		fallthrough
	case 'p':
		io.WriteString(s, fmt.Sprintf("%p", l))
	}
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

type visitMap map[*ListNode]*ListNode

type LNDeepCopyer struct {
	vm visitMap
}

func newLNDeepCopyer() LNDeepCopyer {
	return LNDeepCopyer{
		make(visitMap),
	}
}

func LNDeepCopy(o *ListNode) (n *ListNode) {
	ldc := newLNDeepCopyer()
	n = ldc.lnDeepCopy(o)
	return
}

func (lnc LNDeepCopyer) Loop(o, n *ListNode) bool {
	// fmt.Printf("loop map  %+v\n", lnc.vm)
	_, ok := lnc.vm[o]
	if ok {
		return true
	}
	lnc.vm[o] = n
	return false
}

func (lnc LNDeepCopyer) lnDeepCopy(o *ListNode) (n *ListNode) {
	if o == nil {
		return n
	}
	n = new(ListNode)

	if loop := lnc.Loop(o, n); loop {
		val := lnc.vm[o]
		return val
	}
	n.Val = o.Val
	n.Next = lnc.lnDeepCopy(o.Next)
	return n
}

type ListPrinter struct {
	visitMap visitMap
	head     *ListNode
	cur      *ListNode
}

func ListStringer(l *ListNode) fmt.Stringer {
	return &ListPrinter{
		visitMap: make(visitMap),
		head:     l,
		cur:      l,
	}
}

func (l *ListPrinter) Loop(o, n *ListNode) bool {
	_, ok := l.visitMap[o]
	if ok {
		return true
	}
	l.visitMap[o] = n
	return false
}

func (l *ListPrinter) String() string {
	if l.cur == nil {
		return fmt.Sprintf("(nil)")
	}
	if loop := l.Loop(l.cur, nil); loop {
		return fmt.Sprintf("(loop)%v", l.cur.Val)
	}
	cur := l.cur
	l.cur = l.cur.Next
	n := l
	return fmt.Sprintf("%v->%v", cur.Val, n)
}
