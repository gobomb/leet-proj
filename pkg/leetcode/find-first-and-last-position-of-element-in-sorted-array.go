package leetcode

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	return []int{findfirstmatch(nums, target), findlastmatch(nums, target)}
}

func findlastmatch(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		mi := (lo + hi) / 2
		if nums[mi] == target {
			if (mi == len(nums)-1) || (nums[mi+1] != target) {
				return mi
			}
			lo = mi + 1
		} else if target < nums[mi] {
			hi = mi - 1
		} else if target > nums[mi] {
			lo = mi + 1
		}
	}
	return -1
}

func findfirstmatch(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		mi := (lo + hi) / 2
		if nums[mi] == target {
			if (mi == 0) || (nums[mi-1] != target) {
				return mi
			}
			hi = mi - 1
		} else if target < nums[mi] {
			hi = mi - 1
		} else if target > nums[mi] {
			lo = mi + 1
		}
	}
	return -1
}
