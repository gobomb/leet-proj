package leetcode

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		mid := (lo + hi) / 2
		if target == nums[mid] {
			return mid
		} else if nums[mid] > nums[lo] {
			if nums[lo] <= target && target < nums[mid] {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		} else if nums[mid] < nums[hi] {
			if nums[mid] < target && target <= nums[hi] {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		} else {
			if nums[lo] == nums[mid] {
				lo++
			}
			if nums[hi] == nums[mid] {
				hi--
			}
		}
		// fmt.Printf("%v %v %v\n", lo, mid, hi)
	}

	return -1
}
