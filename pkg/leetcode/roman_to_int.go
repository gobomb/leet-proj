package leetcode

import "fmt"

func romanToInt(s string) int {
	svmap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	var result int
	var lastV int
	for i := range s {
		v := svmap[s[i]]
		if lastV < v {
			// 假设了数据是合法的，IV、IX这种情况属于判断的子集
			result -= 2 * lastV
		}
		result += v
		lastV = v
	}

	return result
}

func testRoamToInt() {
	r := romanToInt("IIII")
	fmt.Println(r)
}
