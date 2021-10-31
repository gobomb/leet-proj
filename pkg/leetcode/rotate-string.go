package leetcode

import "strings"

/*
	796. Rotate String
*/

func rotateString(s string, goal string) bool {
	if goal == "" && s == "" {
		return true
	}
	if len(s) != len(goal) || s == "" || goal == "" {
		return false
	}
	ss := s + s
	return strings.Contains(ss, goal)
}
