package tasks

func findMaxK(nums []int) int {
	maxK := -1
	m := map[int]bool{}

	for i := 0; i < len(nums); i++ {
		if _, ok := m[nums[i]*-1]; ok {
			maxK = max(maxK, nums[i], nums[i]*-1)
		} else {
			m[nums[i]] = true
		}
	}
	return maxK
}
