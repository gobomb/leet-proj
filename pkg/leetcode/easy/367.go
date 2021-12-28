package easy

/*
	367. Valid Perfect Square
*/

func isPerfectSquare(num int) bool {
	lo, hi, mi := 0, num, 0

	for lo < hi {
		mi = lo + (hi-lo)>>1
		if mi*mi == num {
			hi = mi
		} else if mi*mi < num {
			lo = mi + 1
		} else if mi*mi > num {
			hi = mi
		}
	}

	return lo*lo == num
}
