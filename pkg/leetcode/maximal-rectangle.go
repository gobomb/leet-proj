package leetcode

/*
	85. Maximal Rectangle
*/

func maximalRectangle(matrix [][]byte) int {
	max := 0
	rows := len(matrix)
	if rows == 0 {
		return max
	}
	nums := make([]int, len(matrix[0]))

	for i := 0; i < rows; i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] != '0' {
				nums[j] = nums[j] + 1
			} else {
				nums[j] = 0
			}
		}
		maxarea := largestRectangleArea2(nums)
		if maxarea > max {
			max = maxarea
		}
	}
	return max
}
