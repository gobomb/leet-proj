package leetcode

var memoPath [][]int

func minPathSum(grid [][]int) int {
	memoPath = make([][]int, len(grid)+1)
	for i := 0; i < len(grid)+1; i++ {
		memoPath[i] = make([]int, len(grid[0])+1)
		for j := 0; j < len(grid[0])+1; j++ {
			memoPath[i][j] = -1
		}
		memoPath[i][len(grid[0])] = -2
	}

	for j := 0; j < len(grid[0])+1; j++ {
		memoPath[len(grid)][j] = -2
	}
	rs := dpMinPathSum(grid, 0, 0)
	return rs
}

func dpMinPathSum(grid [][]int, a, b int) int {
	checkMemo := func(a, b int) int {
		if memoPath[a][b] == -1 {
			memoPath[a][b] = dpMinPathSum(grid, a, b)
		}
		return memoPath[a][b]
	}
	down := checkMemo(a+1, b)
	right := checkMemo(a, b+1)
	if down < 0 && right < 0 {
		return grid[a][b]
	}
	if down < 0 {
		return right + grid[a][b]
	}
	if right < 0 {
		return down + grid[a][b]
	}
	if down >= right {
		return right + grid[a][b]
	}

	return down + grid[a][b]
}
