package easy

/*
	219. Contains Duplicate II
*/

func containsNearbyDuplicate(nums []int, k int) bool {
	mp := make(map[int]int)

	for i := range nums {
		if ii, ok := mp[nums[i]]; ok {
			if i-ii <= k {
				mp[nums[i]] = i

				return true
			}
		}

		mp[nums[i]] = i
	}

	return false
}
