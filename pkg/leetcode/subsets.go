package leetcode

func subsets(nums []int) [][]int {
	rs := [][]int{}
	r := []int{}
	rs = append(rs, r)
	subsetsBT(nums, &rs, r, 0)
	return rs
}

func subsetsBT(nums []int, rs *[][]int, r []int, a int) {
	for i := a; i < len(nums); i++ {
		r = append(r, nums[i])
		b := make([]int, len(r))
		copy(b, r)
		*rs = append(*rs, b)
		subsetsBT(nums, rs, r, i+1)
		r = r[:len(r)-1]
	}
}
