package leetcode

import "log"

func bitset(s string) int {
	if len(s) == 0 {
		return 0
	}

	var freq [256]uint8
	result, left, right := 0, 0, 0

	for left < len(s) {
		if right < len(s) && freq[s[right]] == 0 {
			freq[s[right]] = 1
			right++
		} else {
			freq[s[left]] = 0
			left++
		}

		result = max(result, right-left)
	}

	return result
}

func longestSubstringWithoutRepeatingCharacters(s string) int {
	if len(s) == 0 {
		return 0
	}

	var freq [256]bool
	result, left, right := 0, 0, 0

	for left < len(s) {
		if right < len(s) && !freq[s[right]-'a'] {
			freq[s[right]-'a'] = true
			right++
		} else {
			freq[s[left]-'a'] = false
			left++
		}

		result = max(result, right-left)
	}

	return result
}

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// func min(a, b int) int {
// 	if a > b {
// 		return b
// 	}
// 	return a
// }

func TestLong() {
	log.Printf("%d\n", longestSubstringWithoutRepeatingCharacters("bbbbb"))
	log.Printf("%d\n", longestSubstringWithoutRepeatingCharacters("abcabcbb"))
	log.Printf("%d\n", longestSubstringWithoutRepeatingCharacters("pwwkew"))

	log.Printf("%d\n", bitset("bbbbb"))
	log.Printf("%d\n", bitset("abcabcbb"))
	log.Printf("%d\n", bitset("pwwkew"))
}
