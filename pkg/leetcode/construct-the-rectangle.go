package leetcode

func constructRectangle(area int) []int {
	var l, w, sub int
	l = area
	res := []int{0, 0}

	for w = 1; w <= l; w++ {
		if area%w != 0 {
			continue
		}

		l = area / w
		if l-w == 0 {
			return []int{l, w}
		}
		if sub == 0 || l-w < sub {
			sub = l - w
			// log.Println(l, w, sub)

			res = []int{w, l}
		}
	}
	return res
}
