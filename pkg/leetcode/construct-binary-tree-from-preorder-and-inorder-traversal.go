package leetcode

func buildTree(preorder []int, inorder []int) *TreeNode {
	return buildTreeByPreInOrder(preorder, inorder, 0, 0, len(preorder))
}

func buildTreeByPreInOrder(preorder []int, inorder []int, pre, in, n int) *TreeNode {
	if n == 0 {
		return nil
	}
	root := &TreeNode{
		Val: preorder[pre],
	}

	k := 0
	for preorder[pre] != inorder[k] {
		k++
	}

	root.Left = buildTreeByPreInOrder(preorder, inorder, pre+1, in, k)
	root.Right = buildTreeByPreInOrder(preorder, inorder, pre+k+1, in+k+1, n-k-1)
	return root
}
