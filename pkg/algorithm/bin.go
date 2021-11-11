package algorithm

import (
	"fmt"
	"log"
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

func BinSearch1(arr []int, k int) int {
	mi, lo, hi := 0, 0, len(arr)
	for lo < hi {
		// fmt.Println(arr[lo], arr[mi], arr[hi])
		// fmt.Println(" lo mid hi", lo, mi, hi)
		mi = lo + (hi-lo)>>1

		if arr[mi] >= k {
			hi = mi
		} else if arr[mi] < k {
			lo = mi + 1
		}
	}
	if lo < 0 || lo == len(arr) || arr[lo] != k {
		return -1
	}

	log.Println(lo, hi)
	return lo
}
