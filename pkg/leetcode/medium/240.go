package medium

import (
	"sort"
)

/*
	240. Search a 2D Matrix II
*/

func searchMatrix(matrix [][]int, target int) bool {
	cal, row := 0, 0
	row = len(matrix)

	if len(matrix) != 0 {
		cal = len(matrix[0])
	}

	for i := 0; i < row; i++ {
		got := sort.Search(cal, func(j int) bool {
			return matrix[i][j] >= target
		})

		if got < cal && matrix[i][got] == target {
			return true
		}
	}

	return false
}
