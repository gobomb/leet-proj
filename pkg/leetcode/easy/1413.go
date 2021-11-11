package easy

func minStartValue(nums []int) int {
	for i := 1; ; i++ {
		sum := i
		j := 0

		for j = 0; j < len(nums); j++ {
			sum += nums[j]

			if sum < 1 {
				i -= sum
				break
			}
		}

		if j == len(nums) {
			return i
		}
	}
}
