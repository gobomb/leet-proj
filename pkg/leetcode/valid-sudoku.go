package leetcode

func isValidSudoku(board [][]byte) bool {
	row := make([][]bool, 9)
	col := make([][]bool, 9)
	box := make([][]bool, 9)
	for i := 0; i < 9; i++ {
		row[i] = make([]bool, 9)
		col[i] = make([]bool, 9)
		box[i] = make([]bool, 9)
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			b := board[i][j]
			if b == '.' {
				continue
			}
			bi := b - '0' - 1
			if bi > 8 {
				return false
			}
			if row[i][bi] || col[j][bi] || box[i/3*3+j/3][bi] {
				return false
			}
			row[i][bi] = true
			col[j][bi] = true
			box[i/3*3+j/3][bi] = true
		}
	}
	return true
}
