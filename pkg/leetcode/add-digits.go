package leetcode

/*
	258. Add Digits
*/

func addDigits(num int) int {
	var rs int
	rs = num
	for {
		if rs/10 == 0 {
			return rs
		}
		temp := 0

		for {
			a := rs % 10
			b := rs / 10
			if b != 0 {
				temp += a
				rs = b
			} else {
				temp += a
				rs = temp
				break
			}
		}
	}
}
