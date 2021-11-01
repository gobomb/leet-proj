package leetcode

/*
	448. Find All Numbers Disappeared in an Array
*/

// 48 ms	7.4 MB
func findDisappearedNumbers(nums []int) []int {
	for i := range nums {
		tmp := nums[i] - 1

		for {
			if tmp < 0 {
				break
			}

			ttmp := tmp
			tmp, nums[ttmp] = nums[ttmp]-1, -1
		}
	}
	rs := nums[0:0]
	for i := range nums {
		if nums[i] >= 0 {
			rs = append(rs, i+1)
		}
	}

	return rs
}
