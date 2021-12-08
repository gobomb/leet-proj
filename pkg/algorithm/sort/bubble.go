package sort

func bubSort(arr []int) {
	for j := len(arr) - 1; j > 0; j-- {
		for i := 0; i < j; i++ {
			if arr[i] > arr[i+1] {
				swap(arr, i, i+1)
			}
		}
	}
	// log.Println(arr)
}
