package leetcode

func searchInsert(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			return i
		}
		if nums[i] > target {
			return i
		}
	}

	return len(nums)
}

func searchInsertBin(nums []int, target int) int {
	lo, hi, mi := 0, len(nums)-1, 0
	for lo <= hi {
		mi = lo + (hi-lo)>>1
		if nums[mi] == target {
			return mi
		}
		if nums[mi] < target {
			lo = mi + 1
		}
		if nums[mi] > target {
			if mi == 0 || nums[mi-1] < target {
				return mi
			}
			hi = mi - 1
		}
	}

	return len(nums)
}
