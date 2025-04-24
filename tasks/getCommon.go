package tasks

func getCommon(nums1 []int, nums2 []int) int {
	// Input: nums1 = [3, 5, 6, 7], nums2 = [2, 7]
	// Output: 2

	index1 := 0
	index2 := 0

	for index1 < len(nums1) && index2 < len(nums2) {
		first := nums1[index1]
		second := nums2[index2]
		if first == second {
			return first
		} else if first > second {
			index2++
		} else {
			index1++
		}
	}
	return -1
}
