package medium

import "math"

/*
	313. Super Ugly Number
*/

/*
	Runtime: 429 ms, faster than 25.81% of Go online submissions for Super Ugly Number.
	Memory Usage: 11.9 MB, less than 19.35% of Go online submissions for Super Ugly Number.
*/
func nthSuperUglyNumber(n int, primes []int) int {
	dp := make([]int, n)
	dp[0] = 1

	p := make([]int, len(primes))

	for i := 1; i < n; i++ {
		maxq := math.MaxInt64

		for j := 0; j < len(primes); j++ {
			v := dp[p[j]] * primes[j]

			if v < maxq {
				maxq = v
			}
		}

		for j := 0; j < len(primes); j++ {
			v := dp[p[j]] * primes[j]

			if v == maxq {
				p[j]++
			}
		}

		dp[i] = maxq
	}

	// log.Println(dp)

	return dp[n-1]
}
