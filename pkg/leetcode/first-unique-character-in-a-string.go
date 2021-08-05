package leetcode

func firstUniqChar(s string) int {
	var m [26]int
	for _, c := range s {
		m[c-'a']++
	}
	for i, c := range s {
		if m[c-'a'] == 1 {
			return i
		}
	}
	return -1
}
