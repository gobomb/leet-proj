package medium

/*
	200. Number of Islands

*/

/*
	只访问1，并留下访问记录
*/

func numIslands(grid [][]byte) int {
	res := 0
	// flag := make([][]bool, len(grid))

	// for i := range flag {
	// 	flag[i] = make([]bool, len(grid[i]))
	// }

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				findIsland(&grid, i, j)
				res++
			}
		}
	}
	return res
}

func findIsland(grid *[][]byte, i, j int) {
	if i < 0 || j < 0 || i >= len(*grid) || j >= len((*grid)[i]) || (*grid)[i][j] == '0' || (*grid)[i][j] == '2' {
		return
	}
	// flag[i][j] = true
	(*grid)[i][j] = '2'

	findIsland(grid, i+1, j)
	findIsland(grid, i, j+1)
	findIsland(grid, i-1, j)
	findIsland(grid, i, j-1)
}
