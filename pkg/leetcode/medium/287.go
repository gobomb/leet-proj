package medium

/*
	287. Find the Duplicate Number
*/

func findDuplicate(nums []int) int {
	var i, j int
	i = nums[0]
	j = nums[nums[0]]

	for i != j {
		i = nums[i]
		j = nums[nums[j]]
	}

	for j := 0; i != j; {
		i = nums[i]
		j = nums[j]
	}

	return i
}
