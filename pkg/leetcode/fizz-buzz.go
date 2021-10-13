package leetcode

import "strconv"

/*
412. Fizz Buzz
*/

func fizzBuzz(n int) []string {
	rs := []string{}
	for i := 0; i < n; i++ {
		if (i+1)%15 == 0 {
			rs = append(rs, "FizzBuzz")
			continue
		}
		if (i+1)%3 == 0 {
			rs = append(rs, "Fizz")
			continue
		}
		if (i+1)%5 == 0 {
			rs = append(rs, "Buzz")
			continue
		}
		rs = append(rs, strconv.Itoa(i+1))
	}
	return rs
}
