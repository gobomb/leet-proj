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

func intersect2(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	rs := []int{}
	for i := range nums1 {
		_, ok := m[nums1[i]]
		if ok {
			m[nums1[i]] = m[nums1[i]] + 1
			continue
		}
		m[nums1[i]] = 1
	}
	for i := range nums2 {
		_, ok := m[nums2[i]]
		if ok {
			rs = append(rs, nums2[i])
			m[nums2[i]] = m[nums2[i]] - 1
			if m[nums2[i]] == 0 {
				delete(m, nums2[i])
			}
		}
	}
	return rs
}
