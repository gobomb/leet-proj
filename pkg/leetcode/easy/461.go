package easy

/*
	461. Hamming Distance
*/

func hammingDistance(x int, y int) int {
	// x = (x | y) & (^x | ^y)
	x ^= y

	res := 0

	for x != 0 {
		x &= (x - 1)
		res++
	}

	return res
}
