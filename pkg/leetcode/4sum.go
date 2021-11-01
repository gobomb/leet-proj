package leetcode

import (
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	ma := make(map[[4]int]bool)

	// var sum int
	var rss [][]int
	// var rs []int
	for i := 0; i < len(nums)-3; i++ {
		subtarget := target - nums[i]
		subrss := threeSumTarget(nums[i+1:], subtarget)
		for j := range subrss {
			subrss[j][0] = nums[i]

			if !ma[subrss[j]] {
				ma[subrss[j]] = true
				rss = append(rss, subrss[j][:])
			}
		}
	}
	// fmt.Printf("map: %+v\n", ma)

	return rss
}

func threeSumTarget(nums []int, target int) [][4]int {
	result := [][4]int{}
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
			if sum == target {
				result = append(result, [4]int{-1, nums[start], nums[index], nums[end]})
				// println(start, index, end)

				start++
				end--
			} else if sum < target {
				start++
			} else {
				end--
			}
		}
	}
	return result
}
