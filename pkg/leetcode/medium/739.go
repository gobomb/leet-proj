package medium

/*
	739. Daily Temperatures

*/

// 1044 ms	9.7 MB
func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))

	for i := range temperatures {
		for j := i; j < len(temperatures); j++ {
			if temperatures[j] > temperatures[i] {
				res[i] = j - i
				break
			}
		}
	}

	return res
}

// 140 ms	9.8 MB
func dailyTemperaturesMonotonousStack(temperatures []int) []int {
	res := make([]int, len(temperatures))
	// stack := make([]int, 0, len(temperatures))
	stack := []int{}

	for i := len(temperatures) - 1; i >= 0; i-- {
		for len(stack) != 0 && temperatures[i] >= temperatures[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) != 0 {
			res[i] = stack[len(stack)-1] - i
		} else {
			res[i] = 0
		}

		stack = append(stack, i)
	}

	return res
}
