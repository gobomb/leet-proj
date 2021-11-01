package leetcode

import (
	"fmt"
)

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

func hashSlice(nums, sl []int) string {
	var s string
	for i := range sl {
		s += fmt.Sprint(nums[sl[i]])
	}
	return s
}

var rig map[string]bool

func permuteUnique(nums []int) [][]int {
	rs := [][]int{}
	rig = make(map[string]bool)

	rs = dfsPermutii(nums, rs, []int{}, 0)

	// fmt.Println(rs)
	for i := range rs {
		for j := range rs[i] {
			rs[i][j] = nums[rs[i][j]]
		}
	}
	return rs
}

func dfsPermutii(nums []int, rs [][]int, r []int, d int) [][]int {
	if d == len(nums) {
		if d == len(nums) {
			hr := hashSlice(nums, r)
			// fmt.Println(hr)
			ok := rig[hr]
			if !ok {
				rig[hr] = true
				rcopy := make([]int, len(r))
				copy(rcopy, r)
				rs = append(rs, rcopy)
				return rs
			}
			return rs
		}
	}
	for i := 0; i < len(nums); i++ {
		for j := range r {
			if r[j] == i {
				goto loop
			}
		}
		r = append(r, i)
		rs = dfsPermutii(nums, rs, r, d+1)
		r = r[:len(r)-1]
	loop:
	}
	return rs
}
