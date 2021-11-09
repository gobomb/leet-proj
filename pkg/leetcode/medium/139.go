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

/*
	140. Word Break II
*/

// https://leetcode.com/problems/word-break-ii/discuss/440571/golang-solution
func wordBreakII(s string, wordDict []string) []string {
	if s == "" || len(wordDict) == 0 {
		return []string{}
	}

	return dfsWordBreak(s, wordDict, 0, make(map[int][]string))
}

func dfsWordBreak(s string, wordDict []string, index int, record map[int][]string) []string {
	if result, ok := record[index]; ok {
		return result
	}

	result := []string{}

	for i := index + 1; i <= len(s); i++ {
		if checkWordInDict(s[index:i], wordDict) {
			if i != len(s) {
				rs := dfsWordBreak(s, wordDict, i, record)
				for _, r := range rs {
					result = append(result, s[index:i]+" "+r)
				}
			} else {
				result = append(result, s[index:i])
				break
			}
		}
	}

	record[index] = result

	return result
}
