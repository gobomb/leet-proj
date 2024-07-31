package review

// 704. Binary Search

func search(nums []int, target int) int {
	lo, hi, mi := 0, len(nums), 0
	for lo < hi {
		mi = lo + (hi-lo)/2

		if nums[mi] == target {
			hi = mi
		} else if nums[mi] > target {
			hi = mi
		} else if nums[mi] < target {
			lo = mi + 1
		}
	}

	if lo == len(nums) || nums[lo] != target {
		return -1
	}

	return lo
}
