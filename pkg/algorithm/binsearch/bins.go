package binsearch

import "log"

/*
	1. 左闭右开
	2. hi 不加减，lo=mi+1
	3. 根据情况，mi等于k时，选择移动hi或mi
*/

// 假如 1 2 3 为相同数字的下标 [1,4)
// 返回 1
// 返回左端闭区间下标
// 返回范围 0~len(arr)
func binSearchL(arr []int, k int) int {
	mi, lo, hi := 0, 0, len(arr)
	for lo < hi {
		mi = lo + (hi-lo)>>1

		if arr[mi] > k {
			hi = mi // arr[mi]!=k 因为 hi 是开区间，所以可以设置为 mi
		} else if arr[mi] < k {
			lo = mi + 1 // arr[mi]!=k 不需要再搜索，lo是闭区间，设置为mi+1
		} else if arr[mi] == k {
			hi = mi // 取决于要找左端点还是右端点，要找左端点，则右下标向左趋近；找右端点，左下标向右趋近
		}
	}

	// 循环结束 lo==hi
	return lo
}

// 假如 1 2 3 为相同数字的下标 [1,4)
// 返回 4
// 返回右端开区间下标
// 返回范围 0~len(arr)
func binSearchR(arr []int, k int) int {
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

	return lo
}

type T int

const (
	T1 T = iota
	T2
	T3
	T4
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func binAll(arr []int, k int, t T) int {
	var index int
	switch t {
	case T1:
		// 找相同数字的左侧边界
		index = binSearchL(arr, k)

		if index == len(arr) || arr[index] != k {
			return -1
		}

		return index
	case T2:
		// 找相同数字的右侧边界
		// 相当于大于target的最小下标-1
		index = binSearchR(arr, k)
		index--

		if index == -1 || arr[index] != k {
			return -1
		}

		return index
	case T3:
		// 找小于target的最大下标
		// 相当于相同数字的左侧边界--
		index = binSearchL(arr, k)

		index--

		return index

	case T4:
		// 找大于target的最小下标

		index = binSearchR(arr, k)

		if index == len(arr) {
			return -1
		}

		return index

	default:
		return -1
	}
}
