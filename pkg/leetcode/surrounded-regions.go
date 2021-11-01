package leetcode

/*
	130. Surrounded Regions
*/

func solve(board [][]byte) {
	for i := range board {
		for j := range board[i] {
			if i == 0 || j == 0 || i == len(board)-1 || j == len(board[i])-1 {
				if board[i][j] == 'O' {
					dfsSolve(i, j, board)
				}
			}
		}
	}
	// for i := range board {
	// 	for j := range board[i] {
	// 		fmt.Printf("%s ", string(board[i][j]))
	// 	}
	// 	log.Println()
	// }
	for i := range board {
		for j := range board[i] {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == '*' {
				board[i][j] = 'O'
			}
		}
	}
}

func dfsSolve(i, j int, board [][]byte) {
	if i < 0 || j < 0 || i > len(board)-1 || j > len(board[i])-1 {
		return
	}
	if board[i][j] == 'O' {
		board[i][j] = '*'
		dfsSolve(i, j-1, board)
		dfsSolve(i, j+1, board)
		dfsSolve(i-1, j, board)
		dfsSolve(i+1, j, board)
	}
}
