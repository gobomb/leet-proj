package leetcode

type position struct {
	x int
	y int
}

func solveSudoku(board [][]byte) {
	pos, find := []position{}, false
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == '.' {
				pos = append(pos, position{
					x: i, y: j,
				})
			}
		}
	}
	dfsSudoku(&board, pos, &find, 0)
	return
}

func dfsSudoku(board *[][]byte, pos []position, find *bool, index int) {
	if *find {
		return
	}
	if index >= len(pos) {
		*find = true

		return
	}

	for i := 1; i < 10; i++ {
		bi := byte(i) + '0'
		if checkSudoku(board, pos[index], bi) && !*find {
			(*board)[pos[index].x][pos[index].y] = bi
			dfsSudoku(board, pos, find, index+1)
			if *find {
				return
			}
			(*board)[pos[index].x][pos[index].y] = '.'
		}
	}
}

func checkSudoku(board *[][]byte, pos position, val byte) bool {
	for i := 0; i < len(*board); i++ {
		if (*board)[pos.x][i] == val {
			return false
		}
		if (*board)[i][pos.y] == val {
			return false
		}
	}

	posi, posy := pos.x/3*3, pos.y/3*3

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if (*board)[i+posi][j+posy] == val {
				return false
			}
		}
	}

	return true
}
