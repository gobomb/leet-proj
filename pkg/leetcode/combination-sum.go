package leetcode

import (
	"fmt"
	"justest/pkg/utils"
	"reflect"
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
			*choose = append(*choose, utils.DeepCopyIntSlice(sums))
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
			return true
		}
		sum -= candidates[i]
	}
	return false
}

func compSliceSlice(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	cmap := make(map[int]bool)
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			if reflect.DeepEqual(a[i], b[j]) {
				_, ok := cmap[j]
				if ok {
					return false
				}
				cmap[j] = true

				break
			}
		}
	}
	return len(cmap) == len(a)
}
