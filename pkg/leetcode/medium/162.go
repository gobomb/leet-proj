package medium

/*
	162. Find Peak Element
*/

// https://books.halfrost.com/leetcode/ChapterFour/0100~0199/0162.Find-Peak-Element/
func findPeakElement(nums []int) int {
	lo, hi, mi := 0, len(nums)-1, 0

	for lo < hi {
		mi = lo + (hi-lo)>>1
		if nums[mi] > nums[mi+1] {
			hi = mi
		} else {
			lo = mi + 1
		}
	}

	return hi
}
