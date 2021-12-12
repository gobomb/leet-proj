package medium

import "math"

/*
	279. Perfect Squares
*/

// https://books.halfrost.com/leetcode/ChapterFour/0200~0299/0279.Perfect-Squares/
func numSquares(n int) int {
	if isSquare(n) {
		return 1
	}

	if squares4(n) {
		return 4
	}

	for i := 0; i*i < n; i++ {
		if isSquare(n - i*i) {
			return 2
		}
	}

	return 3
}

func isSquare(n int) bool {
	s := int(math.Sqrt(float64(n)))

	return n == s*s
}

func squares4(x int) bool {
	for x%4 == 0 {
		x /= 4
	}

	return x%8 == 7
}
