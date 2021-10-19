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
