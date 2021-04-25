package leetcode

func longestValidParentheses(s string) int {
	if s == "" {
		return 0
	}
	st := NewStack()
	for i := range s {
		if s[i] == '(' {
			st.Push(s[i])
		}
		if s[i] == ')' {
			if st.Peek() != nil {
				if st.Peek() == '(' {
					st.Pop()
				} else {
					
				}
			}
		}
	}
	return 0
}
