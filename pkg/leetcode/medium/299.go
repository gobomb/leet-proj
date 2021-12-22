package medium

import "strconv"

/*
	299. Bulls and Cows
*/

func getHint(secret string, guess string) string {
	a, b := 0, 0
	set := make(map[byte]int)
	s := []byte{}

	for i := range secret {
		if secret[i] == guess[i] {
			a++
		} else {
			set[secret[i]]++
			s = append(s, guess[i])
		}
	}

	for _, v := range s {
		if _, ok := set[v]; ok {
			set[v]--
			b++

			if set[v] == 0 {
				delete(set, v)
			}
		}
	}

	return strconv.Itoa(a) + "A" + strconv.Itoa(b) + "B"
}
