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
