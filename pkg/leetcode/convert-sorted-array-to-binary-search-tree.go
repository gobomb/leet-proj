package leetcode

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	val := nums[len(nums)/2]
	tree := &TreeNode{
		Val:   val,
		Left:  sortedArrayToBST(nums[:len(nums)/2]),
		Right: sortedArrayToBST(nums[len(nums)/2+1:]),
	}
	return tree
}
