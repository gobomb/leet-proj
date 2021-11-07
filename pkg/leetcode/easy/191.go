package easy

import (
	"math/bits"
)

/*
	191. Number of 1 Bits
*/

// 0 ms	2.1 MB
func hammingWeight(num uint32) int {
	// fmt.Printf("%b\n", num)
	return bits.OnesCount32(num)
}

// 4 ms	2.1 MB
// https://leetcode.com/problems/number-of-1-bits/discuss/285270/Golang-faster-than-100
func hammingWeight2(num uint32) int {
	val := 0
	for num > 0 {
		val += int(num & 1)
		num >>= 1
	}

	return val
}

// 0 ms	2 MB
// https://books.halfrost.com/leetcode/ChapterFour/0100~0199/0191.Number-of-1-Bits/
func hammingWeight3(num uint32) int {
	count := 0

	for num != 0 {
		// fmt.Printf("%b %b %b\n", num, num-1, num&(num-1))
		num &= (num - 1)
		count++
	}

	return count
}
