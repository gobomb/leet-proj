package leetcode

/*
	263. Ugly Number
*/
func isUgly(n int) bool {
	if n <= 0 {
		return false
	}
	var b bool
	for {
		if n%2 == 0 {
			n /= 2
			b = true
		}
		if n%3 == 0 {
			n /= 3
			b = true
		}
		if n%5 == 0 {
			n /= 5
			b = true
		}
		if !b {
			return n == 1
		}
		b = false
	}
}
