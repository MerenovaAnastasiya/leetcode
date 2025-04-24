package tasks

import "slices"

func Intersect(nums1 []int, nums2 []int) []int {
	l1, l2 := 0, 0

	slices.Sort(nums2)
	slices.Sort(nums1)

	var result []int

	for {
		if l1 >= len(nums1) {
			break
		}
		if l2 >= len(nums2) {
			break
		}
		if nums1[l1] > nums2[l2] {
			l2++
		} else if nums1[l1] < nums2[l2] {
			l1++
		} else {
			result = append(result, nums1[l1])
			l1++
			l2++
		}
	}

	return result

}
