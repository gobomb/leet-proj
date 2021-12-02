package easy

/*
	231. Power of Two
*/

func isPowerOfTwo(n int) bool {
	return n&(n-1) == 0 && n > 0
}
