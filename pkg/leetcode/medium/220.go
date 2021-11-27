package medium

/*
	220. Contains Duplicate III
*/

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	for i := range nums {
		for j := i + 1; j < len(nums) && j-i <= k; j++ {
			if abs(nums[i]-nums[j]) <= t {
				return true
			}
		}
	}

	return false
}

func containsNearbyAlmostDuplicate1(nums []int, k int, t int) bool {
	if k == 0 {
		return false
	}

	buckets := map[int]int{}

	for i := range nums {
		// why t+1
		// 同一个桶都是符合条件的
		// t=3
		// (0,1,2,3)/4=0 3-0<=3
		// (4,5,6,7)/4=1 7-4<=3
		key := nums[i] / (t + 1)

		// why key--
		// 可能nums[i]为负数，比如-4 / 5 以及 4 / 5都等于0，所以负数要向下移动一位
		if nums[i] < 0 {
			key--
		}

		// log.Println(buckets, key, buckets[key-1], nums[i])

		// 如果在桶中则命中
		// 如果在相邻桶，判断一下
		if _, ok := buckets[key]; ok {
			return true
		} else if v, ok := buckets[key-1]; ok && abs(nums[i]-v) <= t {
			return true
		} else if v, ok := buckets[key+1]; ok && abs(nums[i]-v) <= t {
			return true
		}

		if len(buckets) >= k {
			delete(buckets, nums[i-k]/(t+1))
		}

		buckets[key] = nums[i]
	}

	return false
}
