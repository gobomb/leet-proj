package medium

/*
	201. Bitwise AND of Numbers Range
*/

func rangeBitwiseAnd(left int, right int) int {
	for left < right {
		right &= right - 1
	}

	return right
}

func rangeBitwiseAnd1(left int, right int) int {
	// log.Printf("%b %b", left, right)
	i := 0

	for left != right {
		left >>= 1
		right >>= 1
		i += 1
	}
	// log.Printf("%b %b", left, right)

	return right << i
}
