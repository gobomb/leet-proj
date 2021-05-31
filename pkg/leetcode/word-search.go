package leetcode

func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		// word[i]
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == word[0] {
				if findword(board, i, j, word, 0) {
					return true
				}
			}
		}
	}
	return false
}

func findword(board [][]byte, a, b int, word string, wi int) bool {
	if a == len(board) || b == len(board[0]) || a == -1 || b == -1 {
		return false
	}
	// log.Printf("%v %v %v %c\n", a, b, wi, word[wi])

	if board[a][b] != word[wi] {
		return false
	}

	if wi == len(word)-1 {
		return true
	}

	// 防止往回走老路
	board[a][b] = '0'
	found := findword(board, a+1, b, word, wi+1) || findword(board, a, b+1, word, wi+1) || findword(board, a-1, b, word, wi+1) || findword(board, a, b-1, word, wi+1)
	// 改回原值
	board[a][b] = word[wi]

	return found
}
