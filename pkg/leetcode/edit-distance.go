package leetcode

var dp [][]int

func minDistance(word1 string, word2 string) int {
	dp = make([][]int, len(word1)+1)
	for i := 0; i < len(word1)+1; i++ {
		dp[i] = make([]int, len(word2)+1)
	}
	return dpMinDistaance(len(word1), len(word2), word1, word2)
}

/*
https://blog.csdn.net/flying_all/article/details/99535006

1. 画出矩阵，总结状态转移方式
2. 通过状态转移方式得出状态转移方程
*/

func dpMinDistaance(a, b int, word1, word2 string) int {
	if dp[a][b] != 0 {
		return dp[a][b]
	}

	if min(a, b) == 0 {
		dp[a][b] = max(a, b)
	} else {
		d := 0
		if word1[a-1] == word2[b-1] {
			d = 0
		} else {
			d = 1
		}
		x := dpMinDistaance(a-1, b, word1, word2) + 1
		y := dpMinDistaance(a, b-1, word1, word2) + 1
		z := dpMinDistaance(a-1, b-1, word1, word2) + d
		dp[a][b] = min(x, min(y, z))

	}
	return dp[a][b]
}
