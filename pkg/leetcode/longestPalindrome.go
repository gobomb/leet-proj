package leetcode

import "fmt"

func longestPalindrome(s string) string {
	l := len(s)
	maxStr := ""
	maxLen := 0
	for i := 0; i < l; i++ {
		for j := i; j < l; j++ {
			l := isPalindrome(s[i : j+1])
			if l > maxLen {
				maxLen = l
				maxStr = s[i : j+1]
			}
		}
	}
	return maxStr
}

func isPalindrome(s string) int {
	l := len(s)
	if l%2 == 0 {
		for i, j := l/2, l/2-1; i < l; i++ {

			if s[i] != s[j] {
				return 0
			}
			j--
		}
		return l
	}
	for i, j := l/2+1, l/2-1; i < l; i++ {
		if s[i] != s[j] {
			return 0
		}
		j--
	}
	return l
}

func TestPalindrome() {
	b := longestPalindrome("babad")
	fmt.Println(b)
	b = longestPalindrome("cbbd")
	fmt.Println(b)
	b = longestPalindrome("ac")
	fmt.Println(b)
}
