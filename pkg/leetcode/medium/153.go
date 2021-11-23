package medium

/*
	153. Find Minimum in Rotated Sorted Array

*/

func findMin(nums []int) int {
	lo, hi, mi := 0, len(nums)-1, 0

	for lo <= hi {
		mi = lo + (hi-lo)/2

		switch {
		// 处于正常升序，直接返回最低位置
		// 需要加等号防止无限循环
		case nums[lo] <= nums[hi]:
			return nums[lo]
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
		}
	}

	return 0
}

func findMin1(nums []int) int {
	lo, hi, mi := 0, len(nums)-1, 0

	for lo < hi {
		mi = lo + (hi-lo)/2

		if nums[lo] < nums[hi] {
			return nums[lo]
		} else if nums[mi] < nums[lo] {
			hi = mi
		} else if nums[mi] >= nums[lo] {
			lo = mi + 1
		}
	}

	return nums[lo]
}

/*
	154. Find Minimum in Rotated Sorted Array II
*/

func findMinII(nums []int) int {
	lo, hi, mi := 0, len(nums)-1, 0

	for lo < hi {
		mi = lo + (hi-lo)/2

		switch {
		// 处于正常升序，直接返回最低位置
		case nums[lo] < nums[hi]:
			return nums[lo]
		// 比最高的大，最小值在右边
		case nums[mi] > nums[hi]:
			lo = mi + 1
		// 最小值在左边或中间,hi 向左逼近
		case nums[mi] < nums[hi]:
			hi = mi
		case nums[mi] == nums[hi]:
			hi--
		}
	}

	return nums[lo]
}

// 剑指 offer
func findMinII1(nums []int) int {
	lo, hi, mi := 0, len(nums)-1, 0

	for lo < hi {
		// 保证hi在右半部分，lo在左半部分
		if hi-lo == 1 {
			mi = hi
			break
		}
		mi = lo + (hi-lo)/2
		// 处于正常升序，直接返回最低位置
		if nums[lo] < nums[hi] {
			mi = lo
			break
		} else if nums[mi] > nums[hi] {
			// 比最高的大，最小值在右边
			lo = mi
		} else if nums[mi] < nums[lo] {
			// 最小值在左边,hi 向左逼近
			hi = mi
		} else if nums[mi] == nums[hi] {
			// 退化成线性查找
			hi--
		}
	}

	return nums[mi]
}
