package medium

import "strconv"

/*
	150. Evaluate Reverse Polish Notation
*/

func evalRPN(tokens []string) int {
	stack := []int{}

	for _, t := range tokens {
		switch t {
		case "+":
			stack[len(stack)-2] += stack[len(stack)-1]
		case "-":
			stack[len(stack)-2] -= stack[len(stack)-1]
		case "*":
			stack[len(stack)-2] *= stack[len(stack)-1]
		case "/":
			stack[len(stack)-2] /= stack[len(stack)-1]
		default:
			num, err := strconv.Atoi(t)
			if err != nil {
				panic(err)
			}

			stack = append(stack, num)

			continue
		}

		stack = stack[:len(stack)-1]
	}

	return stack[0]
}
