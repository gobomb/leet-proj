package easy

/*
	226. Invert Binary Tree
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	queue := []*TreeNode{root}

	for len(queue) != 0 {
		r := queue[0]
		queue = queue[1:]

		if r == nil {
			continue
		}

		r.Right, r.Left = r.Left, r.Right
		queue = append(queue, r.Left, r.Right)
	}

	return root
}

// https://leetcode.com/problems/invert-binary-tree/discuss/1443168/Golang-Recursive-sol'n-faster-than-100-using-multi-assignment-swap
func invertTree2(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	root.Left, root.Right = invertTree2(root.Right), invertTree2(root.Left)

	return root
}
