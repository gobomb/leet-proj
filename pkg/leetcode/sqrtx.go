package leetcode

func mySqrt(x int) int {
	var hi, lo = x, 0
	var mi int
	for hi >= lo {
		mi = (hi + lo) / 2
		if mi*mi == x {
			return mi
		} else if mi*mi > x {
			if (mi-1)*(mi-1) <= x {
				return mi - 1
			}
			hi = mi - 1
		} else if mi*mi < x {
			lo = mi + 1
		}
	}
	return -1
}
