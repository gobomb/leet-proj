package leetcode

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
