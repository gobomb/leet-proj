package easy

/*
	575. Distribute Candies
*/

func distributeCandies(candyType []int) int {
	c := len(candyType) / 2
	m := make(map[int]int)

	for i := range candyType {
		if c == 0 {
			break
		}
		if _, ok := m[candyType[i]]; ok {

		} else {
			m[candyType[i]] = 1
			c--
		}
	}

	return len(m)
}
