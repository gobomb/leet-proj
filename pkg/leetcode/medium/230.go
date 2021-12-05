package medium

/*
	230. Kth Smallest Element in a BST
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
	v := 0
	kthSmallest1(root, &k, &v)
	return v
}

func kthSmallest1(root *TreeNode, k, v *int) {
	if root == nil {
		return
	}
	kthSmallest1(root.Left, k, v)
	*k--
	if *k == 0 {
		*v = root.Val
	}
	kthSmallest1(root.Right, k, v)
}
