package leetcode

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * func guess(num int) int;
 */

var pick int

func guess(num int) int {
	switch {
	case num > pick:
		return -1
	case num < pick:
		return 1
	case num == pick:
		return 0
	}
	return 0
}

func guessNumber(n int) int {
	l, h, m := 1, n, 0
	for l <= h {
		m = (l + h) / 2
		g := guess(m)
		// log.Println(g)
		switch g {
		case 0:
			return m
		case 1:
			l = m + 1
		case -1:
			h = m - 1
		}
	}
	return m
}
