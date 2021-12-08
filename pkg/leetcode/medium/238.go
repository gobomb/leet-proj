package medium

/*
	238. Product of Array Except Self
*/

func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	res[0] = 1

	for i := 1; i < len(nums); i++ {
		res[i] = res[i-1] * nums[i-1]
	}

	rightProduct := nums[len(nums)-1]

	for i := len(nums) - 2; i >= 0; i-- {
		res[i] *= rightProduct
		rightProduct = rightProduct * nums[i]
	}

	return res
}
