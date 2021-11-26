package easy

/*
	219. Contains Duplicate II
*/

func containsNearbyDuplicate(nums []int, k int) bool {
	mp := make(map[int]int)
	found := -1

	for i := range nums {
		if ii, ok := mp[nums[i]]; ok {
			if i-ii <= k {
				found = i
				mp[nums[i]] = i

				if found != -1 && nums[found] != nums[i] {
					return false
				}
			} else {
				mp[nums[i]] = i
			}
		} else {
			mp[nums[i]] = i
		}
	}

	return found != -1
}
