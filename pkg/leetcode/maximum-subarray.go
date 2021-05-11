package leetcode

import (
	"log"
	"math"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

// 暴力
//88 ms	3.3 MB
func maxSubArray(nums []int) int {
	// log.Printf("%v\n", nums)
	maxRecord := -math.MaxInt64
	for i := 0; i < len(nums); i++ {
		temp := 0
		for j := i; j < len(nums); j++ {
			temp += nums[j]
			if temp > maxRecord {
				maxRecord = temp
			}
		}

	}
	return maxRecord
}

// dp[i] = nums[i] + dp[i-1] (dp[i-1] > 0)
// dp[i] = nums[i] (dp[i-1] ≤ 0)

// 4 ms	3.5 MB
// dp 循环
func maxSubArray2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		if dp[i-1] > 0 {
			dp[i] = nums[i] + dp[i-1]
		} else {
			dp[i] = nums[i]
		}
		res = max(dp[i], res)
	}
	// log.Printf("%v\n", dp)

	return res
}

// 0 ms	6 MB
// dp 递归
func maxSubArray3(nums []int) int {
	dp := make([]int, len(nums))
	return dpsubarray3(nums, dp, len(nums)-1)
}
func dpsubarray3(nums []int, dp []int, i int) int {
	if i == 0 {
		dp[i] = nums[i]
		return dp[i]
	}
	res := dpsubarray3(nums, dp, i-1)
	if dp[i-1] > 0 {
		dp[i] = nums[i] + dp[i-1]
	} else {
		dp[i] = nums[i]
	}
	res = max(dp[i], res)

	return res
}

// 	4 ms 3.3 MB
func maxSubArray1(nums []int) int {
	maxSum := nums[0]
	record := 0
	for i := 0; i < len(nums); i++ {
		record += nums[i]
		if record > maxSum {
			maxSum = record
		}
		if record < 0 {
			record = 0
		}
	}
	return maxSum
}
