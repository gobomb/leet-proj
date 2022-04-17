package review

// 144. binary-tree-preorder-traversal

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

// 迭代法
func preorderTraversal3(root *TreeNode) []int {
	var rs []int

	r := root
	stack := []*TreeNode{}

	// 入栈
	stack = append(stack, r)

	for len(stack) != 0 {
		// 出栈直到栈空
		r = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		rs = append(rs, r.Val)

		// 先压栈右节点
		if r.Right != nil {
			stack = append(stack, r.Right)
		}
		if r.Left != nil {
			stack = append(stack, r.Left)
		}
	}

	return rs
}
