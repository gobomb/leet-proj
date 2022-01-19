package easy

import (
	"strings"
)

/*
	434. Number of Segments in a String
*/

func countSegments(s string) int {
	return len(strings.Fields(s))
}

func countSegments1(s string) int {
	rs := 0

	s = s + " "

	for i := 1; i < len(s); i++ {
		if s[i] == ' ' && s[i-1] != ' ' {
			rs++
		}
	}

	return rs
}

func countSegments2(s string) int {
	rs := 0

	for i := 1; i < len(s); i++ {
		if s[i] == ' ' && s[i-1] != ' ' {
			rs++
		}
	}

	if len(s) != 0 && s[len(s)-1] != ' ' {
		rs++
	}

	return rs
}
