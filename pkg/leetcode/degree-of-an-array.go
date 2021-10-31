package leetcode

/*
	697. Degree of an Array
*/

import (
	"math"
)

func findShortestSubArray(nums []int) int {
	set := make(map[int]int)
	distance := make(map[int]int)
	distanceLast := make(map[int]int)

	for i := range nums {
		n := nums[i]
		set[n]++
		if set[n] == 1 {
			distance[n] = i
		} else {
			distanceLast[n] = i - distance[n]
		}
	}

	ks := []int{}

	counts := make([]int, 0, len(set))
	for _, v := range set {
		counts = append(counts, v)
	}

	maxCount := counts[0]

	for i := range counts {
		if counts[i] > maxCount {
			maxCount = counts[i]
		}
	}

	for k, v := range set {
		if v == maxCount {
			ks = append(ks, k)
		}
	}

	res := math.MaxInt32

	for _, k := range ks {
		if distanceLast[k] < res {
			res = distanceLast[k]
		}
	}
	return res + 1
}
