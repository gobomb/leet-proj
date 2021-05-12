package leetcode

func lengthOfLastWord(s string) int {
	var i, c int
	for i = len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			c++
		} else {
			if i != len(s)-1 {
				if s[i+1] == ' ' {
					continue
				} else {
					return c
				}
			} else {
				continue
			}
		}
	}

	return c
}
