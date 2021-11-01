package leetcode

type direction int

const (
	right direction = iota
	down
	left
	up
)

type top int

func turn(d *direction) {
	if *d == up {
		*d = right
		return
	}
	*d++
}

func generateMatrix(n int) [][]int {
	rs := make([]int, n*n)
	for i := 0; i < len(rs); i++ {
		rs[i] = i + 1
	}
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	k := 0
	i := k
	j := k

	rt := len(matrix[0])
	lt := -1
	ut := -1
	dt := len(matrix)

	var d direction
	for ; ; k++ {
		if k == len(rs) {
			return matrix
		}

		switch d {
		case right:
			matrix[i][j] = rs[k]
			j++
		case down:
			matrix[i][j] = rs[k]

			i++
		case left:
			matrix[i][j] = rs[k]

			j--
		case up:
			matrix[i][j] = rs[k]

			i--
		}

		if j == rt {
			ut = i

			turn(&d)
			j--
			i++
			continue
		}
		if i == dt {
			rt = j

			turn(&d)
			i--
			j--
			continue
		}
		if j == lt {
			dt = i

			turn(&d)
			j++
			i--
			continue
		}

		if i == ut {
			lt = j

			turn(&d)
			i++
			j++
			continue
		}

	}
}

func spiralOrder(matrix [][]int) []int {
	rs := make([]int, len(matrix)*len(matrix[0]))

	k := 0
	i := k
	j := k

	rt := len(matrix[0])
	lt := -1
	ut := -1
	dt := len(matrix)

	var d direction
	for ; ; k++ {
		// log.Printf("%v\n", rs)
		if k == len(rs) {
			return rs
		}

		switch d {
		case right:
			rs[k] = matrix[i][j]
			j++
		case down:
			rs[k] = matrix[i][j]
			i++
		case left:
			rs[k] = matrix[i][j]
			j--
		case up:
			rs[k] = matrix[i][j]
			i--
		}

		if j == rt {
			ut = i

			turn(&d)
			j--
			i++
			continue
		}
		if i == dt {
			rt = j

			turn(&d)
			i--
			j--
			continue
		}
		if j == lt {
			dt = i

			turn(&d)
			j++
			i--
			continue
		}

		if i == ut {
			lt = j

			turn(&d)
			i++
			j++
			continue
		}

	}
}
