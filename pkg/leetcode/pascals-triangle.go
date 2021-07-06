package leetcode

func generate(numRows int) [][]int {
	var rs [][]int
	for i := 0; i < numRows; i++ {
		rs = append(rs, make([]int, i+1))
		for j := 0; j < i+1; j++ {
			if j == 0 || j == i {
				rs[i][j] = 1
			} else {
				rs[i][j] = rs[i-1][j-1] + rs[i-1][j]
			}
		}
	}
	return rs
}
