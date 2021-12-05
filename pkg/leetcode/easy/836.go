package easy

/*
	836. Rectangle Overlap
*/

// https://leetcode.com/problems/rectangle-overlap/discuss/1103101/Golang-solution-faster-than-100-with-explanation-and-images
func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	left := max(rec1[0], rec2[0])
	right := min(rec1[2], rec2[2])
	down := max(rec1[1], rec2[1])
	top := min(rec1[3], rec2[3])

	return right > left && top > down
}
