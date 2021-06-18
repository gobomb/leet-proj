package leetcode

func intersection(nums1 []int, nums2 []int) []int {
	m := make(map[int]struct{})
	rs := []int{}
	for i := range nums1 {
		m[nums1[i]] = struct{}{}
	}
	for i := range nums2 {
		_, ok := m[nums2[i]]
		if ok {
			rs = append(rs, nums2[i])
			delete(m, nums2[i])
		}
	}
	return rs
}
