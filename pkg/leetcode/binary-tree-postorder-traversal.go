package leetcode

var rs []int

func postorderTraversal(root *TreeNode) []int {
	rs = rs[:0]
	postorder(root)

	return rs
}

func postorder(r *TreeNode) {
	if r != nil {
		postorder(r.Left)
		postorder(r.Right)
		rs = append(rs, r.Val)
	}
}
