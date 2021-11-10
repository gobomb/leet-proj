package medium

/*
	153. Find Minimum in Rotated Sorted Array

*/

func findMin(nums []int) int {
	lo, hi, mi := 0, len(nums)-1, 0

	for lo <= hi {
		mi = (hi + lo) / 2

		switch {
		// 比最高的大，最小值在右边
		case nums[mi] > nums[hi]:
			lo = mi + 1
		// 比最低的小，最小值在左边或中间
		case nums[mi] < nums[lo]:
			// mi 刚好就是最小值
			if nums[mi] < nums[mi-1] && nums[mi] < nums[mi+1] {
				return nums[mi]
			}

			hi = mi - 1
		// 处于正常升序，直接返回最低位置
		case nums[lo] <= nums[mi] && nums[mi] <= nums[hi]:
			return nums[lo]
		}
	}

	return 0
}
