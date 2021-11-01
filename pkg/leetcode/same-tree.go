package leetcode

func isSameTree(root, root2 *TreeNode) bool {
	if root == nil || root2 == nil {
		return root == root2
	}

	if !isSameTree(root.Left, root2.Left) {
		return false
	}
	if root.Val != root2.Val {
		return false
	}
	if !isSameTree(root.Right, root2.Right) {
		return false
	}
	return true
}
