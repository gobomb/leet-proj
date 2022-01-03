package medium

/*
	357. Count Numbers with Unique Digits
*/

func countNumbersWithUniqueDigits(n int) int {
	if n == 0 {
		return 1
	}

	rs, uniqueDigits, availableNumber := 10, 9, 9

	for ; n > 1 && availableNumber > 0; n-- {
		uniqueDigits *= availableNumber
		rs += uniqueDigits
		availableNumber--
	}

	return rs
}
