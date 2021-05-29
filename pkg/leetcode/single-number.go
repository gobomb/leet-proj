package leetcode

func singleNumber(nums []int) int {
	mp := make(map[int]struct{})
	for i := range nums {
		_, ok := mp[nums[i]]
		if ok {
			delete(mp, nums[i])
		} else {
			mp[nums[i]] = struct{}{}
		}
	}
	for k := range mp {
		return k
	}
	return 0
}

func sigleNumberXor(nums []int) int {
	rs := 0
	for i := range nums {
		rs ^= nums[i]
	}
	return rs
}

func singleNumberXor2(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		res ^= nums[i]
	}
	return res
}
