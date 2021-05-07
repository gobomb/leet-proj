package leetcode

// 快速幂
func myPow(x float64, n int) float64 {
	if n == 1 {
		return x
	}
	if n == 0 {
		return 1
	}
	if n < 0 {
		n = -n
		x = 1 / x
	}
	// fmt.Println(n)
	res := myPow(x, n/2)

	if n%2 != 0 {

		return res * res * x
	} else {
		return res * res
	}
}
