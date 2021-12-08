package sort

import "log"

func init() {
	log.SetFlags(log.Lshortfile)
}

// 非稳定
func heapSort(arr []int) {
	l := 1

	for ; l <= len(arr); l++ {
		up(arr[:l])
	}

	l -= 1
	i := 0

	for l != 0 {
		l--
		swap(arr, 0, l)
		i = 0

		for {
			child := i*2 + 1
			if child >= l {
				break
			}

			if child+1 < l && arr[child+1] > arr[child] {
				child++
			}

			if arr[child] > arr[i] {
				swap(arr, child, i)
				i = child
			} else {
				break
			}
		}
	}

}

func up(arr []int) {
	for i := len(arr) - 1; arr[i] > arr[i/2]; i /= 2 {
		swap(arr, i, i/2)
	}
}
