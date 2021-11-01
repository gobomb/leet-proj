package leetcode

/*
	115. Distinct Subsequences
*/

/*
Runtime: 16 ms, faster than 23.40% of Go online submissions for Distinct Subsequences.
Memory Usage: 17 MB, less than 14.36% of Go online submissions for Distinct Subsequences.
*/
func numDistinct(s string, t string) int {
	mem := make([][]int, len(s)+1)
	for i := 0; i < len(s)+1; i++ {
		mem[i] = make([]int, len(t)+1)
	}
	for j := 0; j < len(t)+1; j++ {
		mem[len(s)][j] = 0
	}
	for i := 0; i < len(s)+1; i++ {
		mem[i][len(t)] = 1
	}

	for i := len(s) - 1; i >= 0; i-- {
		for j := len(t) - 1; j >= 0; j-- {
			if s[i] == t[j] {
				// 选择：是否匹配当前字符
				// 匹配取 mem[i+1][j+1]
				// 不匹配取 mem[i+1][j]
				mem[i][j] = mem[i+1][j+1] + mem[i+1][j]
			} else {
				mem[i][j] = mem[i+1][j]
			}
		}
	}

	return mem[0][0]
}

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Distinct Subsequences.
Memory Usage: 2.1 MB, less than 86.17% of Go online submissions for Distinct Subsequences.
*/
func numDistinct2(s string, t string) int {
	mem := make([]int, len(t)+1)

	for j := 0; j < len(t); j++ {
		mem[j] = 0
	}
	mem[len(t)] = 1
	tp, tp2 := 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		tp = 1
		for j := len(t) - 1; j >= 0; j-- {
			tp2 = tp
			tp = mem[j]
			if s[i] == t[j] {
				mem[j] = tp2 + mem[j]
			}
		}
	}
	return mem[0]
}
