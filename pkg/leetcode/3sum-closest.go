package leetcode

import (
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	result := 0
	save := 1000
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				sub := abs(nums[i] + nums[j] + nums[k] - target)
				//println(i, j, k, nums[i], nums[j], nums[k], sub, save)
				if sub <= save {
					save = sub
					result = nums[i] + nums[j] + nums[k]
				}
			}
		}
	}
	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
