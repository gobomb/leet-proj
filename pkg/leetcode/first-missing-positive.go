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

func firstMissingPositive2(nums []int) int {
	lnums := len(nums)
	for i := 0; i < lnums; i++ {
		if nums[i] <= 0 {
			nums[i] = lnums + 1
		}
	}
	// fmt.Printf("%v\n", nums)

	for i := 0; i < lnums; i++ {
		n := nums[i]
		if n < 0 {
			n *= -1
		}
		if n > lnums {
			continue
		}
		if nums[n-1] > 0 {
			nums[n-1] *= -1
		}
	}
	// fmt.Printf("%v\n", nums)

	for i := 1; i <= lnums; i++ {
		if nums[i-1] > 0 {
			return i
		}
	}
	return lnums + 1
}
