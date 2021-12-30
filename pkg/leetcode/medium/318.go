package medium

/*
	318. Maximum Product of Word Lengths
*/

/* 
	Runtime: 14 ms, faster than 57.69% of Go online submissions for Maximum Product of Word Lengths.
	Memory Usage: 6.8 MB, less than 15.38% of Go online submissions for Maximum Product of Word Lengths.
*/
func maxProduct2(words []string) int {
	var rs int
	bm := make([]uint32, len(words))

	for i, word := range words {
		for j := range word {
			bm[i] |= 1 << (word[j] - 'a')
		}
	}

	for i, w1 := range words {
		for j := i + 1; j < len(words); j++ {
			w2 := words[j]

			if i != j {
				if bm[i]&bm[j] == 0 && len(w1)*len(w2) > rs {
					rs = len(w1) * len(w2)
				}
			}
		}
	}

	return rs
}
/* 
	Runtime: 105 ms, faster than 7.69% of Go online submissions for Maximum Product of Word Lengths.
	Memory Usage: 7.9 MB, less than 11.54% of Go online submissions for Maximum Product of Word Lengths.
*/
func maxProduct1(words []string) int {
	var rs int

	for i, w1 := range words {
		m := buildCache(w1)

		for j := i + 1; j < len(words); j++ {
			w2 := words[j]

			if i != j {
				if len(w1)*len(w2) > rs && !compareHaveSameEle(m, w2) {
					rs = len(w1) * len(w2)
				}
			}
		}
	}

	return rs
}

func buildCache(s1 string) map[rune]struct{} {
	m := make(map[rune]struct{})

	for _, r := range s1 {
		m[r] = struct{}{}
	}

	return m
}

func compareHaveSameEle(m map[rune]struct{}, s2 string) bool {
	for _, r := range s2 {
		if _, ok := m[r]; ok {
			return true
		}
	}

	return false
}
