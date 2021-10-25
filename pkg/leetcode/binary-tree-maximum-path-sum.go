package leetcode

import "math"

/*
	124. Binary Tree Maximum Path Sum
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxPathSum(root *TreeNode) int {
	rs := math.MinInt
	getMaxPath(root, &rs)
	return rs
}

// https://books.halfrost.com/leetcode/ChapterFour/0100~0199/0124.Binary-Tree-Maximum-Path-Sum/
func getMaxPath(root *TreeNode, rs *int) int {
	if root == nil {
		return math.MinInt
	}
	// 求左/右不穿过根节点的最大值
	left := getMaxPath(root.Left, rs)
	right := getMaxPath(root.Right, rs)

	// 求当前不穿过根节点的最大值 max（左+根，右+根，根）
	curMax := max(max(left+root.Val, right+root.Val), root.Val)
	// max（全局最大值，当前不穿过根节点的最大值，穿过当前根节点的最大值）
	*rs = max(*rs, max(curMax, left+right+root.Val))
	// 穿过当前根节点的最大值不需要返回
	return curMax
}
