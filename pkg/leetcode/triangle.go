package leetcode

func minimumTotal(triangle [][]int) int {
	rs := 0
	j := 0

	for i := range triangle {
		if i == 0 {
			rs += triangle[i][j]
		} else if triangle[i][j] < triangle[i][j+1] {
			rs += triangle[i][j]
		} else {
			rs += triangle[i][j+1]
		}
	}
	return rs
}
