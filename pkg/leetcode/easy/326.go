package easy

/*
	326. Power of Three
*/

// https://books.halfrost.com/leetcode/ChapterFour/0300~0399/0326.Power-of-Three/
func isPowerOfThree(n int) bool {
	// 1162261467 is 3^19,  3^20 is bigger than int
	return n > 0 && (1162261467%n == 0)
}

