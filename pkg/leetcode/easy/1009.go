package easy

/*
	1009. Complement of Base 10 Integer
*/

func bitwiseComplement(n int) int {
	if n == 0 {
		return 1
	}

	rs := 0
	m := 1

	for n > 0 {
		if n&1 == 0 {
			rs += m
		}

		m <<= 1
		n >>= 1
	}

	return rs
}

func bitwiseComplement1(n int) int {
	m := 1

	for n > m {
		m = m<<1 + 1
	}

	return m ^ n
}
