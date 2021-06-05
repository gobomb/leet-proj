package leetcode

func preorderTraversal(root *TreeNode) []int {
	var rs []int
	preorderTraversal2(root, &rs)
	return rs
}

func preorderTraversal2(root *TreeNode, rs *[]int) {
	if root == nil {
		return
	}
	*rs = append(*rs, root.Val)
	preorderTraversal2(root.Left, rs)
	preorderTraversal2(root.Right, rs)
}
