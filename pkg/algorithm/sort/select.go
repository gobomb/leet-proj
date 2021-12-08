package sort

func selectSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		minInd := i
		for j := i + 1; j < len(arr); j++ {
			if arr[minInd] > arr[j] {
				minInd = j
			}
		}
		swap(arr, i, minInd)
	}
}
