package leetcode

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return root
	}
	if val > root.Val {
		return searchBST(root.Right, val)
	} else if val == root.Val {
		return root
	} else {
		return searchBST(root.Left, val)
	}
}
