package tree

import (
	"justest/pkg/utils"
	"math"
)

var Null = math.MaxInt64
var max = utils.Max

func isNull(val int) bool {
	return val == Null
}

//   Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) DeepCopy() *TreeNode {
	var tt = &TreeNode{}
	if t != nil {
		tt.Val = t.Val
		tt.Left = t.Left.DeepCopy()
		tt.Right = t.Right.DeepCopy()
		return tt
	} else {
		return t
	}
}

func (t *TreeNode) Insert(v int) *TreeNode {
	if t == nil {
		t = &TreeNode{
			Val: v,
		}
		return t
	}
	if v > t.Val {
		if t.Right != nil {
			t.Right.Insert(v)
		} else {
			t.Right = &TreeNode{
				Val: v,
			}
		}
	} else {
		if t.Left != nil {
			t.Left.Insert(v)
		} else {
			t.Left = &TreeNode{
				Val: v,
			}
		}
	}
	return t
}

func (t *TreeNode) Find(v int) *TreeNode {
	if t == nil {
		return nil
	}
	if t.Val == v {
		return t
	}
	if v > t.Val {
		return t.Right.Find(v)
	} else {
		return t.Left.Find(v)
	}
}

func (t *TreeNode) FindMax() *TreeNode {
	if t == nil {
		return nil
	}
	if t.Right != nil {
		return t.Right.FindMax()
	} else {
		return t
	}
}
func (t *TreeNode) FindMin() *TreeNode {
	if t == nil {
		return nil
	}
	if t.Left != nil {
		return t.Left.FindMin()
	} else {
		return t
	}
}

func (t *TreeNode) Remove(v int) *TreeNode {
	if t == nil {
		return t
	}
	if v > t.Val {
		t.Right = t.Right.Remove(v)
	} else if v < t.Val {
		t.Left = t.Left.Remove(v)
	} else if t.Left != nil && t.Right != nil {
		m := t.Right.FindMin()
		t.Val = m.Val
		t.Right = t.Right.Remove(v)
	} else {
		if t.Left == nil {
			t = t.Right
		} else if t.Right == nil {
			t = t.Left
		}
	}

	return t

}

func getTreeHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(getTreeHeight(root.Left), getTreeHeight(root.Right))
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

type Tree interface {
	Insert(v int)
	Heigh() int
}

type BinSearchTree interface {
	FindMax() BinSearchTree
	FindMin() BinSearchTree
}
