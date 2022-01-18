package easy

import "strings"

/*
	434. Number of Segments in a String
*/

func countSegments(s string) int {
	return len(strings.Fields(s))
}
