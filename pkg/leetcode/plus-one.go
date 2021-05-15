package leetcode

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i]+1 == 10 {
			digits[i] = 0
		} else {
			digits[i]++
			break
		}
	}
	if digits[0] == 0 {
		var ndigits = []int{1}
		digits = append(ndigits, digits...)
	}
	return digits
}
