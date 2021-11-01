package leetcode

import "justest/pkg/ds"

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if targetSum == root.Val && root.Left == nil && root.Right == nil {
		return true
	}
	targetSum = targetSum - root.Val

	if hasPathSum(root.Left, targetSum) {
		return true
	}
	return hasPathSum(root.Right, targetSum)
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	rs := [][]int{}
	nums := []int{}
	pathSum2(root, targetSum, &rs, nums)
	return rs
}

func pathSum2(root *TreeNode, targetSum int, rs *[][]int, nums []int) {
	if root == nil {
		return
	}
	if targetSum == root.Val && root.Left == nil && root.Right == nil {
		nums = append(nums, root.Val)
		*rs = append(*rs, ds.DeepCopyIntSlice(nums))
		return
	}
	targetSum = targetSum - root.Val
	nums = append(nums, root.Val)
	pathSum2(root.Left, targetSum, rs, nums)

	pathSum2(root.Right, targetSum, rs, nums)
	// nums = nums[:len(nums)-1]
}
