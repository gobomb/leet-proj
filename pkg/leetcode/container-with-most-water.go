package leetcode

import "fmt"

func testMaxArea() {
	tests := [][]int{
		{1, 8, 6, 2, 5, 4, 8, 3, 7},
		{1, 1},
		{4, 3, 2, 1, 4},
		{1, 2, 1},
	}
	for i := range tests {
		r := maxArea(tests[i])
		fmt.Println(r)
	}
}

func maxArea(height []int) int {
	var maxA = 0
	var less, sub int

	lengh := len(height)
	j := lengh - 1
	for i := 0; i < j; {
		sub = j - i
		if height[i] > height[j] {
			less = height[j]
			j--
		} else {
			less = height[i]
			i++
		}
		area := less * sub
		if area > maxA {
			maxA = area
		}
	}
	return maxA
}
