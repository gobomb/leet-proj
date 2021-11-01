package leetcode

func inorderTraversal(root *TreeNode) []int {
	rs := []int{}
	recursInorderTraversal(root, &rs)
	return rs
}

func recursInorderTraversal(root *TreeNode, rs *[]int) {
	if root == nil {
		return
	}

	recursInorderTraversal(root.Left, rs)
	*rs = append(*rs, root.Val)
	recursInorderTraversal(root.Right, rs)
}
