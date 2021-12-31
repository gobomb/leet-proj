package medium

/*
	343. Integer Break
*/

func integerBreak(n int) int {
	dp := make([]int, n+1)

	dp[0] = 1
	dp[1] = 1

	for i := 1; i <= n; i++ {
		for j := 1; j < i; j++ {
			dp[i] = max(dp[i], j*max(dp[i-j], i-j))
		}
	}

	return dp[n]
}
