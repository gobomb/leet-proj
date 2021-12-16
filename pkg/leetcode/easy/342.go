package easy

/*
	342. Power of Four
*/

func isPowerOfFour(n int) bool {
	if n <= 0 {
		return false
	}

	for n != 1 {
		rm := n % 4
		if rm != 0 {
			return false
		}
		n /= 4
	}

	return true
}

func isPowerOfFour2(n int) bool {
	if n&(n-1) != 0 {
		return false
	}

	for n != 1 && n != 0 {
		n >>= 2
	}

	return n == 1
}

// 数论
func isPowerOfFour1(num int) bool {
	return num > 0 && (num&(num-1)) == 0 && (num-1)%3 == 0
}
