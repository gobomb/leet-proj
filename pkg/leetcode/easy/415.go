package easy

/*
	415. Add Strings
*/

func addStrings(num1 string, num2 string) string {
	rs := ""
	next := byte(0)

	for i, j := len(num1)-1, len(num2)-1; ; i, j = i-1, j-1 {
		var a, b byte = 0, 0
		if i >= 0 {
			a = num1[i] - '0'
		}

		if j >= 0 {
			b = num2[j] - '0'
		}

		if next == 0 && i < 0 && j < 0 {
			break
		}

		count := a + b + next

		next = 0

		if count >= 10 {
			next = 1
			count -= 10
		}

		rs = string(count+'0') + rs
	}

	if rs == "" {
		rs = "0"
	}

	return rs
}
