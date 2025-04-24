package tasks

func FindLHS(nums []int) int {
	m := make(map[int]int)
	maxLHS := 0

	for i := 0; i < len(nums); i++ {
		if val, ok := m[nums[i]]; ok {
			m[nums[i]] = val + 1
			maxLHS = max(maxLHS, val+1+m[nums[i]+1], val+1+m[nums[i]-1])
		} else {
			m[nums[i]] = 1
			maxLHS = max(maxLHS, 1+m[nums[i]+1], val+1+m[nums[i]-1])
		}
		maxLHS = max(maxLHS, m[nums[i]])
		if val, ok := m[nums[i]+1]; ok {
			maxLHS = max(maxLHS, val+m[nums[i]])
		}
		if val, ok := m[nums[i]-1]; ok {
			maxLHS = max(maxLHS, val+m[nums[i]])
		}
	}

	return maxLHS
}
