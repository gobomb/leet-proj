package leetcode

func removeDuplicatesII(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}
	j := 0
	c := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			c++
		}
		if nums[i] != nums[i-1] {
			c = 0
		}
		if c >= 2 {
			continue
		}
		j++
		nums[j] = nums[i]
	}
	return j + 1
}
