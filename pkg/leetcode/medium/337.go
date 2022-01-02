package medium

/*
	337. House Robber III
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// https://books.halfrost.com/leetcode/ChapterFour/0300~0399/0337.House-Robber-III/
func robIII(root *TreeNode) int {

	r1, r2 := robDFS(root)

	return max(r1, r2)
}

func robDFS(root *TreeNode) (a, b int) {
	if root == nil {
		return 0, 0
	}

	l1, l2 := robDFS(root.Left)
	r1, r2 := robDFS(root.Right)

	return max(l1, l2) + max(r1, r2), root.Val + l1 + r1
}
