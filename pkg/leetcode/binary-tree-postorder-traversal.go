package leetcode


var btpt_rs []int

func postorderTraversal(root *TreeNode) []int {
	btpt_rs = btpt_rs[:0]
	postorder(root)

	return btpt_rs
}

func postorder(r *TreeNode) {
	if r != nil {
		postorder(r.Left)
		postorder(r.Right)
		btpt_rs = append(btpt_rs, r.Val)
	}
}
