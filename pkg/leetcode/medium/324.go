package medium

import (
	"justest/pkg/leetcode/review"
)

/*
	324. Wiggle Sort II
*/

// https://leetcode.com/problems/wiggle-sort-ii/discuss/77682/Step-by-step-explanation-of-index-mapping-in-Java
func wiggleSort(nums []int) {
	n := len(nums)

	if n < 2 {
		return
	}

	m := review.FindKthLargest1(nums, (n+1)/2)

	left := 0
	right := n - 1
	i := 0
	for i <= right {
		if nums[indexMap(i, n)] > m {
			nums[indexMap(left, n)], nums[indexMap(i, n)] = nums[indexMap(i, n)], nums[indexMap(left, n)]
			i++
			left++
		} else if nums[indexMap(i, n)] == m {
			i++
		} else {
			nums[indexMap(right, n)], nums[indexMap(i, n)] = nums[indexMap(i, n)], nums[indexMap(right, n)]
			right--
		}
	}
}

func indexMap(index, n int) int {
	return (1 + 2*index) % (n | 1)
}
