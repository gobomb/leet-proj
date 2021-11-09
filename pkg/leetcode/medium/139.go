package medium

/*
	139. Word Break
*/

func wordBreak(s string, wordDict []string) bool {
	if s == "" || len(wordDict) == 0 {
		return true
	}

	dp := make([]bool, len(s)+1)

	dp[0] = true

	for i := 1; i < len(dp); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && checkWordInDict(s[j:i], wordDict) {
				dp[i] = true
				break
			}
		}
	}

	return dp[len(s)]
}

func checkWordInDict(s string, wordDict []string) bool {
	for _, w := range wordDict {
		if s == w {
			return true
		}
	}

	return false
}
