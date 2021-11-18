package medium

/*
	172. Factorial Trailing Zeroes
*/

func trailingZeroes(n int) int {
	count := 0

	for n > 0 {
		count += n / 5
		n = n / 5
	}

	return count
}
