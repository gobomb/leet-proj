package medium

/*
	229. Majority Element II
*/

// 25 ms	5.2 MB
func majorityElement(nums []int) []int {
	length := len(nums) / 3

	mp := make(map[int]int)

	rs := []int{}

	for i := range nums {
		mp[nums[i]]++

		if mp[nums[i]] == length+1 {
			rs = append(rs, nums[i])
		}
	}

	return rs
}

// 12 ms	5 MB
func majorityElement1(nums []int) []int {
	length := len(nums) / 3

	rs := []int{}
	a, b, c1, c2 := len(nums), len(nums), 0, 0

	for i := range nums {
		if nums[i] == a {
			c1++
		} else if nums[i] == b {
			c2++
		} else if c1 == 0 {
			a, c1 = nums[i], 1
		} else if c2 == 0 {
			b, c2 = nums[i], 1
		} else {
			c1--
			c2--
		}
	}

	c1, c2 = 0, 0

	for i := range nums {
		if nums[i] == a {
			c1++
		} else if nums[i] == b {
			c2++
		}
	}

	if c1 > length {
		rs = append(rs, a)
	}

	if c2 > length {
		rs = append(rs, b)
	}

	return rs
}
