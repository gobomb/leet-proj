package easy

import "math"

/*
	171. Excel Sheet Column Number
*/

func titleToNumber(columnTitle string) int {
	var res int

	for i := len(columnTitle) - 1; i >= 0; i-- {
		j := len(columnTitle) - i - 1
		res += int(columnTitle[i]-'A'+1) * int(math.Pow('Z'-'A'+1, float64(j)))
	}

	return res
}
