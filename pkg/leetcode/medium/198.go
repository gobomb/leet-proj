package medium

import "log"

/*
	198. House Robber
*/

// 0 ms	2.1 MB
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	dp := make([]int, len(nums))

	dp[0] = nums[0]

	if nums[0] < nums[1] {
		dp[1] = nums[1]
	} else {
		dp[1] = dp[0]
	}

	for i := 2; i < len(nums); i++ {
		if nums[i]+dp[i-2] > dp[i-1] {
			dp[i] = nums[i] + dp[i-2]
		} else {
			dp[i] = dp[i-1]
		}
	}

	log.Println(dp)

	return dp[len(nums)-1]
}

// 0 ms	2 MB
func rob1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	a, b, c := 0, 0, 0
	a = nums[0]

	if nums[0] < nums[1] {
		b = nums[1]
	} else {
		b = a
	}

	for i := 2; i < len(nums); i++ {
		if nums[i]+a > b {
			c = nums[i] + a
		} else {
			c = b
		}
		a, b = b, c
	}

	// log.Println(b, c)

	return b
}
