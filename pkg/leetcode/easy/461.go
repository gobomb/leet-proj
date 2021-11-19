package easy

/*
	461. Hamming Distance
*/

func hammingDistance(x int, y int) int {
	x = (x | y) & (^x | ^y)

	res := 0

	for x != 0 {
		y := x - 1
		x = x & y
		res++
	}

	return res
}
