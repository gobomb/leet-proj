package sort

func getMaxMin(arr []int) (int, int) {
	maxInt := arr[0]
	minInt := arr[0]

	for i := range arr {
		if arr[i] > maxInt {
			maxInt = arr[i]
		}
		if arr[i] < minInt {
			minInt = arr[i]
		}
	}

	return maxInt, minInt
}

type bucket interface {
	makeBucket() []int
	getArr() []int
	getKey(int) int
}

func countSortFunc(bc bucket) {
	bucket := bc.makeBucket()
	arr := bc.getArr()

	for i := range arr {
		key := bc.getKey(i)
		bucket[key]++
	}

	for i := 1; i < len(bucket); i++ {
		bucket[i] += bucket[i-1]
	}

	sorted := make([]int, len(arr))

	for i := len(arr) - 1; i >= 0; i-- {
		key := bc.getKey(i)
		sorted[bucket[key]-1] = arr[i]
		bucket[key]--
	}

	copy(arr, sorted)
}

type countSorting struct {
	arr            []int
	minInt, maxInt int
}

func getCountSorting(arr []int) bucket {
	maxInt, minInt := getMaxMin(arr)

	return &countSorting{
		arr:    arr,
		maxInt: maxInt,
		minInt: minInt,
	}
}

func (c *countSorting) makeBucket() []int {
	return make([]int, c.maxInt-c.minInt+1)
}
func (c *countSorting) getArr() []int {
	return c.arr
}
func (c *countSorting) getKey(i int) int {
	// log.Println(c.arr[i] - c.minInt)
	return c.arr[i] - c.minInt
}

// 计数排序
// 为稳定排序
// 当 maxInt>>O(n) 则不适用
// 时间复杂度为O(n)
func countSort(arr []int) {
	cs := getCountSorting(arr)
	countSortFunc(cs)
}

// 基数排序
func radixSort(arr []int) {
	maxInt, _ := getMaxMin(arr)

	for bit := 1; maxInt/bit > 0; bit *= 10 {
		length := len(arr)

		bucket := make([]int, 10)

		for i := 0; i < length; i++ {
			key := (arr[i] / bit) % 10

			bucket[key]++
		}

		for j := 1; j < len(bucket); j++ {
			bucket[j] += bucket[j-1]
		}

		res := make([]int, length)

		for i := length - 1; i >= 0; i-- {
			key := (arr[i] / bit) % 10
			res[bucket[key]-1] = arr[i]

			bucket[key]--
		}

		copy(arr, res)
	}
}

// 桶排序
func bucketSort(arr []int) {
	maxInt, _ := getMaxMin(arr)
	buckets := make([][]int, len(arr))

	index := 0
	for i := 0; i < len(arr); i++ {
		index = arr[i] * (len(arr) - 1) / maxInt
		// log.Println(index,maxInt)
		buckets[index] = append(buckets[index], arr[i])
	}
	// fmt.Println(buckets)

	tmppose := 0
	for i := 0; i < len(arr); i++ {
		bucketslen := len(buckets[i])

		if bucketslen > 0 {
			selectSort(buckets[i])
			copy(arr[tmppose:], buckets[i])
			tmppose += bucketslen
		}
	}
}
