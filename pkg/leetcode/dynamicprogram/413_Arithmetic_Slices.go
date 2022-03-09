package dynamicprogram

/*
	413. Arithmetic Slices
*/

func numberOfArithmeticSlices(nums []int) int {
	if len(nums) < 3 {
		return 0
	}

	dp := make([]int, len(nums))
	res := 0

	for i := 2; i < len(dp); i++ {
		if nums[i]-nums[i-1] == nums[i-1]-nums[i-2] {
			dp[i] = dp[i-1] + 1
			res += dp[i]
		}
	}

	return res
}
