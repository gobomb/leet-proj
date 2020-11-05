package ds

import "fmt"

// 过大不存在的数/等于最后一位的数 会导致死循环
// lo + hi 可能会溢出
// 注意找不到时的结束条件
func Bin(arr []int, k int) int {
	mi, lo, hi := 0, 0, len(arr)-1

	for lo <= hi {
		mi = lo + (hi-lo)>>1
		fmt.Println("hi lo mid", hi, lo, mi)
		if arr[mi] == k {
			return mi
		} else if arr[mi] < k {
			lo = mi + 1
		} else {
			hi = mi - 1
		}
	}

	return -1

}

func bin2(arr []int, k int) int {
	lo := 0
	hi := len(arr) - 1
	mi := (lo + hi) / 2

	for i := mi; ; i = (lo + hi) / 2 {
		if arr[i] == k {
			return i
		}

		if arr[i] < k {
			lo = i
			continue
		}
		if arr[i] > k {
			hi = i
			continue
		}
	}
	return -1
}

func binarySearchMatrix(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)>>1

		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

func TestBin() {
	arr := []int{1, 12, 23, 34, 45}
	println(Bin(arr, 12))
	r := binarySearchMatrix(arr, 12)
	println(r)
}
