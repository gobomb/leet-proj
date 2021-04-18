package leetcode

func findSubstring(s string, words []string) []int {
	rs := make([]int, 0)
	countMap := make(map[string]int)
	resetCount(countMap, words)

	for i := range s {
		t := matchWords(s, words, countMap, i, 0)

		if t == len(words) {
			rs = append(rs, i)
		}
		resetCount(countMap, words)
	}
	return rs
}

func matchWords(s string, words []string, countMap map[string]int, i int, t int) int {
	for j := range words {
		if countMap[words[j]] != 0 {
			if matchWord(i, s, words[j]) {
				countMap[words[j]]--
				t++
				return matchWords(s, words, countMap, i+len(words[j]), t)
			}
		}
	}
	return t
}

func resetCount(countMap map[string]int, words []string) {
	for k := range countMap {
		delete(countMap, k)
	}
	for _, w := range words {
		countMap[w]++
	}
}

func matchWord(i int, s, word string) bool {
	for k := range word {
		if i+k >= len(s) {
			return false
		}
		if s[i+k] == word[k] {
		} else {
			return false
		}
	}
	return true
}
