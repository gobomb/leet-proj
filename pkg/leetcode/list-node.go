package leetcode

import (
	"fmt"
	"io"
	"sync"
)

// var VmFree = &Vm{}

var vmFree = sync.Pool{
	New: func() interface{} {
		vm := make(visitMap)
		return vm
	},
}

type Looper interface {
	Loop(o, n *ListNode) bool
}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) FindNode(v int) *ListNode {
	for {
		if l.Val == v {
			return l
		}
		l = l.Next
		if l == nil {
			return nil
		}
	}
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
	case 'l':
		fmt.Fprintf(s, "%v", ListLen(l))
		return
	case 'p':
		_, err := io.WriteString(s, fmt.Sprintf("%p", l))
		if err != nil {
			panic(err)
		}
	}
}

func MakeListNode(in ...int) *ListNode {
	return makeListNode(in)
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
		// make(visitMap),
		vmFree.Get().(visitMap),
	}
}

func LNDeepCopy(o *ListNode) (n *ListNode) {
	ldc := newLNDeepCopyer()
	defer func() {
		for k := range ldc.vm {
			delete(ldc.vm, k)
		}

		vmFree.Put(ldc.vm)
	}()
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
	len      int
	// isLoop   bool
}

func ListStringer(l *ListNode) fmt.Stringer {
	return &ListPrinter{
		visitMap: vmFree.Get().(visitMap),
		head:     l,
		cur:      l,
	}
}

func ListLen(l *ListNode) int {
	lp := &ListPrinter{
		visitMap: vmFree.Get().(visitMap),
		head:     l,
		cur:      l,
	}
	defer func() {
		for k := range lp.visitMap {
			delete(lp.visitMap, k)
		}

		vmFree.Put(lp.visitMap)
	}()

	return lp.Len()
}

func (l *ListPrinter) Len() int {
	cur := l.cur
	l.len = 0
	for i := 0; ; i++ {
		if cur == nil {
			return l.len
		}
		if loop := l.Loop(cur, nil); loop {
			return l.len
		}
		l.len++

		cur = cur.Next
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
		defer func() {
			for k := range l.visitMap {
				delete(l.visitMap, k)
			}

			vmFree.Put(l.visitMap)
		}()
		return "(nil)"
	}
	if loop := l.Loop(l.cur, nil); loop {
		for k := range l.visitMap {
			delete(l.visitMap, k)
		}
		return fmt.Sprintf("(loop)%v", l.cur.Val)
	}
	cur := l.cur
	l.cur = l.cur.Next
	n := l
	return fmt.Sprintf("%v->%v", cur.Val, n)
}
