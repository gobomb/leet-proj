package leetcode

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	// fmt.Println(nums)
	result := [][]int{}
	sum := 0
	start, end := 0, 0
	length := len(nums)

	for index := 1; index < length-1; index++ {
		// key
		start, end = 0, length-1

		if index > 1 && nums[index] == nums[index-1] {
			// ?
			start = index - 1
		}
		// fmt.Printf("%d\n", start)

		//for
		for start < index && index < end {
			// fmt.Println(start, index, end)
			// 去重
			if start > 0 && nums[start] == nums[start-1] {
				start++
				continue
			}
			if end < length-1 && nums[end] == nums[end+1] {
				end--
				continue
			}
			// fmt.Println(nums[start], nums[index], nums[end])
			sum = nums[index] + nums[start] + nums[end]
			// fmt.Printf("%d\n", sum)
			if sum == 0 {
				result = append(result, []int{nums[start], nums[index], nums[end]})
				start++
				end--
			} else if sum < 0 {
				start++
			} else {
				end--
			}
		}
	}
	return result
}

func Test3sum() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Printf("%#v \n", threeSum(nums))
	fmt.Printf("%#v \n", threeSum([]int{}))
	fmt.Printf("%#v \n", threeSum([]int{0}))
}
