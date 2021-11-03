package medium

/*
	137. Single Number II
*/

// https://blog.csdn.net/yutianzuijin/article/details/50597413
// https://books.halfrost.com/leetcode/ChapterFour/0100~0199/0137.Single-Number-II/
func singleNumber(nums []int) int {
	hi, lo := 0, 0
	for _, n := range nums {
		lo = ^hi & (lo ^ n)
		hi = ^lo & (hi ^ n)
	}
	return lo
}
