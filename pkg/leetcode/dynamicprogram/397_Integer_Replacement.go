package dynamicprogram

/*
	397. Integer Replacement
*/

/*
	Runtime: 0 ms, faster than 100.00% of Go online submissions for Integer Replacement.
	Memory Usage: 1.9 MB, less than 100.00% of Go online submissions for Integer Replacement.

	使用递归而不是通过数组存所有结果
·*/
func integerReplacement(n int) int {
	// dp := make([]int, n+1)

	// for i := 2; i <= n; i++ {
	// 	if i%2 == 0 {
	// 		dp[i] = dp[i/2] + 1
	// 	} else {
	// 		dp[i] = min(dp[(i+1)/2], dp[(i-1)/2]) + 2
	// 	}
	// }

	// return dp[n]
	if n == 1 {
		return 0
	}
	if n%2 == 0 {
		return integerReplacement(n/2) + 1
	} else {
		return min(integerReplacement((n+1)/2), integerReplacement((n-1)/2)) + 2
	}
}
