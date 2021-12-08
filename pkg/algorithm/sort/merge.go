package sort

func mergeSort(arr []int) {
	tmp := mergeSort1(arr)
	copy(arr, tmp)
}

func mergeSort1(arr []int) []int {
	length := len(arr)

	if length <= 1 {
		return arr
	}

	mid := length / 2
	l := mergeSort1(arr[0:mid])
	r := mergeSort1(arr[mid:length])

	return mergeTwoSortList(l, r)
}

func mergeTwoSortList(a, b []int) []int {
	tmp := []int{}
	i, j := 0, 0

	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			tmp = append(tmp, a[i])
			i++
		} else {
			tmp = append(tmp, b[j])
			j++
		}
	}

	if i < len(a) {
		tmp = append(tmp, a[i:]...)
	}

	if j < len(b) {
		tmp = append(tmp, b[j:]...)
	}

	return tmp
}
