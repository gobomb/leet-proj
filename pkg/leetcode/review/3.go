package review

/*
	3. Longest Substring Without Repeating Characters
*/

/*
	Runtime: 0 ms, faster than 100.00% of Go online submissions for Longest Substring Without Repeating Characters.
	Memory Usage: 2.5 MB, less than 91.00% of Go online submissions for Longest Substring Without Repeating Characters.
*/
func lengthOfLongestSubstring(s string) int {
	cache := [128]int{}

	left, right, maxLen := 0, 0, 0

	for right < len(s) {
		c := s[right]
		right++
		cache[c]++

		// cache 中的字符大于1，触发左指针向前走
		// 同时要更新cache
		for cache[c] > 1 {
			cache[s[left]]--
			left++
		}

		if right-left > maxLen {
			maxLen = right - left
		}
	}

	return maxLen
}
