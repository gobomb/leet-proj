package easy

/*
	217. Contains Duplicate
*/

func containsDuplicate(nums []int) bool {
	mp := make(map[int]struct{})

	for _, n := range nums {
		if _, ok := mp[n]; ok {
			return true
		}

		mp[n] = struct{}{}
	}

	return false
}
