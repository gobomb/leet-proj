package medium

/*
	322. Coin Change
*/

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0

	for i := 1; i <= amount; i++ {
		a := amount + 1

		for _, c := range coins {
			if i >= c && dp[i-c] != -1 && dp[i-c]+1 < a {
				a = dp[i-c] + 1
			}
		}

		if a == amount+1 {
			a = -1
		}

		dp[i] = a
	}

	return dp[amount]
}
