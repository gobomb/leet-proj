package leetcode

// type loc struct {
// 	x int
// 	y int
// }

var pathMemo [][]int

func uniquePaths(m int, n int) int {
	pathMemo = make([][]int, m)
	for i := 0; i < m; i++ {
		pathMemo[i] = make([]int, n)
	}
	return dpUniquePaths(m, n, 0, 0)
}

func uniquePathsLoop(m, n int) int {
	pathMemo = make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		pathMemo[i] = make([]int, n+1)
	}

	for i := m - 1; i >= 0; i-- {
		var down, right int

		for j := n - 1; j >= 0; j-- {
			if i == m-1 && j == n-1 {
				pathMemo[i][j] = 1
				continue
			}
			if pathMemo[i+1][j] != 0 {
				down = pathMemo[i+1][j]
			}
			if pathMemo[i][j+1] != 0 {
				right = pathMemo[i][j+1]
			}
			pathMemo[i][j] = down + right
		}
	}
	return pathMemo[0][0]
}

func dpUniquePaths(m, n, a, b int) int {
	if a == m || b == n {
		return 0
	}
	if a == m-1 && b == n-1 {
		return 1
	}
	var down, right int

	if a+1 != m {
		if pathMemo[a+1][b] != 0 {
			down = pathMemo[a+1][b]
		} else {
			down = dpUniquePaths(m, n, a+1, b)
			pathMemo[a+1][b] = down
		}
	}
	if b+1 != n {
		if pathMemo[a][b+1] != 0 {
			right = pathMemo[a][b+1]
		} else {
			right = dpUniquePaths(m, n, a, b+1)
			pathMemo[a][b+1] = right
		}
	}

	pathMemo[a][b] = down + right
	return pathMemo[a][b]
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				obstacleGrid[i][j] = -1
			}
		}
	}
	pathMemo = obstacleGrid
	rs := dpUniquePathsObs(m, n, 0, 0)

	return rs
}

func dpUniquePathsObs(m, n, a, b int) int {
	if pathMemo[a][b] == -1 {
		return 0
	}
	if a == m || b == n {
		return 0
	}
	if a == m-1 && b == n-1 {
		return 1
	}

	var down, right int

	if a+1 != m {
		if pathMemo[a+1][b] > 0 {
			down = pathMemo[a+1][b]
		} else if pathMemo[a+1][b] == 0 {
			down = dpUniquePathsObs(m, n, a+1, b)
			pathMemo[a+1][b] = down
		}
	}
	if b+1 != n {
		if pathMemo[a][b+1] > 0 {
			right = pathMemo[a][b+1]
		} else if pathMemo[a][b+1] == 0 {
			right = dpUniquePathsObs(m, n, a, b+1)
			pathMemo[a][b+1] = right
		}
	}
	pathMemo[a][b] = down + right

	return pathMemo[a][b]
}
