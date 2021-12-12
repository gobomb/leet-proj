package medium

import "sort"

/*
	274. H-Index
*/

func hIndex(citations []int) int {
	h := 0
	n := len(citations)
	sort.Ints(citations)
	for i := n - 1; i >= 0; i-- {
		if citations[i] >= n-i {
			h++
		} else {
			break
		}
	}

	return h
}
