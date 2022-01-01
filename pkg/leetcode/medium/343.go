package medium

import "math"

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

func integerBreak1(n int) int {
	if n == 2 {
		return 1
	}

	if n == 3 {
		return 2
	}

	a := n / 3
	b := (n % 3) / 2

	if n%3 == 1 {
		a -= 1
		b = 2
	}

	return int(math.Pow(3, float64(a)) * math.Pow(2, float64(b)))
}

func integerBreak2(n int) int {
	if n == 2 {
		return 1
	}

	if n == 3 {
		return 2
	}

	rs := 1

	for n > 4 {
		rs *= 3
		n -= 3
	}

	rs *= n

	return rs
}
