package leetcode

import "fmt"

func smallerNumbersThanCurrent(nums []int) []int {
	var rs []int = make([]int, len(nums))
	for i := range nums {
		rs[i] = 0
		for j := range nums {
			if nums[j] < nums[i] {
				rs[i]++
			}
		}
	}
	return rs
}

func smallerNumbersThanCurrent2(nums []int) []int {
	fmt.Println(nums)

	bk := make([]int, 101)
	for i := range nums {
		bk[nums[i]] += 1

	}
	for i := 1; i <= 100; i++ {
		bk[i] += bk[i-1]
	}
	rs := make([]int, len(nums))
	for i := range nums {
		if nums[i] != 0 {
			rs[i] = bk[nums[i]-1]
		}
	}
	return rs
}
