package leetcode

/*
	242. Valid Anagram
*/

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := make(map[byte]int)
	for i := range s {
		m[s[i]]++
	}
	for j := range t {
		m[t[j]]--
	}
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}
