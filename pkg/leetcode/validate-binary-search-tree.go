package leetcode

import (
	"math"
)

var prev = math.MinInt64

func isValidBST(root *TreeNode) bool {
	// 每次执行用例全局变量需要重新赋值
	prev = math.MinInt64
	return isValidBST2(root)
}
func isValidBST2(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if !isValidBST2(root.Left) {
		return false
	}

	if root.Val <= prev {
		return false
	}
	prev = root.Val

	return isValidBST2(root.Right)
}

func isValidBST3(root *TreeNode) bool {
	arr := []int{}
	getOrder(root, &arr)
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] >= arr[i+1] {
			return false
		}
	}
	return true
}

func getOrder(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}
	getOrder(root.Left, arr)
	*arr = append(*arr, root.Val)
	getOrder(root.Right, arr)
}
