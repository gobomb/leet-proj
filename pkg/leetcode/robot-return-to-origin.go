package leetcode

func judgeCircle(moves string) bool {
	u := [4]int{}

	for i := range moves {
		switch moves[i] {
		case 'U':
			u[0]++
		case 'D':
			u[1]++
		case 'R':
			u[2]++
		case 'L':
			u[3]++
		}
	}
	return u[0] == u[1] && u[2] == u[3]
}
