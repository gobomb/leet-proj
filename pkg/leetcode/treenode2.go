package leetcode

import (
	"container/list"
	"log"
)

type IndexTree struct {
	p *TreeNode
	i int
}

// 非完全序列
func MakeTree2(vals []int) *TreeNode {
	if len(vals) == 0 {
		return nil
	}
	l := listFree.Get().(*list.List)
	defer listFree.Put(l.Init())

	tp := &treePrinter{
		list: l,
	}

	rs := &TreeNode{
		Val: vals[0],
	}
	tp.list.PushBack(&IndexTree{rs, 0})
	log.Printf("\n")
	tp.BuildLoop(vals)
	return rs //makeTree3(l,vals, 0)
}

// todo 队列是 tree
// 队列是下标

func (t *treePrinter) BuildLoop(vs []int) *TreeNode {

	e := t.list.Front()
	if e == nil {
		return nil
	}
	t.list.Remove(e)
	it := e.Value.(*IndexTree)
	if it == nil {
		return nil
	}
	v := it.p
	i := it.i
	if i == len(vs) {
		return nil
	}
	log.Printf("%v %v\n", v.Val, i)

	if i+1 < len(vs) && vs[i+1] != Null {
		v.Left = &TreeNode{
			Val: vs[i+1],
		}
		t.list.PushBack(&IndexTree{
			p: v.Left,
			i: i + 1,
		})
	}
	if i+2 < len(vs) && vs[i+2] != Null {
		v.Right = &TreeNode{
			Val: vs[i+2],
		}
		t.list.PushBack(&IndexTree{
			p: v.Right,
			i: i + 2,
		})
	}

	return t.BuildLoop(vs)
}
