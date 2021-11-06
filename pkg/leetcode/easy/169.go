package easy

/*
	169. Majority Element
*/

// https://www.zhihu.com/question/49973163/answer/235921864
// Runtime: 16 ms, faster than 93.66% of Go online submissions for Majority Element.
// Memory Usage: 6.1 MB, less than 66.67% of Go online submissions for Majority Element.
func majorityElement(nums []int) int {
	major, count := 0, 0

	for i := 0; i < len(nums); i++ {
		if count == 0 {
			major = nums[i]
		}

		if nums[i] != major {
			count--
		} else {
			count++
		}
	}

	return major
}
