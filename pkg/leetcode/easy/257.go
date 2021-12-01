package easy

import "strconv"

/*
	257. Binary Tree Paths
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
	Runtime: 0 ms, faster than 100.00% of Go online submissions for Binary Tree Paths.
	Memory Usage: 2.4 MB, less than 100.00% of Go online submissions for Binary Tree Paths.
*/
func binaryTreePaths(root *TreeNode) []string {
	rs := []string{}
	binaryTreePaths1(root, "", &rs)

	return rs
}

func binaryTreePaths1(root *TreeNode, ss string, rs *[]string) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		ss = ss + strconv.FormatInt(int64(root.Val), 10)
		(*rs) = append((*rs), ss)

		return
	}

	if root.Left != nil {
		binaryTreePaths1(root.Left, ss+strconv.FormatInt(int64(root.Val), 10)+"->", rs)
	}

	if root.Right != nil {
		binaryTreePaths1(root.Right, ss+strconv.FormatInt(int64(root.Val), 10)+"->", rs)
	}
}
