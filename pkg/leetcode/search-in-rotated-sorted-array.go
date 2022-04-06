package leetcode

func search2(arr []int, t int) int {
	var hi, lo, mi int
	lo = 0
	hi = len(arr) - 1
	// 两边都是开区间
	for lo <= hi {
		mi = lo + (hi-lo)/2

		// 必然是一边升序，一边不升序或重合
		if arr[mi] == t {
			// 先判断mi是否与t相等
			return mi
		} else if arr[lo] <= arr[mi] {
			// 左边有序 or 左边端点重合
			if arr[lo] <= t && t <= arr[mi] {
				// t落在左半边
				hi = mi - 1
			} else {
				lo = mi + 1
			}
		} else if arr[mi] <= arr[hi] {
			//右边有序 or 右边端点重合
			if arr[mi] <= t && t <= arr[hi] {
				// t落在右半边
				lo = mi + 1
			} else {
				hi = mi - 1
			}
		}

	}

	return -1
}

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
