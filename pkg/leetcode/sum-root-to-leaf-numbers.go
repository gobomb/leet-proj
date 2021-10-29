package leetcode

/*
	129. Sum Root to Leaf Numbers
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return sumNumbersFind(root, 0)
}

func sumNumbersFind(root *TreeNode, num int) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return num*10 + root.Val
	}

	num = num*10 + root.Val
	numL := sumNumbersFind(root.Left, num)
	numR := sumNumbersFind(root.Right, num)
	return numL + numR
}
