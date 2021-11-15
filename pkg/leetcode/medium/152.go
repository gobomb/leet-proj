package medium

/*
	152. Maximum Product Subarray
*/

/*
	https://books.halfrost.com/leetcode/ChapterFour/0100~0199/0152.Maximum-Product-Subarray/

	如果最后一个数是负数，最大值就在负数 * 最小值中产生，如果最后一个数是正数，最大值就在正数 * 最大值中产生

	https://segmentfault.com/a/1190000018154455
*/
func maxProduct(nums []int) int {
	maximum, minimum, res := nums[0], nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		pmaximum := maximum
		pminimum := minimum
		maximum = max(max(nums[i], maximum*nums[i]), pminimum*nums[i])
		minimum = min(min(nums[i], minimum*nums[i]), pmaximum*nums[i])
		res = max(maximum, res)
	}

	return res
}
