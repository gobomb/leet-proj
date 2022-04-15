package sort

import (
	"container/heap"
	"log"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

type h struct {
	a []int
}

// Len is the number of elements in the collection.
func (a *h) Len() int {
	return len(a.a)
}

func (a *h) Less(i int, j int) bool {
	return a.a[i] > a.a[j]
}

// Swap swaps the elements with indexes i and j.
func (a *h) Swap(i int, j int) {
	a.a[i], a.a[j] = a.a[j], a.a[i]
}

func (a *h) Push(x interface{}) {
	a.a = append(a.a, x.(int))
}

func (a *h) Pop() interface{} {
	x := a.a[len(a.a)-1]
	a.a = a.a[:len(a.a)-1]
	return x
}

func heapSortSTL(arr []int) {
	a := &h{
		a: arr,
	}

	heap.Init(a)

	for a.Len() != 0 {
		heap.Pop(a)
	}
}

// O(n)
func buildHeap(arr []int) {
	for i := (len(arr) - 2) / 2; i >= 0; i-- {
		down(arr, i, len(arr))
	}
}

// O(nlog(n))
func buildHeap2(arr []int) {
	l := 1
	for ; l <= len(arr); l++ {
		up(arr[:l])
	}
}

// 非稳定
func heapSort(arr []int) {
	// 建堆，升序排序使用大顶堆
	buildHeap(arr)

	l := len(arr)

	for l != 0 {
		l--
		// 交换最后叶子和根，将最大值放到最后
		swap(arr, 0, l)
		// 调整堆
		down(arr, 0, l)
	}
}

// O(log(n))
func down(arr []int, i, l int) {
	// 调整堆 down
	for {
		// 左节点
		child := i*2 + 1
		if child >= l {
			break
		}

		if child+1 < l && arr[child+1] > arr[child] {
			// 与右节点比较，取比较大的一个
			child++
		}

		if arr[child] > arr[i] {
			// 与父节点比较，如果比较大就交换
			swap(arr, child, i)
			// 继续检查子节点
			i = child
		} else {
			break
		}
	}
}

func up(arr []int) {
	// (i-1)/2取父节点
	for i := len(arr) - 1; arr[i] > arr[(i-1)/2]; i = (i - 1) / 2 {
		swap(arr, i, (i-1)/2)
	}
}
