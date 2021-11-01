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

func levelOrderBottom(root *TreeNode) [][]int {
	rs := [][]int{}
	getLevelOrderBottom(root, 0, &rs)

	return rs
}

func getLevelOrderBottom(root *TreeNode, depth int, rs *[][]int) {
	if root == nil {
		return
	}
	if len(*rs) == depth {
		*rs = append([][]int{{}}, *rs...)
	}
	ind := len(*rs) - 1 - depth
	// log.Printf("%+v %+v\n", len(*rs), depth)
	(*rs)[ind] = append((*rs)[ind], root.Val)

	getLevelOrderBottom(root.Left, depth+1, rs)
	getLevelOrderBottom(root.Right, depth+1, rs)
}
