package medium

/*
	304. Range Sum Query 2D - Immutable
*/

type NumMatrix struct {
	matrix [][]int
}

func ConstructorMatrix(matrix [][]int) NumMatrix {
	return NumMatrix{matrix: matrix}
}

func (n *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	rs := 0

	for i := row1; i <= row2; i++ {
		for j := col1; j <= col2; j++ {
			rs += n.matrix[i][j]
		}
	}

	return rs
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
