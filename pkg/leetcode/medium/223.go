package medium

/*
	223. Rectangle Area
*/

func computeArea(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int) int {
	a1 := (ax2 - ax1) * (ay2 - ay1)
	a2 := (bx2 - bx1) * (by2 - by1)
	cover := 0

	left := max(ax1, bx1)
	right := min(ax2, bx2)
	down := max(ay1, by1)
	top := min(ay2, by2)

	if right > left && top > down {
		cover = (right - left) * (top - down)
	}

	return a1 + a2 - cover
}
