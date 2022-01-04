package medium

import (
	"math"
)

/*
	375. Guess Number Higher or Lower II
*/

// https://www.cnblogs.com/grandyang/p/5677550.html
func getMoneyAmount(n int) int {
	dp := make([][]int, n+1)

	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for hi := 2; hi < len(dp); hi++ {
		for lo := hi - 1; lo > 0; lo-- {
			globalMin := math.MaxInt
			for x := lo + 1; x < hi; x++ {
				localMin := x + max(dp[lo][x-1], dp[x+1][hi])
				globalMin = min(globalMin, localMin)
			}

			if lo+1 == hi {
				dp[lo][hi] = lo
			} else {
				dp[lo][hi] = globalMin
			}
		}
	}

	return dp[1][n]
}
