package medium

import (
	"sort"
)

func minSetSize(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	count := map[int]int{}
	for _, num := range arr {
		count[num]++
	}

	counts := []int{}
	for _, v := range count {
		counts = append(counts, v)
	}
	sort.Ints(counts)

	c := 0
	sum := 0
	for i := len(counts) - 1; i >= 0; i-- {
		sum += counts[i]
		c++

		if sum >= len(arr)/2 {
			break
		}
	}

	return c
}
