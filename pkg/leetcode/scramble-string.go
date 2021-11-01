package leetcode

/*
	87. Scramble String
*/

var cacheScr map[string]bool

func isScramble(s1 string, s2 string) bool {
	cacheScr = make(map[string]bool)
	return isScramble2(s1, s2)
}

func isScramble2(s1 string, s2 string) bool {
	if b, ok := cacheScr[s1+s2]; ok {
		return b
	}
	if s1 == s2 {
		cacheScr[s1+s2] = true
		return true
	}
	if !cmpStrFreq(s1, s2) {
		cacheScr[s1+s2] = false
		return false
	}

	l := len(s1)
	for i := 1; i < l; i++ {
		if isScramble2(s1[0:i], s2[0:i]) && isScramble2(s1[i:], s2[i:]) {
			cacheScr[s1+s2] = true
			return true
		}
		if isScramble2(s1[0:i], s2[l-i:]) && isScramble2(s1[i:], s2[0:l-i]) {
			cacheScr[s1+s2] = true
			return true
		}
	}
	cacheScr[s1+s2] = false
	return false
}

func cmpStrFreq(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	m := make(map[byte]int)
	for i := range s1 {
		m[s1[i]]++
	}
	for i := range s2 {
		m[s2[i]]--
	}
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}
