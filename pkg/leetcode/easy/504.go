package easy

/*
	504. Base 7
*/

func convertToBase7(num int) string {
	sign := ""

	if num < 0 {
		sign += "-"
		num = -num
	}

	s := ""

	for num != 0 {
		n := num % 7
		s = string(rune(n+'0')) + s
		num /= 7
	}

	if s == "" {
		s = "0"
	}

	return sign + s
}
