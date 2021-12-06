package algorithm

func shellSort(arr []int) {
	for gap := len(arr) / 2; gap > 0; gap /= 2 {
		for i := 0; i < len(arr); i += gap {
			cur := arr[i]
			j := i - gap
			for ; j >= 0; j -= gap {
				if arr[j] > cur {
					arr[j+gap] = arr[j]
				} else {
					break
				}
			}
			arr[j+gap] = cur
		}
	}
}
