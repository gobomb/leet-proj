package medium

/*
	289. Game of Life
*/

/*
	0 d->d
	1 l->l
	2 l->d
	3 d->l
*/
func gameOfLife(board [][]int) {
	var l1, l2 int

	l1 = len(board)

	if l1 > 0 {
		l2 = len(board[0])
	}

	dx := []int{-1, -1, -1, 0, 1, 1, 1, 0}
	dy := []int{-1, 0, 1, 1, 1, 0, -1, -1}

	for i := range board {
		for j := range board[i] {
			cnt := 0

			for k := 0; k < 8; k++ {
				x, y := i+dx[k], j+dy[k]
				if x >= 0 && x < l1 && y >= 0 && y < l2 &&
					(board[x][y] == 1 || board[x][y] == 2) {
					cnt++
				}
			}

			if board[i][j]%2 == 1 {
				if cnt < 2 || cnt > 3 {
					board[i][j] = 2
				} else if cnt == 2 || cnt == 3 {
					board[i][j] = 1
				}
			} else if board[i][j]%2 == 0 {
				if cnt == 3 {
					board[i][j] = 3
				}
			}
		}
	}

	for i := range board {
		for j := range board[i] {
			board[i][j] %= 2
		}
	}
}
