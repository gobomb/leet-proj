package leetcode

func sortColors(nums []int) {
	// 计数排序
	bucket := make([]int, 3)

	for i := 0; i < len(nums); i++ {
		bucket[nums[i]] += 1
	}

	sorti := 0

	for i := 0; i < 3; i++ {
		for bucket[i] > 0 {
			nums[sorti] = i
			sorti += 1
			bucket[i] -= 1
		}
	}
}
