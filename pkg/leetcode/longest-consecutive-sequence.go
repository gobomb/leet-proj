package leetcode

/*
	128. Longest Consecutive Sequence
*/
func longestConsecutive(nums []int) int {
	set := make(map[int]int)
	res := 0
	for _, n := range nums {
		if _, ok := set[n]; ok {
		} else {
			r := set[n+1]
			l := set[n-1]

			sum := l + r + 1
			set[n] = sum
			if r > 0 {
				set[n+r] = sum
			}
			if l > 0 {
				set[n-l] = sum
			}
			if sum > res {
				res = sum
			}
		}
	}
	return res
}
