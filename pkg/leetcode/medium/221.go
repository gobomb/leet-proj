package medium

/*
	221. Maximal Square
*/

// https://leetcode.com/problems/maximal-square/discuss/600149/Python-Thinking-Process-Diagrams-DP-Approach
func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}

	dp := make([][]int, len(matrix)+1)

	for i := range matrix {
		dp[i] = make([]int, len(matrix[0])+1)
	}

	dp[len(matrix)] = make([]int, len(matrix[0])+1)

	maxInt := 0

	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == '1' {
				dp[i+1][j+1] = min(min(dp[i][j], dp[i+1][j]), dp[i][j+1]) + 1

				if dp[i+1][j+1] > maxInt {
					maxInt = dp[i+1][j+1]
				}
			}
		}
	}

	return maxInt * maxInt
}
