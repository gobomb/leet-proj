package leetcode

/*
	693. Binary Number with Alternating Bits
*/

import "math"

func hasAlternatingBits(n int) bool {
	m := n
	for i := 0; ; i = i + 2 {
		if m == 0 {
			return true
		}
		if m < 0 {
			break
		}
		m = m - int(math.Pow(2, float64(i)))
	}
	m = n
	for i := 1; ; i = i + 2 {
		if m == 0 {
			return true
		}
		if m < 0 {
			break
		}
		m = m - int(math.Pow(2, float64(i)))
	}
	return false
}
