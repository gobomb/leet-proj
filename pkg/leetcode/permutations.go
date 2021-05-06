package leetcode

func permute(nums []int) [][]int {
	rs := [][]int{}
	rs = dfsPermut(nums, rs, []int{}, 0)
	return rs
}

func dfsPermut(nums []int, rs [][]int, r []int, d int) [][]int {
	if d == len(nums) {
		// temp := ds.DeepCopyIntSlice(r)
		rcopy := make([]int, len(r))
		copy(rcopy, r)
		rs = append(rs, rcopy)
		return rs
	}
	for i := 0; i < len(nums); i++ {
		for j := range r {
			if r[j] == nums[i] {
				goto loop
			}
		}
		r = append(r, nums[i])
		rs = dfsPermut(nums, rs, r, d+1)
		r = r[:len(r)-1]
	loop:
	}
	return rs
}
