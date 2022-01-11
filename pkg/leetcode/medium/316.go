package medium

/*
	316. Remove Duplicate Letters
*/

func removeDuplicateLetters(s string) string {
	cnt := make([]int, 26)
	visited := make([]bool, 26)
	stack := []byte{}
	rs := ""

	for i := range s {
		cnt[s[i]-'a']++
	}

	for i := range s {
		cnt[s[i]-'a']--

		if visited[s[i]-'a'] {
			continue
		}

		for len(stack) != 0 && s[i] < stack[len(stack)-1] && cnt[stack[len(stack)-1]-'a'] > 0 {
			visited[stack[len(stack)-1]-'a'] = false
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, s[i])
		visited[s[i]-'a'] = true
	}

	for i := range stack {
		rs = rs + string(rune(stack[i]))
	}

	return rs
}
