package leetcode

import (
	"justest/pkg/utils"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	choose := &[][]int{}

	sort.Ints(candidates)

	combination22(candidates, target, []int{}, choose, 0)
	// fmt.Printf("sums3 %v\n", *choose)

	// a := []int{1, 2}
	// testslice(a)
	// fmt.Printf("%v\n", a)

	return *choose
}

// var dept = 0

func combination22(candidates []int, target int, sums []int, choose *[][]int, ind int) bool {
	// dept++
	// defer func() {
	// 	dept--
	// }()
	sum := count(sums)
	if sum > target {
		return false
	}

	if sum == target {
		return true
	}

	for i := ind; i < len(candidates); i++ {
		// fmt.Println(dept)
		if i > ind && candidates[i] == candidates[i-1] {
			continue
		}

		sums = append(sums, candidates[i])
		rs := combination22(candidates, target, sums, choose, i+1)

		if rs {
			*choose = append(*choose, utils.DeepCopyIntSlice(sums))
		}
		sums = sums[:len(sums)-1]
	}

	return false
}
