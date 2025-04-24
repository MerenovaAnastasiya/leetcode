package tasks

func MaximumUniqueSubarray(nums []int) int {
	m := make(map[int]int)
	l, sum, maxSum := 0, 0, 0

	for r := 0; r < len(nums); r++ {
		sum += nums[r]
		if _, ok := m[nums[r]]; ok {
			for l <= m[nums[r]] {
				sum -= nums[l]
				l++
			}
		}
		m[nums[r]] = r
		maxSum = max(maxSum, sum)
	}
	return maxSum
}
