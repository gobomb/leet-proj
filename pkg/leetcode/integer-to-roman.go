package leetcode

var d = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
var rm = []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

func intToRoman(num int) string {
	var a int
	var r string
	last := num
	for i := 0; i < len(d); i++ {
		for {
			a = last - d[i]
			if a < 0 {
				break
			}
			r = r + rm[i]
			last = a
		}
	}
	return string(r)
}

func testIntToRoman() {
	tests := []int{
		3, 4, 9, 59, 1994,
	}
	for i := range tests {
		r := intToRoman(tests[i])
		println(r)
	}
}
