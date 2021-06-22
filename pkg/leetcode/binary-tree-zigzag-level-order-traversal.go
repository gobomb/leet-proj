package leetcode

func zigzagLevelOrder(root *TreeNode) [][]int {
	rs := [][]int{}
	getZigzagLevelOrder(root, 0, &rs)

	return rs
}

func getZigzagLevelOrder(root *TreeNode, depth int, rs *[][]int) {
	if root == nil {
		return
	}
	if len(*rs) == depth {
		*rs = append(*rs, []int{})
	}
	if depth%2 == 0 {
		(*rs)[depth] = append((*rs)[depth], root.Val)
	} else {
		temp := append([]int{}, root.Val)
		(*rs)[depth] = append(temp, (*rs)[depth]...)
	}
	getZigzagLevelOrder(root.Left, depth+1, rs)
	getZigzagLevelOrder(root.Right, depth+1, rs)
}
