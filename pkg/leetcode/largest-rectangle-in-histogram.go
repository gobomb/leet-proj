package leetcode

/*

	84. Largest Rectangle in Histogram

*/

// brute force
func largestRectangleArea(heights []int) int {
	rs := 0
	h := 0
	for i := range heights {
		h = heights[i]
		for j := i + 1; j < len(heights); j++ {
			if heights[j] < h {
				h = heights[j]
			}
			area := h * (j - i + 1)
			if area > rs {
				rs = area
			}
		}
	}
	return rs
}

// stack
// ref: https://mp.weixin.qq.com/s/l08qmMpFM4uqRYr7cm_0QQ
func largestRectangleArea2(heights []int) int {
	var stack []int
	heights = append(heights, -1)
	rs := 0
	i := 0
	for i < len(heights) {
		if len(stack) == 0 ||
			heights[stack[len(stack)-1]] <= heights[i] {
			stack = append(stack, i)
			i++
		} else {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1] // Pop

			area := 0

			if len(stack) == 0 {
				area = i * heights[top]
			} else {
				area = (i - stack[len(stack)-1] - 1) * heights[top]
			}
			if area > rs {
				rs = area
			}

		}
	}
	return rs
}
