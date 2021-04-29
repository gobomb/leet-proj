package leetcode

// 2 1 9 8 7
// 7 8 9 1 2
// 2 9 2
func firstMissingPositive(nums []int) int {
	m := make(map[int]struct{})
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = struct{}{}
	}
	for i := 1; ; i++ {
		_, ok := m[i]
		if !ok {
			return i
		}
	}
}
