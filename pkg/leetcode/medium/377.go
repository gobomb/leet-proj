package medium

/*
	377. Combination Sum IV
*/

func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)

	dp[0] = 1

	for i := 1; i <= target; i++ {
		for j := range nums {
			if nums[j] <= i {
				dp[i] += dp[i-nums[j]]
			}
		}
	}

	return dp[target]
}
