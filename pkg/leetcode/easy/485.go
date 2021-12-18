package easy

/*
	485. Max Consecutive Ones
*/

func findMaxConsecutiveOnes(nums []int) int {
	c := 0
	maxC := 0

	for i := range nums {
		if nums[i] == 1 {
			c++

			if c > maxC {
				maxC = c
			}
		} else {
			c = 0
		}
	}

	return maxC
}
