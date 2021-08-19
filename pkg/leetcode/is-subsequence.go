package leetcode

func isSubsequence(s string, t string) bool {
	j := 0
	for i := range s {
		for ; j <= len(t); j++ {
			if j == len(t) {
				return false
			}
			// log.Printf("%v %v %c %c", i, j, s[i], t[j])

			if s[i] == t[j] {
				j++
				break
			}

		}

	}
	return true
}
