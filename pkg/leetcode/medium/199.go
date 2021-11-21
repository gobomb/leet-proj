package medium

/*
	199. Binary Tree Right Side View
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
	res := []int{}
	dfsRightSideView(root, &res, 1)
	return res
}

func dfsRightSideView(root *TreeNode, res *[]int, depth int) {
	if root == nil {
		return
	}

	if len(*res) < depth {
		*res = append(*res, root.Val)
	}

	dfsRightSideView(root.Right, res, depth+1)
	dfsRightSideView(root.Left, res, depth+1)
}
