package easy

/*
	168. Excel Sheet Column Title
*/

func convertToTitle(columnNumber int) string {
	var res string

	length := int('Z' - 'A' + 1)

	for columnNumber != 0 {
		n := columnNumber % length

		if n == 0 {
			n = length
			columnNumber--
		}

		res = string(rune('A'+n-1)) + res
		columnNumber /= length
	}

	return res
}
