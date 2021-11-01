package leetcode

func flatten(root *TreeNode) {
	n := preorderTraversalFlatten(root)
	if root != nil {
		root.Left = n.Left
		root.Right = n.Right
	}
}

func preorderTraversalFlatten(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	newNode := &TreeNode{
		Val: root.Val,
	}

	right := preorderTraversalFlatten(root.Left)
	if right != nil {
		newNode.Right = right
	}
	right = preorderTraversalFlatten(root.Right)
	next := newNode
	for next.Right != nil {
		next = next.Right
	}
	if right != nil {
		next.Right = right
	}

	return newNode
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
			next := n.Left
			for next.Right != nil {
				next = next.Right
			}
			next.Right = n.Right
			n.Right = n.Left
			n.Left = nil
		}

		n = n.Right
	}
}
