package leetcode

func findTheDifference(s string, t string) byte {
	m := make(map[byte]int)
	for i := range s {
		m[s[i]] = m[s[i]] + 1
	}
	for i := range t {
		m[t[i]] = m[t[i]] - 1
		if m[t[i]] < 0 {
			return t[i]
		}
	}

	return t[len(t)-1]
}
