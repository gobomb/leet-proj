package medium

import "strconv"

/*
	241. Different Ways to Add Parentheses
*/

// https://leetcode.com/problems/different-ways-to-add-parentheses/discuss/66364/Golang-6ms
func diffWaysToCompute(expression string) []int {
	res := []int{}

	for i := range expression {
		if expression[i] == '*' || expression[i] == '-' || expression[i] == '+' {
			part1 := diffWaysToCompute(expression[:i])
			part2 := diffWaysToCompute(expression[i+1:])

			for _, a := range part1 {
				for _, b := range part2 {
					c := 0

					switch expression[i] {
					case '+':
						c = a + b
					case '-':
						c = a - b
					case '*':
						c = a * b
					}

					res = append(res, c)
				}
			}
		}
	}

	if len(res) == 0 {
		num, _ := strconv.Atoi(expression)
		res = append(res, num)
	}

	return res
}
