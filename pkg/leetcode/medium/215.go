package medium

import (
	"sort"
)

/*
	215. Kth Largest Element in an Array
*/

func findKthLargest(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}

func findKthLargest1(nums []int, k int) int {
	k1 := len(nums) - k

	return qselect(nums, k1, 0, len(nums)-1)
}

func qselect(arr []int, k, left, right int) int {
	pivot := left
	index := pivot + 1

	for i := index; i <= right; i++ {
		if arr[i] < arr[pivot] {
			swap(arr, i, index)
			index += 1
		}
	}
	swap(arr, pivot, index-1)

	i := index - 1

	if k < i {
		return qselect(arr, k, left, i-1)
	} else if k > i {
		return qselect(arr, k, i+1, right)
	}

	return arr[i]
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
