package leetcode

import "strings"

/*
	500. Keyboard Row
*/

func findWords(words []string) []string {
	rows := []string{"qwertyuiop", "asdfghjkl", "zxcvbnm"}
	rs := []string{}
	for _, w := range words {
		if len(w) == 0 {
			continue
		}
		s := strings.ToLower(w)
		oneRow := false
		for _, r := range rows {
			if strings.ContainsAny(s, r) {
				oneRow = !oneRow
				if !oneRow {
					break
				}
			}
		}
		if oneRow {
			rs = append(rs, w)
		}
	}
	return rs
}

// https://leetcode.com/problems/keyboard-row/discuss/548831/Golang-O(N)-simple-solution-using-int-array

func findWords2(words []string) []string {
	rs := []string{}
	rows := []int{1, 2, 2, 1, 0, 1, 1, 1, 0, 1, 1, 1, 2, 2, 0, 0, 0, 0, 1, 0, 0, 2, 0, 2, 0, 2}
	for _, w := range words {
		s := strings.ToLower(w)
		level := rows[s[0]-'a']
		b := true
		for i := 1; i < len(s); i++ {
			if level != rows[s[i]-'a'] {
				b = false
				break
			}
		}
		if b {
			rs = append(rs, w)
		}
	}
	return rs
}
