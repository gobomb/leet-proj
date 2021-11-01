package leetcode

import "fmt"

func removeElement(nums []int, val int) int {
	j := 0
	l := len(nums)

	for i := 0; i < l; i++ {
		if nums[i] == val {
			continue
		}
		nums[j] = nums[i]
		j++
	}
	return j
}

func testremoveElement() {
	tests := [][]int{
		{1},
		{1, 1, 1, 3, 4, 4},
		{1, 1, 2},
		{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
		{3, 2, 2, 3},
		{0, 1, 2, 2, 3, 0, 4, 2},
	}
	for i := range tests {
		r := removeElement(tests[i], 3)
		fmt.Printf("%+v\n", tests[i])
		fmt.Printf("result %v\n", r)
		fmt.Printf("%+v\n\n", tests[i][:r])
	}
}
