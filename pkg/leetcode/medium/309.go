package medium

import (
	"math"
)

/*
	309. Best Time to Buy and Sell Stock with Cooldown
*/
/*
	https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/discuss/75928/Share-my-DP-solution-(By-State-Machine-Thinking)

	Runtime: 0 ms, faster than 100.00% of Go online submissions for Best Time to Buy and Sell Stock with Cooldown.
	Memory Usage: 2.6 MB, less than 47.40% of Go online submissions for Best Time to Buy and Sell Stock with Cooldown.
*/
func maxProfit(prices []int) int {
	l := len(prices)
	s0 := make([]int, l)
	s1 := make([]int, l)
	s2 := make([]int, l)

	s0[0] = 0
	s1[0] = -prices[0]
	s2[0] = math.MinInt

	for i := 1; i < l; i++ {
		s0[i] = max(s0[i-1], s2[i-1])
		s1[i] = max(s1[i-1], s0[i-1]-prices[i])
		s2[i] = s1[i-1] + prices[i]
	}

	return max(s0[l-1], s2[l-1])
}

/*
	Runtime: 0 ms, faster than 100.00% of Go online submissions for Best Time to Buy and Sell Stock with Cooldown.
	Memory Usage: 2.4 MB, less than 64.29% of Go online submissions for Best Time to Buy and Sell Stock with Cooldown.
*/
func maxProfit1(prices []int) int {
	l := len(prices)

	a := 0
	b := -prices[0]
	c := math.MinInt

	for i := 1; i < l; i++ {
		a, b, c = max(a, c), max(b, a-prices[i]), b+prices[i]
	}

	return max(a, c)
}
