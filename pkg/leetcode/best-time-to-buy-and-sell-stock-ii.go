package leetcode

/*
	122. Best Time to Buy and Sell Stock II
*/

func maxProfitII(prices []int) int {
	max := 0
	for i := 0; i < len(prices)-1; i++ {
		if prices[i+1] > prices[i] {
			max += prices[i+1] - prices[i]
		}
	}
	return max
}
