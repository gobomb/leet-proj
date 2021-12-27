package easy

/*
	338. Counting Bits
*/

func countBits(n int) []int {
	rs := make([]int, n+1)

	for i := 1; i <= n; i++ {
		rs[i] += rs[i&(i-1)] + 1
	}

	return rs
}
