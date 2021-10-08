package leetcode

/*
	125. Valid Palindrome
*/

import (
	"strings"
)

func isPalindrome2(s string) bool {
	s = strings.ToLower(s)
	t := ""

	for i := range s {
		if isAlphanumeric(s[i]) {
			t += string(s[i])
			continue
		}
	}

	return isPalindrome(t) != 0
}

func isAlphanumeric(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= '0' && b <= '9')
}

/*
	Runtime: 0 ms, faster than 100.00% of Go online submissions for Valid Palindrome.
	Memory Usage: 2.8 MB, less than 68.36% of Go online submissions for Valid Palindrome.
*/
func isPalindrome4(s string) bool {
	s = strings.ToLower(s)

	l := len(s)

	for i, j := 0, l-1; i < j; i++ {
		if !isAlphanumeric(s[i]) {
			continue
		}
	checkJ:
		if !isAlphanumeric(s[j]) {
			j--
			goto checkJ
		}
		if s[i] != s[j] {
			return false
		}
		j--
	}
	return true
}
