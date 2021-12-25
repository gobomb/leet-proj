package medium

import (
	"strconv"
	"strings"
)

/*
	306. Additive Number
*/
func isAdditiveNumber(num string) bool {
	if len(num) < 3 {
		return false
	}

	for i := 1; i < len(num); i++ {
		if i > 1 && strings.HasPrefix(num[0:i], "0") {
			break
		}

		a, _ := strconv.Atoi(num[0:i])

		for j := i + 1; j < len(num); j++ {
			if j-i > 1 && strings.HasPrefix(num[i:j], "0") {
				break
			}

			b, _ := strconv.Atoi(num[i:j])

			if isAdditiveNumber1(num, a, b, j) {
				return true
			}
			// log.Println()
		}
	}

	return false
}

func isAdditiveNumber1(num string, a, b, j int) bool {
	if j == len(num) {
		// log.Println("j=len")
		return true
	}

	for i := j + 1; i <= len(num); i++ {
		if i-j > 1 && strings.HasPrefix(num[j:i], "0") {
			return false
		}

		c, _ := strconv.Atoi(num[j:i])
		// log.Println(a, b, c, j, i)
		if a+b == c {
			return isAdditiveNumber1(num, b, c, i)
		}
	}

	return false
}
