package medium

import (
	"justest/pkg/tree"
	"justest/pkg/utils"
)

var max = utils.Max
var min = utils.Min

var Null = tree.Null
var null = tree.Null

var isNull = tree.IsNull

type TreeNode = tree.TreeNode

func Abs(x int) int {
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
