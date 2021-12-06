package hard

// https://leetcode.com/problems/basic-calculator/discuss/1457982/Clear-Go-Solution
func calculate(s string) int {
	res, num, op := 0, 0, add

	for i := 0; i < len(s); i++ {
		switch s[i] {
		case ' ':
		case '+':
			res, num = op(res, num), 0
			op = add

		case '-':
			res, num = op(res, num), 0
			op = sub
		case '(':
			end := findEnd(s, i)
			num, i = calculate(s[i+1:end]), end
		default:
			num *= 10
			num += int(s[i] - '0')
		}
	}
	res, num = op(res, num), 0

	return res
}

func findEnd(s string, start int) int {
	lv := 0
	for i := start; i < len(s); i++ {
		switch s[i] {
		case '(':
			lv++
		case ')':
			lv--
			if lv == 0 {
				return i
			}
		}
	}
	return len(s) - 1
}

func add(a, b int) int {
	return a + b
}
func sub(a, b int) int {
	return a - b
}

// // wrong answer
// func calculate1(s string) int {
// 	num := ""
// 	readNum := false
// 	numStack := []int{}
// 	symbolStack := []byte{}
// 	sign := false

// 	appendNum := func() {
// 		if !readNum && len(num) != 0 {

// 			n, err := strconv.Atoi(num)
// 			if err != nil {
// 				return
// 			}
// 			if sign {
// 				n = -n
// 				sign = false
// 			}
// 			numStack = append(numStack, n)

// 			num = ""
// 		}
// 	}

// 	count := func() {
// 		n2 := numStack[len(numStack)-1]
// 		n1 := numStack[len(numStack)-2]
// 		numStack = numStack[:len(numStack)-2]
// 		symbol := symbolStack[len(symbolStack)-1]
// 		symbolStack = symbolStack[:len(symbolStack)-1]

// 		if symbol == '+' {
// 			sum := n1 + n2
// 			numStack = append(numStack, sum)
// 		} else if symbol == '-' {
// 			diff := n1 - n2
// 			numStack = append(numStack, diff)
// 		}
// 	}

// 	for _, c := range s {
// 		if c == ' ' {
// 			readNum = false
// 			appendNum()

// 			continue
// 		} else if c == '(' {
// 			readNum = false
// 			appendNum()

// 			symbolStack = append(symbolStack, '(')
// 		} else if c == ')' {
// 			readNum = false
// 			appendNum()

// 			for len(symbolStack) != 0 && symbolStack[len(symbolStack)-1] != '(' {
// 				count()
// 			}
// 			if len(symbolStack) != 0 && symbolStack[len(symbolStack)-1] == '(' {
// 				symbolStack = symbolStack[:len(symbolStack)-1]
// 			}
// 		} else if c >= '0' && c <= '9' {
// 			num = num + string(c)
// 			readNum = true
// 		} else if c == '+' {
// 			readNum = false
// 			appendNum()
// 			if len(symbolStack) != 0 && symbolStack[len(symbolStack)-1] != '(' {
// 				count()
// 			}
// 			symbolStack = append(symbolStack, '+')

// 		} else if c == '-' {
// 			readNum = false
// 			appendNum()
// 			sign = true
// 			if len(numStack) != 0 && len(symbolStack) != 0 && symbolStack[len(symbolStack)-1] != '(' {
// 				count()
// 			}
// 			symbolStack = append(symbolStack, '+')
// 			numStack = append(numStack, 0)
// 			symbolStack = append(symbolStack, '+')

// 		}

// 		log.Println(numStack)
// 		log.Printf("%s\n\n", symbolStack)

// 	}

// 	if len(num) != 0 {
// 		readNum = false
// 		appendNum()
// 	}

// 	if len(numStack) != 1 {
// 		for len(symbolStack) != 0 && symbolStack[len(symbolStack)-1] != '(' {
// 			count()
// 		}
// 		if len(symbolStack) != 0 && symbolStack[len(symbolStack)-1] == '(' {
// 			symbolStack = symbolStack[:len(symbolStack)-1]
// 		}
// 	}

// 	return numStack[0]
// }
