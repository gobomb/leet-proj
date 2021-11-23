package medium

import (
	"sort"
)

/*
	368. Largest Divisible Subset
*/

// https://leetcode-cn.com/problems/largest-divisible-subset/solution/gong-shui-san-xie-noxiang-xin-ke-xue-xi-0a3jc/
/*
	Runtime: 36 ms, faster than 100.00% of Go online submissions for Largest Divisible Subset.
	Memory Usage: 3 MB, less than 100.00% of Go online submissions for Largest Divisible Subset.
*/
func largestDivisibleSubset(nums []int) []int {
	sort.Ints(nums)

	// dp[i] 表示 i 位置的数字对应 i 以前可互相整除的数字的最大长度
	dp := make([]int, len(nums))
	// 用于回溯出结果数组
	bt := make([]int, len(nums))

	for i := range nums {
		// 至少为 1 即 i 本身
		l := 1
		prev := i

		for j := 0; j < i; j++ {
			if nums[i]%nums[j] == 0 {
				// 选择出最大长度对应的下标·
				if dp[j]+1 > l {
					l = dp[j] + 1
					prev = j
				}
			}
		}
		dp[i] = l
		bt[i] = prev
	}

	ind := -1
	maxi := -1

	for i := range dp {
		if dp[i] > maxi {
			maxi = dp[i]
			ind = i
		}
	}

	res := make([]int, maxi)
	for i := len(res) - 1; i >= 0; i-- {
		res[i] = nums[ind]
		ind = bt[ind]
	}

	return res
}
