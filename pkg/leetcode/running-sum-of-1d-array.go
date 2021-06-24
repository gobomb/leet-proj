package leetcode

func runningSum(nums []int) []int {
	var sum int
	for i := range nums {
		sum += nums[i]
		nums[i] = sum
	}
	return nums
}
