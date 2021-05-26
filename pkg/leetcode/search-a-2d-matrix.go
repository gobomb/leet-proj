package leetcode

// import "log"
func searchMatrix(matrix [][]int, target int) bool {
	rowl := len(matrix[0])
	lowi := 0
	hii := (len(matrix)-1)*rowl + (rowl - 1)
	x, y := 0, 0

	for hii >= lowi {
		x, y = ((hii+lowi)/2)/rowl, ((hii+lowi)/2)%rowl

		if matrix[x][y] == target {
			return true
		} else if matrix[x][y] > target {
			hii = ((hii + lowi) / 2) - 1
		} else {
			lowi = ((hii + lowi) / 2) + 1
		}
	}
	return false
}

func searchMatrix2(matrix [][]int, target int) bool {
	rowl := len(matrix[0])

	type pos struct {
		x, y int
	}

	gi := func(s pos) int {
		return s.x*rowl + s.y
	}
	gp := func(i int) pos {
		return pos{
			i / rowl,
			i % rowl,
		}
	}

	mpf := func(e, s int) int {
		return (e + s) / 2
	}

	low := pos{
		0, 0,
	}
	hi := pos{
		len(matrix) - 1, rowl - 1,
	}

	lowi := gi(low)
	hii := gi(hi)

	for hii >= lowi {
		mi := mpf(hii, lowi)
		m := gp(mi)
		// log.Printf("%v\n", m)

		if matrix[m.x][m.y] == target {
			return true
		} else if matrix[m.x][m.y] > target {
			hii = mi - 1
		} else {
			lowi = mi + 1
		}
	}
	return false
}
