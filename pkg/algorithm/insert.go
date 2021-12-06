package algorithm

func insertSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		cur := arr[i]
		j := i - 1

		for ; j >= 0 && arr[j] > cur; j-- {
			arr[j+1] = arr[j]
		}

		arr[j+1] = cur
	}
}
