func merge(nums1 []int, m int, nums2 []int, n int) {
	last := m + n - 1
	l1, l2 := m-1, n-1

	for last >= 0 {
		if l1 >= 0 && l2 >= 0 && nums1[l1] > nums2[l2] {
			nums1[last] = nums1[l1]
			l1--
		} else if l2 >= 0 {
			nums1[last] = nums2[l2]
			l2--
		}

		last--
	}
}
