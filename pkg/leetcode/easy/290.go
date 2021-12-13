package easy

import (
	"strings"
)

/*
	290. Word Pattern
*/

func wordPattern(pattern string, s string) bool {
	fs := strings.Fields(s)

	if len(pattern) != len(fs) {
		return false
	}

	mp := make(map[rune]string)
	exist := make(map[string]struct{})

	for i, p := range pattern {
		if v, ok := mp[p]; !ok {
			if _, ok1 := exist[fs[i]]; ok1 {
				return false
			}

			mp[p] = fs[i]
			exist[fs[i]] = struct{}{}
		} else if v != fs[i] {
			return false
		}
	}

	return true
}
