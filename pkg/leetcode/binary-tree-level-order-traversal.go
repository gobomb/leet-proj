package leetcode

func levelOrder(root *TreeNode) [][]int {
	rs := [][]int{}
	getLevelOrder(root, 0, &rs)

	return rs
}

func getLevelOrder(root *TreeNode, depth int, rs *[][]int) {
	if root == nil {
		return
	}
	if len(*rs) == depth {
		*rs = append(*rs, []int{})
	}
	(*rs)[depth] = append((*rs)[depth], root.Val)

	getLevelOrder(root.Left, depth+1, rs)
	getLevelOrder(root.Right, depth+1, rs)
}
