package leetcode

func merge88(nums1 []int, m int, nums2 []int, n int) {
	getmax := func(a, b *int) int {
		if *a < 0 {
			*b--
			return nums2[*b+1]
		}
		if *b < 0 {
			*a--
			return nums1[*a+1]
		}

		if nums1[*a] >= nums2[*b] {
			*a--
			return nums1[*a+1]
		}
		*b--
		return nums2[*b+1]
	}
	a, b := m-1, n-1
	for i := m + n - 1; i >= 0; i-- {
		largest := getmax(&a, &b)
		nums1[i] = largest
	}
	return
}
