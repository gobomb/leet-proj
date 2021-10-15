package leetcode

/*
	404. Sum of Left Leaves
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumOfLeftLeaves3(root *TreeNode) int {
	return sumOfLeftLeaves2(root, false)
}

func sumOfLeftLeaves2(root *TreeNode, l bool) int {
	var val int

	if root == nil {
		return val
	}
	if l && root.Left == nil && root.Right == nil {
		return root.Val
	}

	val += sumOfLeftLeaves2(root.Left, true)
	val += sumOfLeftLeaves2(root.Right, false)

	return val
}

func sumOfLeftLeaves(root *TreeNode) int {
	var val int
	if root == nil {
		return val
	}
	if root.Left != nil {
		if root.Left.Left == nil && root.Left.Right == nil {
			val += root.Left.Val
		} else {
			val += sumOfLeftLeaves(root.Left)
		}
	}
	val += sumOfLeftLeaves(root.Right)

	return val
}
