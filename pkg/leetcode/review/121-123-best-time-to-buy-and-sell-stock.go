package review

/*
	121. Best Time to Buy and Sell Stock
*/

func maxProfit(prices []int) int {
	max := 0
	min := prices[0]

	for i := 1; i < len(prices); i++ {
		if prices[i]-min > max {
			max = prices[i] - min
		}
		if prices[i] < min {
			min = prices[i]
		}
	}
	return max
}

/*
	123. Best Time to Buy and Sell Stock III
*/

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iii/discuss/149383/Easy-DP-solution-using-state-machine-O(n)-time-complexity-O(1)-space-complexity
func maxProfitIII(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	minINT := -10001

	s1, s2, s3, s4 := -prices[0], minINT, minINT, minINT
	for i := 1; i < len(prices); i++ {
		s1 = max(s1, -prices[i])
		s2 = max(s2, s1+prices[i])
		s3 = max(s3, s2-prices[i])
		s4 = max(s4, s3+prices[i])
	}
	return max(0, s4)
}

// https://zhuanlan.zhihu.com/p/77666061
// 用 dp[i][k] 表示前i天最多交易k次的最高收益
// dp[i][k] = Max(dp[i-1][k],(prices[i] - prices[0] + dp[0][k-1]),(prices[i] - prices[1] + dp[1][k-1])...(prices[i] - prices[i] + dp[i][k-1]))

//179 ms	15.8 MB
func maxProfitIIIdp(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	t := 2
	dp := make([][]int, len(prices))

	for i := range dp {
		dp[i] = make([]int, t+1)
	}

	for k := 1; k <= t; k++ {
		m := prices[0]
		for i := 1; i < len(prices); i++ {
			m = min(m, prices[i]-dp[i][k-1])

			dp[i][k] = max(dp[i-1][k], prices[i]-m)
		}
	}

	return dp[len(prices)-1][t]
}

//180 ms	9.5 MB
func maxProfitIIIdp2(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	t := 2
	m1 := make([]int, t+1)
	m2 := make([]int, t+1)

	for i := 0; i <= t; i++ {
		m1[i] = prices[0]
	}

	for i := 1; i < len(prices); i++ {
		for k := 1; k <= t; k++ {
			m1[k] = min(m1[k], prices[i]-m2[k-1])
			m2[k] = max(m2[k], prices[i]-m1[k])
		}
	}

	return m2[t]
}
