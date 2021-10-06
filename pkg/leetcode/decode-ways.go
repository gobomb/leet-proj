package leetcode

/*
	91. Decode Ways
*/

import (
	"strconv"
)

func checkNum(s string) bool {
	l := len(s)
	if l == 0 || l > 2 || s[0] == '0' {
		return false
	}
	n, err := strconv.Atoi(s)
	if err != nil || n <= 0 || n > 26 {
		return false
	}
	return true
}

// 去掉map，使用 checkNum 来检查数字是否合法
// 重新归纳递推的两种情况：可组合/不可组合
func numDecodings2(s string) int {
	if s == "" {
		return 0
	}

	rs := make([]int, len(s))

	// 讨论 i==0
	if checkNum(s[0:1]) {
		rs[0] = 1
	} else {
		rs[0] = 0
		return rs[0]
	}

	if len(s) == 1 {
		return rs[0]
	}

	// 讨论 i==1
	if checkNum(s[:2]) {
		// 可组合
		rs[1]++
	}
	if checkNum(s[1:2]) {
		// 可不组合
		rs[1]++
	}

	for i := 2; i < len(s); i++ {
		rs[i] = 0
		if checkNum(s[i-1 : i+1]) {
			// 可组合
			rs[i] += rs[i-2]
		}
		if checkNum(s[i : i+1]) {
			// i 单独合法
			rs[i] += rs[i-1]
		}
	}

	return rs[len(s)-1]
}

func numDecodings(s string) int {
	if s == "" {
		return 0
	}

	numMap := make(map[string]byte)

	var i int64
	for i = 1; i <= 26; i++ {
		numMap[strconv.FormatInt(i, 10)] = 'A' + byte(i) - 1
	}

	rs := make([]int, len(s))

	// 讨论 i==0
	if s[0] == '0' {
		rs[0] = 0
		return rs[0]
	} else {
		rs[0] = 1
	}

	if len(s) == 1 {
		return rs[0]
	}

	// 讨论 i==1

	// 组合
	t := 0
	if _, ok := numMap[s[:2]]; ok {
		t++
	}
	// 不组合
	if s[1] != '0' {
		t++
	}
	// 2/1
	rs[1] = t

	for i := 2; i < len(s); i++ {
		if s[i-1] == '0' {
			// 前一位为 0 ，无法组合
			if s[i] != '0' {
				rs[i] = rs[i-1]
			} else {
				// 00 非法
				rs[i] = 0
			}
		} else {
			if s[i] == '0' {
				// 当前为 0，必须组合
				if _, ok := numMap[s[i-1:i+1]]; ok {
					// 组合合法
					rs[i] = rs[i-2]
				} else {
					rs[i] = 0
				}
			} else {
				if _, ok := numMap[s[i-1:i+1]]; ok {
					// 组合合法
					rs[i] = rs[i-1] + rs[i-2]
				} else {
					// 组合非法，当前不为0，使用上一个结果
					rs[i] = rs[i-1]
				}
			}
		}
	}

	// for k, v := range numMap {
	// 	log.Printf("%+v %c", k, v)
	// }
	return rs[len(s)-1]
}
