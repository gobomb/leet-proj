package medium

/*
	260. Single Number III
*/
func singleNumberIII(nums []int) []int {
	xor := nums[0]

	// 全部异或
	// xor 中 1 表示两数该位一个为1一个为0
	for i := 1; i < len(nums); i++ {
		xor ^= nums[i]
	}

	// 取反加一获得负数
	// 与负数相与获得最后一位1
	// xor & (~(xor-1))
	bitFlag := xor & (-xor)

	res := []int{xor, xor}

	for i := 0; i < len(nums); i++ {
		// 利用 bitFlag 分成两组，不同的两个数字各自在两组中
		// 然后用 136 的方式找到两数
		if (bitFlag^nums[i])&bitFlag == bitFlag {
			res[0] ^= nums[i]
		} else {
			res[1] ^= nums[i]
		}
	}

	return res
}
