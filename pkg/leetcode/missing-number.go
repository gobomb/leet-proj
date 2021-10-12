package leetcode

/*
	268. Missing Number
*/

// 16 ms	6.3 MB
func missingNumber(nums []int) int {
	l := len(nums)
	rs := (l + 1) * l / 2
	sum := 0

	for i := 0; i < l; i++ {
		sum += nums[i]
	}
	return rs - sum
}

// 36 ms	6.7 MB
func missingNumber2(nums []int) int {
	rs := 0
	// for i := 1; i <= len(nums); i++ {
	// 	rs += i
	// }
	for i := 0; i < len(nums); i++ {
		rs += i
		rs -= nums[i]
	}
	rs += len(nums)
	return rs
}
