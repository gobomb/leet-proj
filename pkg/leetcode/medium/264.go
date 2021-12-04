package medium

/*
	264. Ugly Number II
*/

func nthUglyNumber(n int) int {
	dp := make([]int, n)
	dp[0] = 1
	p1, p2, p3 := 0, 0, 0

	for i := 1; i < n; i++ {
		v1 := dp[p1] * 2
		v2 := dp[p2] * 3
		v3 := dp[p3] * 5

		m := min(v1, min(v2, v3))
		dp[i] = m

		if m == v1 {
			p1++
		}

		if m == v2 {
			p2++
		}

		if m == v3 {
			p3++
		}
	}

	return dp[n-1]
}
