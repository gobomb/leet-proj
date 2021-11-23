package easy

/*
	205. Isomorphic Strings
*/

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	mp1 := make(map[byte]int)
	mp2 := make(map[byte]int)
	for i := range s {
		if mp1[s[i]] != mp2[t[i]] {
			return false
		} else {
			mp1[s[i]] = i + 1
			mp2[t[i]] = i + 1
		}
	}

	return true
}
