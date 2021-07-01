package leetcode

func flatten(root *TreeNode) {
	new := preorderTraversalFlatten(root)
	if root != nil {
		root.Left = new.Left
		root.Right = new.Right
	}
}

func preorderTraversalFlatten(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	new := &TreeNode{
		Val: root.Val,
	}

	right := preorderTraversalFlatten(root.Left)
	if right != nil {
		new.Right = right
	}
	right = preorderTraversalFlatten(root.Right)
	next := new
	for next.Right != nil {
		next = next.Right
	}
	if right != nil {
		next.Right = right
	}

	return new
}

func preorderTraversalFlatten2(root *TreeNode) {
	if root == nil {
		return
	}

	tmp := root.Right
	preorderTraversalFlatten2(root.Left)
	preorderTraversalFlatten2(root.Right)
	if root.Left == nil {
		return
	}
	root.Right = root.Left

	next := root.Left
	root.Left = nil
	for next.Right != nil {
		next = next.Right
	}
	next.Right = tmp
}

func flatten2(root *TreeNode) {
	n := root
	for n != nil {
		if n.Left != nil {
			tmp := n.Right
			n.Right = n.Left
			next := n.Left
			n.Left = nil
			for next.Right != nil {
				next = next.Right
			}
			next.Right = tmp
		}

		n = n.Right
	}
}
