package medium

import (
	"justest/pkg/ds/listnode"
	"justest/pkg/ds/tree"
	"justest/pkg/utils"
	"log"
)

var (
	max             = utils.Max
	min             = utils.Min
	Null            = tree.Null
	null            = tree.Null
	MakeListNode    = listnode.MakeListNode
	LNDeepCopy      = listnode.LNDeepCopy
	isNull          = tree.IsNull
	funcName        = utils.FuncName
	ShouldEqualDiff = utils.ShouldEqualDiff
)

type (
	ListNode     = listnode.ListNode
	LNDeepCopyer = listnode.LNDeepCopyer
	TreeNode     = tree.TreeNode
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func MakeTree(arg ...int) *TreeNode {
	return tree.MakeTree(arg)
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}
