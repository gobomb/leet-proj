package leetcode

import (
	"justest/pkg/ds/listnode"
	"justest/pkg/ds/tree"
	"justest/pkg/utils"
)

var (
	max          = utils.Max
	min          = utils.Min
	MakeTree     = tree.MakeTree
	Null         = tree.Null
	null         = tree.Null
	MakeListNode = listnode.MakeListNode
	LNDeepCopy   = listnode.LNDeepCopy
	funcName     = utils.FuncName
)

type (
	ListNode     = listnode.ListNode
	LNDeepCopyer = listnode.LNDeepCopyer
	TreeNode     = tree.TreeNode
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func MakeTree2(arg ...int) *TreeNode {
	return MakeTree(arg)
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}
