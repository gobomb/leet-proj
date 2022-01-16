package medium

import (
	"math"
)

/*
	396. Rotate Function
*/

// https://leetcode.com/problems/rotate-function/discuss/87853/Java-O(n)-solution-with-explanation
func maxRotateFunction(nums []int) int {
	rs := math.MinInt
	sum := 0
	f := 0

	for i, n := range nums {
		f += i * n
		sum += n
	}

	for i := len(nums) - 1; i >= 0; i-- {
		f = f + sum - len(nums)*nums[i]
		if f > rs {
			rs = f
		}
	}

	return rs
}
