package leetcode

// 654. Maximum Binary Tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	var max, ind int = -1, -1
	for i := range nums {
		if nums[i] > max {
			max = nums[i]
			ind = i
		}
	}
	var left, right *TreeNode
	if ind == len(nums)-1 {
		left = constructMaximumBinaryTree(nums[:ind])
	} else if ind == 0 {
		right = constructMaximumBinaryTree(nums[ind+1:])
	} else {
		left = constructMaximumBinaryTree(nums[:ind])
		right = constructMaximumBinaryTree(nums[ind+1:])
	}
	return &TreeNode{
		Val:   max,
		Left:  left,
		Right: right,
	}
}
