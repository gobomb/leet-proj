package leetcode

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	mid := len(nums) / 2
	tree := &TreeNode{
		Val:   nums[mid],
		Left:  sortedArrayToBST(nums[:mid]),
		Right: sortedArrayToBST(nums[mid+1:]),
	}
	return tree
}

func sortedListToBST(head *ListNode) *TreeNode {
	nums := []int{}
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}
	return sortedArrayToBST(nums)
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	l := maxDepth(root.Left)
	r := maxDepth(root.Right)
	return Abs(l-r) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}
