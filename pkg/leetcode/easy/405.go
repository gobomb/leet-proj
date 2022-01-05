package easy

/*
	405. Convert a Number to Hexadecimal
*/

func toHex(num int) string {
	max := 4294967296
	n := 0
	rs := ""

	if num == 0 {
		rs = "0"
	} else if num > 0 {
		n = num
	} else {
		n = max + num
	}

	for n > 0 {
		a := n % 16
		n = n / 16

		if a >= 10 {
			rs = string(rune(a-10+'a')) + rs
		} else {
			rs = string(rune(a-0+'0')) + rs
		}
	}

	return rs
}
