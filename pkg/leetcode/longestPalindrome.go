package leetcode

import (
	"fmt"
)

func longestPalindrome(s string) string {
	l := len(s)
	maxStr := ""
	maxLen := 0
	for i := 0; i < l; i++ {
		for j := i; j < l; j++ {
			sl := isPalindrome(s[i : j+1])
			if sl > maxLen {
				maxLen = sl
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

func manacher(s string) string {
	l := len(s)
	t := "^#"

	for i := 0; i < l; i++ {
		t += string(s[i])
		t += "#"
	}
	t += "$"
	lt := len(t)
	p := make([]int, lt)

	center := 0
	R := 0

	for i := 1; i < lt-1; i++ {
		// get the mirror of i about center
		mirror := center*2 - i

		if R > i {
			p[i] = min(R-i, p[mirror])
		} else {
			p[i] = 0
		}

		for {
			if t[i-p[i]] == t[i+p[i]] {
				p[i]++
			} else {
				break
			}
		}

		// compute center and R
		if i+p[i] > R {
			center = i
			R = center + p[center]
		}
	}

	maxI := 0
	for i := 0; i < lt; i++ {
		if p[maxI] < p[i] {
			maxI = i
		}
	}
	fmt.Printf("%d %d", maxI, p[maxI])

	return s[(maxI-p[maxI])/2 : (maxI+p[maxI])/2-1]
}

func expendPalindrome(s string) string {
	l := len(s)
	maxStr := ""
	maxLen := 0
	for i := 0; i < l; i++ {
		l1 := expendAroundCenter(s, i, i)
		l2 := expendAroundCenter(s, i, i+1)
		ml := max(l1, l2)

		if ml > maxLen {
			maxLen = ml
			maxStr = s[i-(ml-1)/2 : i+ml/2+1]
		}
	}
	return maxStr
}

func expendAroundCenter(s string, l, r int) int {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return r - l - 1
}

func testCase(fn func(string) string) {
	b := fn("aaacccc")
	fmt.Println(b)
	b = fn("cbbd")
	fmt.Println(b)
	b = fn("ac")
	fmt.Println(b)
}

func TestPalindrome() {
	// s := "babcbabcbaccba"
	// l := len("babcbabcbaccba")
	// t := "^#"
	// for i := 0; i < l; i++ {
	// 	t += string(s[i])
	// 	t += "#"
	// }
	// t += "$"
	// fmt.Println(t)
	// tt := expendAroundCenter(t, 12-(15-12), 15)
	// print(tt)
	// return
	testCase(longestPalindrome)
	testCase(manacher)
	testCase(expendPalindrome)
}
