package medium

import (
	"sort"
)

/*
	275. H-Index II
*/

func hIndexII(citations []int) int {
	n := len(citations)

	hi := sort.Search(n, func(i int) bool {
		return citations[i] >= n-i
	})

	return n - hi
}
