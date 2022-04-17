package review

/*
	3. Longest Substring Without Repeating Characters
	6min
*/

func lengthOfLongestSubstring(s string) int {
	var i, j, maxN int

	m := make(map[byte]int)

	for j < len(s) {
		m[s[j]]++

		for m[s[j]] > 1 {
			m[s[i]]--
			i++
		}

		maxN = max(maxN, j-i+1)

		j++
	}
	return maxN
}
