package medium

/*
	227. Basic Calculator II
*/

func calculate(s string) int {
	var p int
	res := calculate2(s, &p)
	return res
}

// https://github.com/labuladong/fucking-algorithm/blob/master/%E6%95%B0%E6%8D%AE%E7%BB%93%E6%9E%84%E7%B3%BB%E5%88%97/%E5%AE%9E%E7%8E%B0%E8%AE%A1%E7%AE%97%E5%99%A8.md
func calculate2(s string, p *int) int {
	res := 0
	stack := []int{}
	sigh := uint8('+')
	num := 0

	for i := 0; i < len(s); i++ {
		ch := s[i]

		isDigit := '0' <= ch && ch <= '9'

		if isDigit {
			num = num*10 + int(ch-'0')
		}

		if ch == '(' {
			num = calculate2(s[i+1:], p)
			i = i + *p + 1
		}

		if !isDigit && ch != ' ' || i == len(s)-1 {
			if sigh == '+' {
				stack = append(stack, num)
			} else if sigh == '-' {
				stack = append(stack, -num)
			} else if sigh == '*' {
				stack[len(stack)-1] *= num
			} else if sigh == '/' {
				stack[len(stack)-1] /= num
			}

			num = 0
			sigh = ch
		}

		if ch == ')' {
			*p = i
			break
		}
	}

	for i := range stack {
		res += stack[i]
	}

	return res
}

/*
def calculate(s: str) -> int:

    def helper(s: List) -> int:
        stack = []
        sign = '+'
        num = 0

        while len(s) > 0:
            c = s.pop(0)
            if c.isdigit():
                num = 10 * num + int(c)

            if (not c.isdigit() and c != ' ') or len(s) == 0:
                if sign == '+':
                    stack.append(num)
                elif sign == '-':
                    stack.append(-num)
                elif sign == '*':
                    stack[-1] = stack[-1] * num
                elif sign == '/':
                    # python 除法向 0 取整的写法
                    stack[-1] = int(stack[-1] / float(num))
                num = 0
                sign = c

        return sum(stack)
    # 需要把字符串转成列表方便操作
    return helper(list(s))

*/
