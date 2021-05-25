package leetcode

func setZeroes(matrix [][]int) {
	ii, jj := false, false
	for i := 0; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			ii = true
			break
		}
	}

	for i := 0; i < len(matrix[0]); i++ {
		if matrix[0][i] == 0 {
			jj = true
			break
		}
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}

	for i := 0; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			if i == 0 {
				continue
			}
			for j := 0; j < len(matrix[i]); j++ {
				matrix[i][j] = 0
			}
		}
	}

	for i := 0; i < len(matrix[0]); i++ {
		if matrix[0][i] == 0 {
			if i == 0 {
				continue
			}
			for j := 1; j < len(matrix); j++ {
				matrix[j][i] = 0
			}
		}
	}

	if ii {
		for i := range matrix {
			matrix[i][0] = 0
		}
	}
	if jj {
		for i := range matrix[0] {
			matrix[0][i] = 0
		}
	}
	return
}
