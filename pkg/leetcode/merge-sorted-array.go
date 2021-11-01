package leetcode

// 0 ms	2.4 MB
func merge88(nums1 []int, m int, nums2 []int, n int) {
	a, b := m-1, n-1
	for i := m + n - 1; i >= 0; i-- {
		tmp := 0

		if a < 0 {
			b--
			tmp = nums2[b+1]
		} else if b < 0 {
			a--
			tmp = nums1[a+1]
		}

		if a >= 0 && b >= 0 {
			if nums1[a] >= nums2[b] {
				a--
				tmp = nums1[a+1]
			} else {
				b--
				tmp = nums2[b+1]
			}
		}

		nums1[i] = tmp
	}
}

// 	0 ms	2.3 MB
func merge881(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}
	for i := 0; i < m; i++ {
		if nums1[i] > nums2[0] {
			nums2[0], nums1[i] = nums1[i], nums2[0]

			for j := 0; j < n-1 && nums2[j] > nums2[j+1]; j++ {
				nums2[j], nums2[j+1] = nums2[j+1], nums2[j]
			}
		}
	}

	for k := m; k < n+m; k++ {
		nums1[k] = nums2[k-m]
	}
}
