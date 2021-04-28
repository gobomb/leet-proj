package leetcode

import (
	"fmt"
	"justest/pkg/ds"
)

func combinationSum(candidates []int, target int) [][]int {
	choose := new([][]int)
	combination2(candidates, target, []int{}, choose, 0)
	// fmt.Printf("%+v\n", len(*choose))

	// for i := 0; i < len(*choose); i++ {
	// 	fmt.Printf("%+v\n", (*choose)[i])
	// }
	return *choose
}

func count(sums []int) (cs int) {
	for i := 0; i < len(sums); i++ {
		cs += sums[i]
	}
	return
}

/*
func DeepCopyIntSlice(arr []int) []int {
	rcopy := make([]int, len(arr))
	copy(rcopy, arr)
	return rcopy
}
*/

func combination2(candidates []int, target int, sums []int, choose *[][]int, ind int) bool {
	sum := count(sums)
	if sum > target {
		return false
	}

	if sum == target {
		return true
	}

	for i := ind; i < len(candidates); i++ {
		sums = append(sums, candidates[i])
		rs := combination2(candidates, target, sums, choose, i)
		if rs {
			*choose = append(*choose, ds.DeepCopyIntSlice(sums))
		}
		sums = sums[:len(sums)-1]
	}

	return false
}

// 只找一个解
func combination(candidates []int, target int, sum int, choose *[]*[]int) bool {
	if sum > target {
		return false
	}
	if sum == target {
		*choose = append(*choose, &[]int{})
		return true
	}

	for i := 0; i < len(candidates); i++ {
		sum += candidates[i]

		if combination(candidates, target, sum, choose) {
			*(*choose)[len(*choose)-1] = append(*(*choose)[len(*choose)-1], candidates[i])
			fmt.Printf("sum %v i %v\n", sum, candidates[i])
			sum -= candidates[i]
			return true
		}
		sum -= candidates[i]

	}
	return false
}
