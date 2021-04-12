package leetcode

import (
	"fmt"
	"sync"
)

type Looper interface {
	Loop(o, n *ListNode) bool
}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
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






// FIXME:  不使用公共map，需要考虑是否内存泄漏
var initvm visitMap
var once sync.Once

func init() {
	once.Do(func() {
		initvm = make(visitMap)
	})
}

// func (l *ListNode) Loop(o, n *ListNode) bool {
// 	_, ok := initvm[o]
// 	if ok {
// 		// FIXME:  清空公共map，需要考虑是否内存泄漏
// 		initvm = make(visitMap)
// 		return true
// 	}
// 	initvm[o] = n
// 	return false
// }

func (l *ListNode) String() string {
	if l == nil {
		return fmt.Sprintf("(nil)")
	}
	// if loop := l.Loop(l, nil); loop {
	// 	return fmt.Sprintf("(loop)")
	// }
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
