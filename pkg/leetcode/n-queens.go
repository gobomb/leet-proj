package leetcode

func solveNQueens(n int) [][]string {
	var rst [][]string
	if n == 0 {
		return [][]string{{}}
	}
	if n == 1 {
		return [][]string{{"Q"}}
	}

	rs := make([][]string, n)
	for i := range rs {
		rs[i] = make([]string, n)
		for j := range rs[i] {
			rs[i][j] = "."
		}
	}

	dfsnqueens(rs, &rst, 0, n)

	return rst
}


func makeresult(rs [][]string) []string {
	var rss []string
	for i := range rs {
		temp := ""
		for j := range rs {
			temp = temp + rs[i][j]
		}
		rss = append(rss, temp)
	}
	return rss
}




func dfsnqueens(rs [][]string, rst *[][]string, r, qs int) {
	if qs == 0 {
		// fmt.Printf("%v\n", (rs))
		*rst = append(*rst, makeresult(rs))
		return
	}
	for j := 0; j < len(rs); j++ {
		if rs[r][j] != "Q" && checkQueen(r, j, rs) {
			rs[r][j] = "Q"
			qs--
			dfsnqueens(rs, rst, r+1, qs)
			rs[r][j] = "."
			qs++
		}
		if rs[r][j] != "Q" {
			rs[r][j] = "."
		}
	}
	return
}

func checkQueen(a, b int, rs [][]string) bool {
	n := len(rs)
	for i := 0; i < len(rs); i++ {
		if rs[a][i] == "Q" {
			return false
		}
		if rs[i][b] == "Q" {
			return false
		}
		if a+i < n && b+i < n && rs[a+i][b+i] == "Q" {
			return false
		}
		if a-i >= 0 && b-i >= 0 && rs[a-i][b-i] == "Q" {
			return false
		}
		if a-i >= 0 && b+i < n && rs[a-i][b+i] == "Q" {
			return false
		}
		if a+i < n && b-i >= 0 && rs[a+i][b-i] == "Q" {
			return false
		}
	}
	return true
}




func totalNQueens(n int) int {
	rs := make([][]string, n)
	for i := range rs {
		rs[i] = make([]string, n)
		for j := range rs[i] {
			rs[i][j] = "."
		}
	}
	var rst2 = new(int)
	dfsnqueens2(rs, rst2, 0, n)

	return *rst2
}

func dfsnqueens2(rs [][]string, rst *int, r, qs int) {
	if qs == 0 {
		*rst++
		return
	}
	for j := 0; j < len(rs); j++ {
		if rs[r][j] != "Q" && checkQueen(r, j, rs) {
			rs[r][j] = "Q"
			qs--
			dfsnqueens2(rs, rst, r+1, qs)
			rs[r][j] = "."
			qs++
		}
		if rs[r][j] != "Q" {
			rs[r][j] = "."
		}
	}
	return
}