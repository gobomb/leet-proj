package dynamicprogram

/*
	473. Matchsticks to Square
*/

func makesquare(matchsticks []int) bool {
	if len(matchsticks) < 4 {
		return false
	}

	sum := 0

	for i := range matchsticks {
		sum += matchsticks[i]
	}

	if sum%4 != 0 {
		return false
	}

	side := sum / 4

	// 全部可能的组合
	all := 1<<len(matchsticks) - 1

	// 和等于side的组合
	okSubset := []int{}
	// 和等于2*side的组合
	validSet := make([]bool, 1<<len(matchsticks))

	for i := 0; i <= all; i++ {
		sum := 0

		for j := 0; j < 32; j++ {
			if (i>>j)&1 == 1 {
				sum += matchsticks[j]
			}
		}

		if sum == side {
			// 符合条件的组合
			for _, v := range okSubset {
				// 遍历已有的复合条件的组合
				if v&i != 0 {
					// 看是否有重复的元素
					// 1110
					// 0011
					continue
				}

				// 1110 0000
				// 0001 0000
				// 1111 0000
				half := v | i
				validSet[half] = true

				// 1111 1111
				// 1111 0000
				// 0000 1111
				// 与 all 异或得到另一半组合
				if validSet[all^half] {
					return true
				}
			}
			// 全部都会存下来
			okSubset = append(okSubset, i)
		}
	}

	return false
}
