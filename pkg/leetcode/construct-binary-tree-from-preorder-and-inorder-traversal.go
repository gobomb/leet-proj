package leetcode

func buildTree(preorder []int, inorder []int) *TreeNode {
	pre := 0
	return buildTreeByPreInOrder(preorder, inorder, &pre, 0, len(inorder)-1)
}

// pre: 指向 preorder 的当前下标，用于获取根节点
// inStart,inEnd: inorder 的初始坐标和结束坐标
func buildTreeByPreInOrder(preorder []int, inorder []int, pre *int, inStart, inEnd int) *TreeNode {
	// 如果为空，返回 nil，pre 减 1,因为这一层没有根节点
	if inStart > inEnd {
		*pre--
		return nil
	}
	root := &TreeNode{
		Val: preorder[*pre],
	}

	// 找到 inorder 中根所在的下标
	k := inStart
	for preorder[*pre] != inorder[k] {
		k++
	}

	// preorder 的下一个坐标为根
	*pre++
	// 以 k 为中心，分别递归处理 inorder
	root.Left = buildTreeByPreInOrder(preorder, inorder, pre, inStart, k-1)
	*pre++
	root.Right = buildTreeByPreInOrder(preorder, inorder, pre, k+1, inEnd)
	return root
}
