package leetcode

import (
	"fmt"
	"math"
)

func MedianSortedArrays(nums1 []int, nums2 []int) float64 {
	x, y := len(nums1), len(nums2)

	if x > y {
		return MedianSortedArrays(nums2, nums1)
	}

	var mid1, mid2, m1left, m1right, m2left, m2right int

	// hi 相比一般坐标会 +1，因为由此除以2产生的中间坐标表示切分线，mid-1 | mid
	lo, hi := 0, x

	for lo <= hi {
		mid1 = lo + (hi-lo)>>1
		mid2 = (x+y+1)>>1 - mid1

		if mid1 == 0 {
			m1left = math.MinInt64
		} else {
			m1left = nums1[mid1-1]
		}
		if mid1 == x {
			m1right = math.MaxInt64
		} else {
			m1right = nums1[mid1]
		}

		if mid2 == 0 {
			m2left = math.MinInt64
		} else {
			m2left = nums2[mid2-1]
		}
		if mid2 == y {
			m2right = math.MaxInt64
		} else {
			m2right = nums2[mid2]
		}

		if m1left <= m2right && m2left <= m1right {
			if (x+y)&1 == 1 {
				return float64(max(m1left, m2left))
			}
			return float64(max(m1left, m2left)+min(m1right, m2right)) / 2
		} else if m1left > m2right {
			hi = mid1 - 1
		} else {
			lo = mid1 + 1
		}
	}
	return -1
}

// https://books.halfrost.com/leetcode/ChapterFour/0004.Median-of-Two-Sorted-Arrays/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// 假设 nums1 的长度小
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}
	low, high, k, nums1Mid, nums2Mid := 0, len(nums1), (len(nums1)+len(nums2)+1)>>1, 0, 0
	for low <= high {
		// nums1:  ……………… nums1[nums1Mid-1] | nums1[nums1Mid] ……………………
		// nums2:  ……………… nums2[nums2Mid-1] | nums2[nums2Mid] ……………………
		nums1Mid = low + (high-low)>>1                           // 分界限右侧是 mid，分界线左侧是 mid - 1
		nums2Mid = k - nums1Mid                                  // ???
		if nums1Mid > 0 && nums1[nums1Mid-1] > nums2[nums2Mid] { // nums1 中的分界线划多了，要向左边移动
			high = nums1Mid - 1
		} else if nums1Mid != len(nums1) && nums1[nums1Mid] < nums2[nums2Mid-1] { // nums1 中的分界线划少了，要向右边移动
			low = nums1Mid + 1
		} else {
			// 找到合适的划分了，需要输出最终结果了
			// 分为奇数偶数 2 种情况
			break
		}
	}
	midLeft, midRight := 0, 0
	if nums1Mid == 0 {
		midLeft = nums2[nums2Mid-1]
	} else if nums2Mid == 0 {
		midLeft = nums1[nums1Mid-1]
	} else {
		midLeft = max(nums1[nums1Mid-1], nums2[nums2Mid-1])
	}
	// odd
	if (len(nums1)+len(nums2))&1 == 1 {
		return float64(midLeft)
	}
	if nums1Mid == len(nums1) {
		midRight = nums2[nums2Mid]
	} else if nums2Mid == len(nums2) {
		midRight = nums1[nums1Mid]
	} else {
		midRight = min(nums1[nums1Mid], nums2[nums2Mid])
	}
	return float64(midLeft+midRight) / 2
}

func TestMedian() {
	a1 := []int{1, 3, 4, 7, 8}
	a2 := []int{2, 4, 6, 9, 10, 30}
	fmt.Println(MedianSortedArrays(a1, a2))
	fmt.Println(findMedianSortedArrays(a1, a2))

	b1 := []int{13, 132, 134}
	b2 := []int{}
	fmt.Println(MedianSortedArrays(b1, b2))
	fmt.Println(findMedianSortedArrays(b1, b2))
}
