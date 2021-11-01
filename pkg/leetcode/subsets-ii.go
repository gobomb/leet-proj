package leetcode

import (
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	rs := [][]int{}
	r := []int{}
	rs = append(rs, r)

	sort.Ints(nums)
	subsetsBT2(nums, &rs, r, 0)
	return rs
}

func subsetsBT2(nums []int, rs *[][]int, r []int, a int) {
	for i := a; i < len(nums); i++ {

		if i > a && nums[i] == nums[i-1] {
			continue
		}
		r = append(r, nums[i])

		b := make([]int, len(r))
		copy(b, r)

		*rs = append(*rs, b)

		subsetsBT2(nums, rs, r, i+1)
		r = r[:len(r)-1]
	}
}
