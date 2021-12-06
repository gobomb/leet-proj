package algorithm

func quick(arr []int) {
	QuickSort(arr, 0, len(arr)-1)
}

func QuickSort(arr []int, startIndex, endIndex int) {
	if startIndex >= endIndex {
		return
	}

	pivotIndex := partition(arr, startIndex, endIndex)
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

	for i != j {
		// <= ;顺序
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
	swap(arr, startIndex, i)

	return i
}

func swap(nums []int, a, b int) {
	nums[a], nums[b] = nums[b], nums[a]
}
