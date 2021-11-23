package easy

/*
	1. Two Sum
*/

func twoSum1(nums []int, target int) []int {
	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return []int{}
}

// https://leetcode.com/problems/two-sum/discuss/261590/100-Golang
func twoSum2(nums []int, target int) []int {
	lookup := make(map[int]int)

	for i, v := range nums {
		j, ok := lookup[-v]
		lookup[v-target] = i

		if ok {
			return []int{j, i}
		}
	}

	return []int{}
}
