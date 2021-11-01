package leetcode

func searchRotated(nums []int, target int) bool {
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		mid := (lo + hi) / 2

		if nums[mid] == target {
			return true
		} else if nums[mid] > nums[lo] {
			if nums[lo] == target {
				return true
			}
			if nums[lo] < target && target < nums[mid] {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		} else if nums[mid] < nums[hi] {
			if nums[hi] == target {
				return true
			}
			if nums[mid] < target && target < nums[hi] {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		} else {
			if nums[mid] == nums[lo] {
				lo++
			}
			if nums[mid] == nums[hi] {
				hi--
			}
		}
	}
	return false
}
