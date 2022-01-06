package easy

/*
	409. Longest Palindrome
*/

func longestPalindrome(s string) int {
	cache := map[int]bool{}

	for _, b := range s {
		cache[int(b)] = !cache[int(b)]
	}

	count := 0

	for _, v := range cache {
		if v {
			count++
		}
	}

	if count == 0 || count == 1 {
		return len(s)
	}

	return len(s) - count + 1
}
