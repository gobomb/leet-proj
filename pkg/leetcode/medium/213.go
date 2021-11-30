package medium

/*
	213. House Robber II
*/

/*
	Runtime: 0 ms, faster than 100.00% of Go online submissions for House Robber II.
	Memory Usage: 2 MB, less than 70.39% of Go online submissions for House Robber II.
*/
func rob2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	r1 := rob1(nums[:len(nums)-1])
	r2 := rob1(nums[1:])

	return max(r1, r2)
}
