package algorithm

import (
	"fmt"
)

// https://mp.weixin.qq.com/s/uA2suoVykENmCQcKFMOSuQ
// 过大不存在的数/等于最后一位的数 会导致死循环
// lo + hi 可能会溢出
// 注意找不到时的结束条件
func BinSearch(arr []int, k int) int {
	mi, lo, hi := 0, 0, len(arr)-1

	for lo <= hi { // <---- 需要包含相等情况
		mi = lo + (hi-lo)>>1
		fmt.Println("hi lo mid", hi, lo, mi)
		if arr[mi] == k {
			return mi
		} else if arr[mi] < k {
			lo = mi + 1 // <---- lo/hi 变化需要+-1
		} else {
			hi = mi - 1 // <----
		}
	}

	return -1
}

// https://labuladong.gitee.io/algo/1/9/
// 找相同数字的左侧边界
func BinSearch1(arr []int, k int) int {
	mi, lo, hi := 0, 0, len(arr)
	for lo < hi {
		mi = lo + (hi-lo)>>1

		if arr[mi] > k {
			hi = mi
		} else if arr[mi] < k {
			lo = mi + 1
		} else if arr[mi] == k {
			hi = mi
		}
	}

	if lo == len(arr) || arr[lo] != k {
		return -1
	}

	return lo
}

// 找相同数字的右侧边界
func BinSearch2(arr []int, k int) int {
	mi, lo, hi := 0, 0, len(arr)
	for lo < hi {
		mi = lo + (hi-lo)>>1

		if arr[mi] > k {
			hi = mi
		} else if arr[mi] < k {
			lo = mi + 1
		} else if arr[mi] == k {
			lo = mi + 1
		}
	}

	if arr[lo-1] != k {
		return -1
	}

	return lo - 1
}

// 找大于k的最小数字下标
func BinSearch3(arr []int, k int) int {
	mi, lo, hi := 0, 0, len(arr)
	for lo < hi {
		mi = lo + (hi-lo)>>1

		if arr[mi] > k {
			hi = mi
		} else if arr[mi] < k {
			lo = mi + 1
		} else if arr[mi] == k {
			lo = mi + 1
		}
	}

	if lo >= len(arr) {
		return -1
	}

	return lo
}

// 找小于k的最大数字下标
func BinSearch4(arr []int, k int) int {
	mi, lo, hi := 0, 0, len(arr)
	for lo < hi {
		mi = lo + (hi-lo)>>1

		if arr[mi] > k {
			hi = mi
		} else if arr[mi] < k {
			lo = mi + 1
		} else if arr[mi] == k {
			hi = mi
		}
	}

	if lo > len(arr) {
		return -1
	}

	return hi - 1
}
