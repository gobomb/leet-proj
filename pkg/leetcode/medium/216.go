package medium

import (
	"justest/pkg/utils"
)

/*
	216. Combination Sum III

	Runtime: 0 ms, faster than 100.00% of Go online submissions for Combination Sum III.
	Memory Usage: 2 MB, less than 55.56% of Go online submissions for Combination Sum III.
*/

func combinationSum3(k int, n int) [][]int {
	choose := &[][]int{}

	combination22(k, n, []int{}, choose, 1)

	return *choose
}

func count(sums []int) (cs int) {
	for i := 0; i < len(sums); i++ {
		cs += sums[i]
	}

	return
}

func combination22(k int, target int, sums []int, choose *[][]int, ind int) bool {
	if len(sums) > k {
		return false
	}

	sum := count(sums)
	if sum > target {
		return false
	}

	if sum == target {
		return true
	}

	for i := ind; i <= 9; i++ {
		sums = append(sums, i)
		rs := combination22(k, target, sums, choose, i+1)

		if rs && len(sums) == k {
			*choose = append(*choose, utils.DeepCopyIntSlice(sums))
		}
		sums = sums[:len(sums)-1]
	}

	return false
}
