package medium

import (
	"log"
	"sort"
)

/*
	300. Longest Increasing Subsequence
*/

func lengthOfLIS(nums []int) int {
	res := 0

	dp := make([]int, len(nums))

	for i := range dp {
		dp[i] = 1
	}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}

		res = max(res, dp[i])
	}

	return res
}

func lengthOfLIS1(nums []int) int {
	dp := []int{}

	for _, num := range nums {
		i := sort.SearchInts(dp, num)

		if i == len(dp) {
			dp = append(dp, num)
		} else {
			dp[i] = num
		}
	}

	return len(dp)
}
