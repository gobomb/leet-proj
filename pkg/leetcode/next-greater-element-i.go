package leetcode

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	rs := make([]int, len(nums1))
	for i, n1 := range nums1 {
		j := 0
		for ; j < len(nums2); j++ {
			if nums2[j] == n1 {
				j++
				for ; j < len(nums2); j++ {
					if nums2[j] > n1 {
						rs[i] = nums2[j]
						break
					}
				}
				break
			}
		}
		if j >= len(nums2) {
			rs[i] = -1
		}
	}
	return rs
}
