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

func moveZeroes1(nums []int) {
	for i := range nums {
		if nums[i] == 0 {
			for j := i + 1; j < len(nums); j++ {
				if nums[j] != 0 {
					nums[i], nums[j] = nums[j], nums[i]

					break
				}
			}
		}
	}
}
