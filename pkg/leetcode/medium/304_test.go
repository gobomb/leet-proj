package medium

import (
	"fmt"
)

func ExampleConstructorMatrix() {
	matrix := [][]int{{3, 0, 1, 4, 2}, {5, 6, 3, 2, 1}, {1, 2, 0, 1, 5}, {4, 1, 0, 1, 7}, {1, 0, 3, 0, 5}}

	counts := [][]int{{2, 1, 4, 3}, {1, 1, 2, 2}, {1, 2, 2, 4}}

	numMatrix := ConstructorMatrix(matrix)

	for _, c := range counts {
		fmt.Println(numMatrix.SumRegion(c[0], c[1], c[2], c[3]))
	}
	// OutPut: 8
	// 11
	// 12
}
