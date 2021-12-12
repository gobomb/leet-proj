package golang

import (
	"reflect"
	"sort"
	"unsafe"
)

var a = []float64{4, 2, 5, 7, 2, 1, 88, 1, 1}
var b = []float64{4, 2, 5, 7, 2, 1, 88, 1, 1}

// 《Go语言高级编程》1.3
func SortFloat64FastV1(a []float64) {
	p := unsafe.Pointer(&a[0])
	var b []int = ((*[1 << 20]int)(p))[:len(a):cap(a)]
	sort.Ints(b)
	// fmt.Println(a, b)
}

func SortFloat64FastV2(a []float64) {
	var c []int

	aHdr := (*reflect.SliceHeader)(unsafe.Pointer(&a))
	cHdr := (*reflect.SliceHeader)(unsafe.Pointer(&c))

	*cHdr = *aHdr

	sort.Ints(c)
	// fmt.Println(a, c)
}
