package medium

import (
	"log"
	"strconv"
)

/*
	166. Fraction to Recurring Decimal
*/

// https://leetcode.com/problems/fraction-to-recurring-decimal/discuss/51119/Golang-solution-using-math
func fractionToDecimal(numerator int, denominator int) string {
	// 处理正负号
	var flag string
	if numerator < 0 && denominator < 0 {
		numerator = -numerator
		denominator = -denominator
	} else if numerator > 0 && denominator < 0 {
		denominator, flag = -denominator, "-"
	} else if numerator < 0 && denominator > 0 {
		numerator, flag = -numerator, "-"
	}

	quotient, remainder := numerator/denominator, numerator%denominator
	res := []byte(strconv.FormatInt(int64(quotient), 10))

	if remainder == 0 {
		return flag + string(res)
	}

	res = append(res, '.')
	// 用于记录位置，检测循环
	cache := make(map[int]int)
	ind := len(res) - 1

	if remainder != 0 {
		for remainder != 0 {
			numerator = remainder * 10

			quotient, remainder = numerator/denominator, numerator%denominator

			if index, ok := cache[numerator]; ok {
				// 多加一个空格预留给 copy
				log.Printf("%s %v %v", res, index, quotient)
				res = append(res, ')', ' ')
				log.Printf("%s", res)

				copy(res[index+1:], res[index:])
				log.Printf("%s", res)

				res[index] = '('
				log.Printf("%s", res)

				break
			}

			res = append(res, digToByte(quotient))
			ind++
			// key是被除数
			cache[numerator] = ind
		}
	}
	return flag + string(res)
}

// dig是单个数字
func digToByte(dig int) byte {
	return byte('0' + dig)
}
