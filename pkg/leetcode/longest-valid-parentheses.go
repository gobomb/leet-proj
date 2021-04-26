package leetcode

func longestValidParentheses(s string) int {
	if s == "" {
		return 0
	}
	var left, right, valid = 0, 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left++
		}
		if s[i] == ')' {
			right++
		}

		if right == left {
			valid = max(left*2, valid)
		} else if right > left {
			right, left = 0, 0
		}
	}

	if left > right {
		left, right = 0, 0
		for i := len(s) - 1; i >= 0; i-- {
			if s[i] == '(' {
				left++
			}
			if s[i] == ')' {
				right++
			}

			if right == left {
				valid = max(left*2, valid)
			} else if left > right {
				right, left = 0, 0
			}
		}
	}

	return valid
}
