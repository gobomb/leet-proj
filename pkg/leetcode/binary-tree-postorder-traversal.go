package leetcode

var btptRs []int

func postorderTraversal(root *TreeNode) []int {
	btptRs = btptRs[:0]

	postorder(root)

	return btptRs
}

func postorder(r *TreeNode) {
	if r != nil {
		postorder(r.Left)
		postorder(r.Right)
		btptRs = append(btptRs, r.Val)
	}
}
