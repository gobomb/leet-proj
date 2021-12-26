package easy

import (
	"fmt"
)

func ExampleConstructorNumArray() {

	toTest := []func([]int) NumArray{Constructor3, Constructor31}

	for _, f := range toTest {
		got := f([]int{-2, 0, 3, -5, 2, -1})
		fmt.Println(got.SumRange(0, 2))
		fmt.Println(got.SumRange(2, 5))
		fmt.Println(got.SumRange(0, 5))
	}
	// OutPut:
	// 1
	// -1
	// -3
	// 1
	// -1
	// -3
}
