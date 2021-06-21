package leetcode

func levelOrder(root *TreeNode) [][]int {
	rs := [][]int{}
	rs = append(rs, []int{})
	getLevelOrder(root, 0, &rs)
	for i := len(rs) - 1; i >= 0; i-- {
		if len(rs[i]) == 0 {
			rs = rs[:len(rs)-1]
		}
	}
	return rs
}

func getLevelOrder(root *TreeNode, depth int, rs *[][]int) {
	if root == nil {
		return
	}
	(*rs)[depth] = append((*rs)[depth], root.Val)
	*rs = append(*rs, []int{})

	getLevelOrder(root.Left, depth+1, rs)
	getLevelOrder(root.Right, depth+1, rs)

}
