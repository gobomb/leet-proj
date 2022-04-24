package review

// 1143. longest-common-subsequence

func longestCommonSubsequence(t1, t2 string) int {
	dp := make([][]int, len(t1)+1)

	for k := 0; k < len(t1)+1; k++ {
		dp[k] = make([]int, len(t2)+1)
	}

	for i := 0; i < len(t1); i++ {
		for j := 0; j < len(t2); j++ {
			if t1[i] == t2[j] {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] =
					max(dp[i+1][j], dp[i][j+1])
			}
		}
	}

	return dp[len(t1)][len(t2)]
}
