package tasks

func Merge(nums1 []int, m int, nums2 []int, n int) {
	if len(nums1) == m {
		return
	}

	if len(nums1) == n {
		copy(nums1, nums2)
		return
	}

	k, j, curr := m-1, n-1, m+n-1

	for k >= 0 && j >= 0 {
		if nums1[k] < nums2[j] {
			nums1[curr] = nums2[j]
			j--
		} else {
			nums1[curr] = nums1[k]
			k--
		}
		curr--
	}

	for j >= 0 {
		nums1[curr] = nums1[j]
		curr--
		j--
	}

}
