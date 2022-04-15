package sort

func quick(arr []int) {
	QuickSort(arr, 0, len(arr)-1)
}

func QuickSort(arr []int, startIndex, endIndex int) {
	if startIndex >= endIndex {
		return
	}

	pivotIndex := partition(arr, startIndex, endIndex)
	// pivotIndex 的数字已经在正确的位置上了
	QuickSort(arr, startIndex, pivotIndex-1)
	QuickSort(arr, pivotIndex+1, endIndex)
}

// https://segmentfault.com/a/1190000020229121
func partition(arr []int, startIndex, endIndex int) int {
	var (
		i   = startIndex
		j   = endIndex
		piv = arr[startIndex]
	)

	// i+ j- 不会越界
	for i != j {
		// 先移动j
		for i < j && arr[j] > piv {
			j--
		}

		for i < j && arr[i] <= piv {
			i++
		}

		if i < j {
			swap(arr, i, j)
		}
	}

	//i == j
	swap(arr, startIndex, i)

	return i
}

func swap(nums []int, a, b int) {
	nums[a], nums[b] = nums[b], nums[a]
}
