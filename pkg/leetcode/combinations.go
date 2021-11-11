package leetcode

import (
	"justest/pkg/utils"
)

func combine(n int, k int) [][]int {
	rs := [][]int{}
	aa := []int{}
	combineBacktrack(1, n, k, aa, &rs)

	return rs
}

func combineBacktrack(a, n, k int, aa []int, rs *[][]int) {
	if len(aa) == k {
		*rs = append(*rs, utils.DeepCopyIntSlice(aa))
		return
	}
	if len(aa) > k {
		return
	}
	for i := a; i <= n; i++ {
		aa = append(aa, i)
		combineBacktrack(i+1, n, k, aa, rs)
		aa = aa[:len(aa)-1]
	}
}
