package easy

/*
	283. Move Zeroes
*/

func moveZeroes(nums []int) {
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] == 0 {
			copy(nums[i:], nums[i+1:])
			nums[len(nums)-1] = 0
		}
	}
}
