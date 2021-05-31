package leetcode

import (
	"container/list"
	"fmt"
	"io"
	"log"
	"math"
	"sync"
)

var listFree = sync.Pool{
	New: func() interface{} {
		vm := list.New()
		return vm
	},
}

var Null = math.MaxInt64

func isNull(val int) bool {
	return val == Null
}

//   Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type treePrinter struct {
	list  *list.List
	root  *TreeNode
	count int
}

// TODO: 按树型打印一颗二叉树
// Ref: github.com/pkg/errors
func (l *TreeNode) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			fmt.Fprintf(s, "%+v", TreeStringer(l))
			return
		}
		fallthrough
	case 'l':
		// fmt.Fprintf(s, "%v", ListLen(l))
		return
	case 'p':
		io.WriteString(s, fmt.Sprintf("%p", l))
	}
}

func TreeStringer(l *TreeNode) fmt.Stringer {
	tp := &treePrinter{
		list: listFree.Get().(*list.List).Init(),
		root: l,
	}

	return tp
}

func (t *treePrinter) String() string {
	defer func() {
		listFree.Put(t.list.Init())
	}()

	if t.root != nil {
		t.list.PushBack(t.root)
		t.count = 0
	}

	// log.Printf("%v\n", t.root.Val)
	return t.StringLoop("")
}

const (
	gap     = "\t"
	newline = "\n"
)

func (t *treePrinter) StringLoop(s string) string {
	e := t.list.Front()
	if e == nil {
		// log.Printf("%v\n", "e nil")

		return s
	}
	t.list.Remove(e)
	v := e.Value.(*TreeNode)
	if v == nil {
		// log.Printf("%v\n", "v nil")
		if t.count&(t.count+1) == 0 {
			s = fmt.Sprintf("%s%v%v", s, gap, newline)
		}
		t.count++

		s = fmt.Sprintf("%s%v%v", s, gap, "Null")
		// t.list.PushBack(new(TreeNode))
		// t.list.PushBack(new(TreeNode))
		return t.StringLoop(s)
	}

	log.Printf("val %v\n", v.Val)
	if t.count&(t.count+1) == 0 {
		s = fmt.Sprintf("%s%v%v", s, gap, newline)
	}
	t.count++

	s = fmt.Sprintf("%s%v%v", s, gap, v.Val)
	t.list.PushBack(v.Left)
	t.list.PushBack(v.Right)

	return t.StringLoop(s) //fmt.Sprintf("%s %v ", s, )
}

// 完全序列
func MakeTree(vals []int) *TreeNode {
	return makeTree(vals, 0)
}

func makeTree(vals []int, i int) *TreeNode {
	if i >= len(vals) || isNull(vals[i]) {
		return nil
	}
	// log.Printf("%v\n", vals[i])

	return &TreeNode{
		Val:   vals[i],
		Left:  makeTree(vals, i*2+1),
		Right: makeTree(vals, i*2+2),
	}
}
