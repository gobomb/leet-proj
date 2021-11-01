package leetcode

func rotate(matrix [][]int) {
	for ii := 0; ii < len(matrix)/2; ii++ {
		for jj := ii; jj < len(matrix[ii])-1-ii; jj++ {
			// move(matrix, 0, ii, jj)
			i := ii
			j := jj
			lasttmp := matrix[i][j]

			for t := 0; t < 4; t++ {
				matrix[j][len(matrix[i])-1-i], lasttmp = lasttmp, matrix[j][len(matrix[i])-1-i]

				//  注意这里 i 会变
				ti := i
				i = j
				j = len(matrix[ti]) - 1 - ti
			}
		}
	}
}

func move(matrix [][]int, t, i, j int) int {
	t++
	if t == 4 {
		tmp := matrix[j][len(matrix[i])-1-i]
		matrix[j][len(matrix[i])-1-i] = matrix[i][j]
		return tmp
	}
	tmp := move(matrix, t, j, len(matrix[i])-1-i)

	if t == 1 {
		matrix[j][len(matrix[i])-1-i] = tmp
		return 0
	}

	matrix[j][len(matrix[i])-1-i] = matrix[i][j]
	return tmp
}
