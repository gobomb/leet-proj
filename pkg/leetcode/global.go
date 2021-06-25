package leetcode

import (
	"justest/pkg/tree"
	"justest/pkg/utils"
)

var max = utils.Max
var min = utils.Min

var MakeTree = tree.MakeTree
var Null = tree.Null
var null = tree.Null

type TreeNode = tree.TreeNode

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func MakeTree2(arg ...int) *TreeNode {
	return MakeTree(arg)
}
