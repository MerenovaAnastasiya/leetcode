package tasks

func intersection(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	result := make([]int, 0)
	for i := 0; i < len(nums1); i++ {
		m[nums1[i]] = 1
	}
	for i := 0; i < len(nums2); i++ {
		if res, ok := m[nums2[i]]; ok {
			if res == 1 {
				result = append(result, nums2[i])
				m[nums2[i]] = 0
			}
		}
	}
	return result
}
