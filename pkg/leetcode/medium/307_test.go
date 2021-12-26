package medium

import "fmt"

func ExampleConstructorNumArray() {
	toTest := []func([]int) NumArrayMutable{ConstructorNumArray, ConstructorNumArray1}

	for _, f := range toTest {
		got := f([]int{1, 3, 5})
		fmt.Println(got.SumRange(0, 2))
		got.Update(1, 2)
		fmt.Println(got.SumRange(0, 2))
	}
	// OutPut:
	// 9
	// 8
	// 9
	// 8
}
