package leetcode

/*
1. 从尾部找到第一个不符合升序的数字a
2. 从尾部找到第一个比a大的数字b
3. 交换 a b 位置
4. 将 b（交换后的b位置）后的序列翻转
*/

func nextPermutation(nums []int) {
	var a, b, i int
	for i = len(nums) - 1; i > 0; i-- {
		if nums[i-1] < nums[i] {
			a = i - 1
			break
		}
	}
	// 从尾部算全部都是降序，已经是最大了。直接逆序返回
	if i == 0 {
		reverseInt(nums)
		return
	}

	for i := len(nums) - 1; i > a; i-- {
		if nums[a] < nums[i] {
			b = i
			break
		}
	}

	swap(nums, a, b)

	reverseInt(nums[a+1:])
}

func reverseInt(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		swap(nums, i, j)
	}
}

func swap(nums []int, a, b int) {
	nums[a], nums[b] = nums[b], nums[a]
}
