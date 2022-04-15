package review

import "log"

/*
	215. Kth Largest Element in an Array
*/

// O(n)
func findKthLargestQuickSelect(nums []int, k int) int {
	if k <= 0 || k > len(nums) {
		return -1
	}

	// 由小到大排序，需要转换下标
	k1 := len(nums) - k

	return qselect(nums, k1, 0, len(nums)-1)
}

func qselect(arr []int, k, left, right int) int {
	// pivot := left
	// index := pivot + 1

	// for i := index; i <= right; i++ {
	// 	if arr[i] < arr[pivot] {
	// 		swap(arr, i, index)
	// 		index += 1
	// 	}
	// }

	// // 大于等于 index 的都大于等于 piv
	// swap(arr, pivot, index-1)

	// i := index - 1

	i := partition(arr, left, right)

	if k < i {
		return qselect(arr, k, left, i-1)
	} else if k > i {
		return qselect(arr, k, i+1, right)
	}

	return arr[i]
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func partition(nums []int, st, ed int) int {
	// select a piv
	piv := nums[st]
	i := st
	j := ed

	// 4 5

	for i != j {
		for i < j && nums[j] > piv {
			j--
		}

		for i < j && nums[i] <= piv {
			i++
		}

		if i < j {
			swap(nums, i, j)
		}

	}

	swap(nums, st, i) // runtime.Goexit()

	return i
}

// 快速选择和优先队列的时间复杂度
// 堆排序的时间复杂度
// O(klog(n))
func findKthLargestHeap(arr []int, k0 int) int {
	if k0 <= 0 || k0 > len(arr) {
		return -1
	}

	// 由小到大排序，需要转换下标
	k := len(arr) - k0

	// 遍历 up，建堆
	// O(klog(n))
	for i := 0; i <= k; i++ {
		up(arr, i)
	}

	log.Println(arr, k)

	// 最后下标与堆顶交换，并调整堆

	// O(klog(n))
	for i := k + 1; i < len(arr); i++ {
		if arr[i] > arr[0] {
			continue
		}

		swap(arr, i, 0)
		log.Println(arr, k, i)

		// down

		p := 0
		for p <= k {
			// 判断左右子树
			son := p*2 + 1

			if son > k {
				break
			}

			if son+1 <= k && arr[son+1] > arr[son] {
				son++
			}

			// 左右子树与父节点比较，然后交换父节点和子树
			if arr[p] < arr[son] {
				swap(arr, p, son)
			}

			p = son
		}
	}

	return arr[0]
}

func up(arr []int, r int) {
	for i := r; arr[i] > arr[(i-1)/2]; i = (i - 1) / 2 {
		swap(arr, i, (i-1)/2)
	}
}
