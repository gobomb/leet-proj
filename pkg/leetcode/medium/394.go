package medium

import (
	"strconv"
)

/*
	394. Decode String
	Runtime: 0 ms, faster than 100.00% of Go online submissions for Decode String.
	Memory Usage: 2.1 MB, less than 75.93% of Go online submissions for Decode String.
*/
func decodeString(s string) string {
	st := []byte{}

	for i := 0; i < len(s); i++ {
		switch {
		case s[i] == '[':
			fallthrough
		case s[i] >= '0' && s[i] <= '9':
			fallthrough
		case s[i] >= 'a' && s[i] <= 'z':
			st = append(st, s[i])
		case s[i] == ']':
			j := -2

			for j = len(st) - 1; j >= 0; j-- {
				if st[j] == '[' {
					break
				}
			}

			a := st[j+1:]
			k := -2
			numString := ""

			for k = j - 1; k >= 0; k-- {
				if !(st[k] >= '0' && st[k] <= '9') {
					break
				}

				numString = string(st[k]) + numString
			}

			num := 0

			if numString != "" {
				num, _ = strconv.Atoi(numString)
			}

			aa := []byte{}

			for ii := 0; ii < num; ii++ {
				aa = append(aa, a...)
			}

			if k != -2 {
				st = st[:k+1]
				st = append(st, aa...)
			}
		}
	}

	return string(st)
}
