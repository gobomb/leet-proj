package leetcode

import (
	"math"
)

func grayCode(n int) []int {
	l := int(math.Pow(2, float64(n)))
	bin := make([]int, n)
	rs := make([]int, 0)
	for i := 0; i < l; i++ {

		if i%2 == 1 {
			bin[n-1] = flip(bin[n-1])
		} else {
			for k := n - 1; k >= 1; k-- {
				if bin[k] == 1 {
					bin[k-1] = flip(bin[k-1])
					break
				}
			}
		}
		rs = append(rs, bintoin(bin))
	}
	return rs
}

func flip(i int) int {
	if i == 1 {
		i = 0
	} else {
		i = 1
	}
	return i
}

func bintoin(bin []int) int {
	var sum int
	for i := 0; i < len(bin); i++ {
		if bin[i] == 1 {
			sum += int(math.Pow(2, float64(len(bin)-i-1)))
		}
	}
	return sum
}
