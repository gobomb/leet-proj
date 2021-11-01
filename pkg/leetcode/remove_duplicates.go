package leetcode

import "fmt"

func removeDuplicates2(nums []int) int {
	var j int
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] != nums[i+1] {
			j++
			nums[j] = nums[i+1]
		}
	}
	return j + 1
}

func removeDuplicates(nums []int) int {
	var lastIndex = len(nums) - 1
	for i := 0; i < len(nums)-1; i++ {
		if i >= lastIndex {
			break
		}
		lastIndex = detect(nums, i, lastIndex)
	}
	return lastIndex + 1
}

func detect(nums []int, i, lastIndex int) int {
	if nums[i] == nums[i+1] {
		copy(nums[i:], nums[i+1:])
		lastIndex--
		if i == lastIndex {
			return lastIndex
		}
		lastIndex = detect(nums, i, lastIndex)
	}
	return lastIndex
}

func testremoveDuplicates() {
	tests := [][]int{
		{1},
		{1, 1, 1, 3, 4, 4},
		{1, 1, 2},
		{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
	}
	for i := range tests {
		r := removeDuplicates2(tests[i])
		fmt.Println(r)
	}
}
