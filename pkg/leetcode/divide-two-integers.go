package leetcode

import (
	"fmt"
	"math"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func divide(dividend int, divisor int) (rs int) {
	defer func() {
		// 默认是64位int，代码不会溢出，但题目要求是32位，超出32位的结果就判定为溢出
		/*
			Assume we are dealing with an environment that could only store integers within the 32-bit signed integer range: [−231, 231 − 1].
			For this problem, assume that your function returns 231 − 1 when the division result overflows.
		*/
		if rs > math.MaxInt32 || rs < math.MinInt32 {
			rs = math.MaxInt32
		}
	}()
	var sign int = -1

	if dividend == 0 {
		return 0
	}
	if divisor == 1 {
		return dividend
	}
	if divisor == -1 {
		return -dividend
	}

	// 计算符号
	if (dividend < 0 && divisor < 0) || (dividend > 0 && divisor > 0) {
		sign = 1
	} else {
		sign = -1
	}

	// 去除符号
	sdd := Abs(dividend)
	sdr := Abs(divisor)

	dd := sdd
	dr := sdr
	// 总的商
	// rs := 0
	// 除数倍数
	quo := 1
	// 余数
	var remainder int
	for {
		// 计算余数，被除数-除数
		remainder = dd - dr // 64 - 32

		// 被除数小于原始除数，结束
		if remainder < 0 {
			break
		}

		// 余数等于0，商+除数倍数，返回
		if remainder == 0 {
			rs += quo
			break
		}

		// 余数等于原始除数，商+除数倍数+余数本身，返回
		if remainder == sdr {
			rs += quo
			rs += 1
			break
		}
		// 余数等于除数，商+除数倍数+余数倍数，返回
		if remainder == dr {
			rs += quo
			rs += quo
			break
		}

		// 余数小于除数
		if remainder < dr {
			// 除数倍数加到总的商里
			rs += quo
			// 除数倍数重置为1
			quo = 1
			// 被除数重置为余数
			dd = remainder
			// 除数重置为原始的除数
			dr = sdr
			// 继续计算余数
			continue
		}
		// 余数大于除数
		if remainder > dr {
			// 除数左移一位（乘以2）
			dr <<= 1
			// 除数倍数左移一位
			quo <<= 1
			// 继续计算余数
			continue
		}
	}

	return rs * sign
}

func testDivide() {
	var r int
	divide(10, 3)
	divide(19, 5)
	divide(7, -3)
	divide(2, 5)

	divide(0, 1)
	divide(8000, 1)
	divide(64, 1)
	r = divide(math.MaxInt32, 1)
	fmt.Printf("result %v \n", r)

	r = divide(math.MinInt32, 1)
	fmt.Printf("result %v \n", r)
	r = divide(math.MinInt32, -1)
	fmt.Printf("result %v \n", r)
	fmt.Println(math.MaxInt32)
	fmt.Println(math.MinInt32)

}
