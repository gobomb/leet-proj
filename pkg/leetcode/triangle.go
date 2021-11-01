package leetcode

/*
	120. Triangle
*/

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}
	rs := 10001
	dp := make([]int, len(triangle))

	for i := 0; i < len(triangle[0]); i++ {
		dp[i] = triangle[0][i]
	}

	for i := 1; i < len(triangle); i++ {
		for j := len(triangle[i]) - 1; j >= 0; j-- {
			if j == 0 {
				dp[j] += triangle[i][j]
			} else if j == len(triangle[i])-1 {
				dp[j] += dp[j-1] + triangle[i][j]
			} else {
				dp[j] = min(dp[j-1]+triangle[i][j], dp[j]+triangle[i][j])
			}
		}
	}

	// log.Printf("%v", dp)
	for i := 0; i < len(dp); i++ {
		if dp[i] < rs {
			rs = dp[i]
		}
	}
	return rs
}
