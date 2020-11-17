package leetcode

import "fmt"

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	n := len(s)
	cycle := 2*numRows - 2
	ret := make([]byte, 0)
	for i := 0; i < numRows; i++ {
		for j := 0; j < n-i; j += cycle {
			ret = append(ret, s[i+j])
			if i != 0 && i != numRows-1 && j+cycle-i < n {
				ret = append(ret, s[j+cycle-i])
			}
		}
	}

	return string(ret)
}

func testConvert(fn func(string, int) string) {
	var b string
	b = fn("PAYPALISHIRING", 3)
	fmt.Println(b)
	b = fn("PAYPALISHIRING", 4)
	fmt.Println(b)
	b = fn("ABCDEFGHIJKLM", 3)
	fmt.Println(b)
}

func TestConvert() {
	testConvert(convert)
}
