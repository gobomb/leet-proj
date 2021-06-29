package leetcode

func shuffle(nums []int, n int) []int {
	rs := make([]int, len(nums))
	for i, j := 0, n; i < n; i, j = i+1, j+1 {
		rs[i*2] = nums[i]
		rs[i*2+1] = nums[j]
	}
	return rs
}
